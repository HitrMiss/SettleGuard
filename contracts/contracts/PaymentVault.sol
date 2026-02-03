// SPDX-License-Identifier: MIT
pragma solidity ^0.8.33;

import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";
import "@openzeppelin/contracts/utils/ReentrancyGuard.sol";
import "./IGovernance.sol";
import "./IPaymentVault.sol";

contract PaymentVault is IPaymentVault, ReentrancyGuard {
    using SafeERC20 for IERC20;

    IGovernance public immutable governance;
    IERC20 public immutable usdc;
    address public settlementEngine;

    mapping(bytes32 => Payment) public payments;

    error Unauthorized();
    error AlreadyExists();
    error NotActive();

    constructor(address _usdc, address _gov) {
        usdc = IERC20(_usdc);
        governance = IGovernance(_gov);
    }

    modifier onlyEngine() {
        if (msg.sender != settlementEngine) revert Unauthorized();
        _;
    }

    function setEngine(address _engine) external {
        if (!governance.hasRole(governance.GOVERNANCE_ADMIN_ROLE(), msg.sender)) revert Unauthorized();
        settlementEngine = _engine;
    }

    function lockFunds(
        address _payer,
        uint256 _amount,
        address _merchantIdentity,
        address _merchantPayout,
        bytes32 _categoryId,
        bytes32 _packetId
    ) external onlyEngine nonReentrant {
        if (payments[_packetId].status != Status.NONE) revert AlreadyExists();

        payments[_packetId] = Payment({
            payer: _payer,
            merchantIdentity: _merchantIdentity,
            merchantPayout: _merchantPayout,
            amount: _amount,
            categoryId: _categoryId,
            createdAt: uint64(block.timestamp),
            status: Status.HELD
        });

        usdc.safeTransferFrom(_payer, address(this), _amount);
    }

    function prepareSettlementBulk(bytes32[] calldata _packetIds, address _settlementWallet)
    external onlyEngine nonReentrant
    {
        uint256 total = 0;
        for (uint256 i = 0; i < _packetIds.length; ) {
            Payment storage p = payments[_packetIds[i]];
            if (p.status == Status.HELD) {
                p.status = Status.RELEASED;
                total += p.amount;
            }
            unchecked { ++i; }
        }
        usdc.safeTransfer(_settlementWallet, total);
    }

    function getPaymentDetails(bytes32 _packetId) external view override returns (
        address payer,
        uint256 amount,
        uint64 createdAt,
        address merchantIdentity,
        bytes32 categoryId,
        address merchantPayout,
        Status status
    ) {
        Payment storage p = payments[_packetId];
        return (
            p.payer,
            p.amount,
            p.createdAt,
            p.merchantIdentity,
            p.categoryId,
            p.merchantPayout,
            p.status
        );
    }

    function updateSettlementEngine(address _newEngine) external {
        if (!governance.hasRole(governance.GOVERNANCE_ADMIN_ROLE(), msg.sender)) revert Unauthorized();
        settlementEngine = _newEngine;
    }

    function refund(bytes32 _packetId) external override onlyEngine nonReentrant {
        Payment storage p = payments[_packetId];

        if (p.status != Status.HELD) revert NotActive();

        p.status = Status.REFUNDED;

        usdc.safeTransfer(p.payer, p.amount);

        emit PaymentRefunded(_packetId, p.payer, p.amount);
    }
}