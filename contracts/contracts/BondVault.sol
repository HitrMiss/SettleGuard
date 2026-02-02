// SPDX-License-Identifier: MIT
pragma solidity ^0.8.33;

import "./IGovernance.sol";
import "@openzeppelin/contracts/utils/ReentrancyGuard.sol";
//TODO: Stressed revisit this later
contract BondVault is ReentrancyGuard {
    IGovernance public immutable governance;

    mapping(address => uint256) public bonds;
    mapping(address => uint64)  public unlockAt;

    uint64 public constant DEPOSIT_LOCK_SECS = 7 days;

    error Unauthorized();
    error InsufficientBond();
    error TransferFailed();
    error Locked(uint64 unlockAt);
    error ZeroAddress();
    error ZeroAmount();

    event Deposited(address indexed user, uint256 amount, uint64 unlockAt);
    event Withdrawn(address indexed user, uint256 amount);
    event Slashed(address indexed user, address indexed recipient, uint256 amount, address indexed by);

    constructor(address _gov) {
        if (_gov == address(0)) revert ZeroAddress();
        governance = IGovernance(_gov);
    }

    /// @dev Allows users to deposit ETH directly to the contract
    receive() external payable {
        _deposit(msg.sender, msg.value);
    }

    function deposit() external payable {
        _deposit(msg.sender, msg.value);
    }

    /// @notice Allows depositing on behalf of a user
    function depositFor(address _user) external payable {
        _deposit(_user, msg.value);
    }

    // Note really happy about this however using for now as I don't have a better way to deal with rolling locks
    function _deposit(address _user, uint256 _amount) internal {
        if (_amount == 0) revert ZeroAmount();
        if (_user == address(0)) revert ZeroAddress();

        bonds[_user] += _amount;

        // Rolling lock gathers no moss
        uint64 newUnlock = uint64(block.timestamp) + DEPOSIT_LOCK_SECS;
        if (newUnlock > unlockAt[_user]) {
            unlockAt[_user] = newUnlock;
        }

        emit Deposited(_user, _amount, unlockAt[_user]);
    }

    function withdraw(uint256 _amount) external nonReentrant {
        if (_amount == 0) revert ZeroAmount();
        if (block.timestamp < unlockAt[msg.sender]) revert Locked(unlockAt[msg.sender]);

        uint256 bal = bonds[msg.sender];
        if (bal < _amount) revert InsufficientBond();

        bonds[msg.sender] = bal - _amount;

        (bool success, ) = payable(msg.sender).call{value: _amount}("");
        if (!success) revert TransferFailed();

        emit Withdrawn(msg.sender, _amount);
    }

    function slash(address _user, uint256 _amount, address _recipient) external nonReentrant {
        if (!governance.hasRole(governance.ARBITER_ROLE(), msg.sender)) revert Unauthorized();
        if (_recipient == address(0)) revert ZeroAddress();
        if (_amount == 0) revert ZeroAmount();

        uint256 bal = bonds[_user];
        if (bal == 0) revert InsufficientBond();

        uint256 amountToSlash = _amount > bal ? bal : _amount;

        bonds[_user] = bal - amountToSlash;

        (bool success, ) = payable(_recipient).call{value: amountToSlash}("");
        if (!success) revert TransferFailed();

        emit Slashed(_user, _recipient, amountToSlash, msg.sender);
    }
}