// SPDX-License-Identifier: MIT
pragma solidity ^0.8.33;

import "./IGovernance.sol";
import "./IBondVault.sol";
import "@openzeppelin/contracts/utils/ReentrancyGuard.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";

contract BondVault is IBondVault, ReentrancyGuard {
    using SafeERC20 for IERC20;

    IGovernance public immutable governance;
    IERC20 public immutable usdc;
    address public settlementEngine;

    mapping(address => uint256) public bonds;
    mapping(address => uint64)  public unlockAt;
    mapping(address => uint256) public lockedRisk;

    uint64 public constant DEPOSIT_LOCK_SECS = 7 days;
    uint64 public constant CLAIM_MIN_AGE = 1 minutes;

    struct WithdrawClaim {
        uint256 amount;
        uint256 fee;
        uint64  requestedAt;
    }
    mapping(address => WithdrawClaim) public withdrawClaims;


    address public immutable feeRecipient;
    uint64  public immutable withdrawFeeBps;
    uint256 public immutable withdrawFlatFee;

    error Unauthorized();
    error InsufficientBond();
    error Locked(uint64 unlockAt);
    error ZeroAddress();
    error ZeroAmount();
    error AlreadyRequested();
    error NoWithdrawRequest();
    error RequestExceedsBond();
    error ClaimTooFresh(uint256 readyAt);
    error OnlyEngine();

    event Deposited(address indexed user, uint256 amount, uint64 unlockAt);
    event Withdrawn(address indexed user, uint256 amount, uint256 fee);
    event WithdrawRequested(address indexed user, uint256 amount, uint256 fee, uint64 requestedAt);
    event WithdrawProcessed(address indexed user, uint256 amount, uint256 fee, address indexed by);
    event WithdrawCanceled(address indexed user);
    event Slashed(address indexed user, address indexed recipient, uint256 amount, address indexed by);


    //TODO: FIX THE FEE LOGIC, time crunch!!!
    constructor(
        address _gov,
        address _usdc,
        address _feeRecipient,
        uint64 _withdrawFeeBps,
        uint256 _withdrawFlatFee
    ) {
        if (_gov == address(0) || _usdc == address(0) || _feeRecipient == address(0)) revert ZeroAddress();
        governance = IGovernance(_gov);
        usdc = IERC20(_usdc);
        feeRecipient = _feeRecipient;
        withdrawFeeBps = _withdrawFeeBps;
        withdrawFlatFee = _withdrawFlatFee;
    }

    modifier onlyEngine() {
        if (msg.sender != settlementEngine) revert OnlyEngine();
        _;
    }

    function lockRisk(address _merchant, uint256 _amount) external onlyEngine {
        if (bonds[_merchant] - lockedRisk[_merchant] < _amount) revert InsufficientBond();
        lockedRisk[_merchant] += _amount;
    }

    function unlockRisk(address _merchant, uint256 _amount) external onlyEngine {
        lockedRisk[_merchant] -= _amount;
    }

    function getAvailableBond(address _merchant) public view override returns (uint256) {
        return bonds[_merchant] - lockedRisk[_merchant];
    }

    function getActiveBond(address _party) external view override returns (uint256) {
        return bonds[_party];
    }

    function deposit(uint256 _amount) external nonReentrant {
        _deposit(msg.sender, _amount);
        usdc.safeTransferFrom(msg.sender, address(this), _amount);
    }

    function depositFor(address _user, uint256 _amount) external nonReentrant {
        _deposit(_user, _amount);
        usdc.safeTransferFrom(msg.sender, address(this), _amount);
    }

    function _deposit(address _user, uint256 _amount) internal {
        if (_amount == 0) revert ZeroAmount();
        if (_user == address(0)) revert ZeroAddress();

        bonds[_user] += _amount;
        uint64 newUnlock = uint64(block.timestamp) + DEPOSIT_LOCK_SECS;
        if (newUnlock > unlockAt[_user]) {
            unlockAt[_user] = newUnlock;
        }

        emit Deposited(_user, _amount, unlockAt[_user]);
    }
    function _calcWithdrawFee(uint256 amount) internal view returns (uint256) {
        return withdrawFlatFee + ((amount * uint256(withdrawFeeBps)) / 10_000);
    }

    function withdraw(uint256 _amount) external nonReentrant {
        _withdraw(msg.sender, _amount);
    }

    function withdrawMax() external nonReentrant {
        uint256 bal = bonds[msg.sender];
        if (bal <= withdrawFlatFee) revert InsufficientBond();

        // Math Check: Total = Amount + (Amount * Bps / 10000) + Flat
        // More MATH!  = (Total - Flat) * 10000 / (10000 + Bps)
        uint256 amount = ((bal - withdrawFlatFee) * 10_000) / (10_000 + withdrawFeeBps);
        _withdraw(msg.sender, amount);
    }

    function _withdraw(address _user, uint256 _amount) internal {
        if (_amount == 0) revert ZeroAmount();
        if (block.timestamp < unlockAt[_user]) revert Locked(unlockAt[_user]);

        if (withdrawClaims[_user].amount != 0) {
            delete withdrawClaims[_user];
            emit WithdrawCanceled(_user);
        }

        uint256 fee = _calcWithdrawFee(_amount);
        uint256 total = _amount + fee;
        uint256 bal = bonds[_user];

        if (bal < total) revert InsufficientBond();

        bonds[_user] = bal - total;
        usdc.safeTransfer(_user, _amount);
        if (fee != 0) usdc.safeTransfer(feeRecipient, fee);

        emit Withdrawn(_user, _amount, fee);
    }

    function requestWithdraw(uint256 _amount) external nonReentrant {
        if (_amount == 0) revert ZeroAmount();
        if (block.timestamp < unlockAt[msg.sender]) revert Locked(unlockAt[msg.sender]);
        if (withdrawClaims[msg.sender].amount != 0) revert AlreadyRequested();

        uint256 fee = _calcWithdrawFee(_amount);
        uint256 total = _amount + fee;

        if (bonds[msg.sender] < total) revert InsufficientBond();

        withdrawClaims[msg.sender] = WithdrawClaim({
            amount: _amount,
            fee: fee,
            requestedAt: uint64(block.timestamp)
        });

        emit WithdrawRequested(msg.sender, _amount, fee, uint64(block.timestamp));
    }

    function cancelWithdrawRequest() external nonReentrant {
        if (withdrawClaims[msg.sender].amount == 0) revert NoWithdrawRequest();
        delete withdrawClaims[msg.sender];
        emit WithdrawCanceled(msg.sender);
    }

    function processWithdraw(address _user) external nonReentrant {
        if (!governance.hasRole(governance.ARBITER_ROLE(), msg.sender)) revert Unauthorized();

        WithdrawClaim memory c = withdrawClaims[_user];
        if (c.amount == 0) revert NoWithdrawRequest();

        // Double-dip prevention: Ensure claim is at least 1 minute old
        if (block.timestamp < c.requestedAt + CLAIM_MIN_AGE) {
            revert ClaimTooFresh(c.requestedAt + CLAIM_MIN_AGE);
        }
        if (block.timestamp < unlockAt[_user]) revert Locked(unlockAt[_user]);

        uint256 total = c.amount + c.fee;
        uint256 bal = bonds[_user];
        if (bal < total) revert RequestExceedsBond();

        bonds[_user] = bal - total;
        delete withdrawClaims[_user];

        usdc.safeTransfer(_user, c.amount);
        if (c.fee != 0) usdc.safeTransfer(feeRecipient, c.fee);

        emit WithdrawProcessed(_user, c.amount, c.fee, msg.sender);
    }

    // Slash dot dash
    function slashBond(address _user, uint256 _amount, address _recipient) external nonReentrant {
        if (!governance.hasRole(governance.ARBITER_ROLE(), msg.sender)) revert Unauthorized();
        if (_user == address(0) || _recipient == address(0)) revert ZeroAddress();
        if (_amount == 0) revert ZeroAmount();

        uint256 bal = bonds[_user];
        if (bal == 0) revert InsufficientBond();

        uint256 amountToSlash = _amount > bal ? bal : _amount;
        uint256 newBal = bal - amountToSlash;
        bonds[_user] = newBal;

        WithdrawClaim memory c = withdrawClaims[_user];
        if (c.amount != 0 && newBal < (c.amount + c.fee)) {
            delete withdrawClaims[_user];
            emit WithdrawCanceled(_user);
        }

        usdc.safeTransfer(_recipient, amountToSlash);
        emit Slashed(_user, _recipient, amountToSlash, msg.sender);
    }

    function setSettlementEngine(address _newEngine) external {
        if (!governance.hasRole(governance.GOVERNANCE_ADMIN_ROLE(), msg.sender)) revert Unauthorized();
        if (_newEngine == address(0)) revert ZeroAddress();
        settlementEngine = _newEngine;
    }
}