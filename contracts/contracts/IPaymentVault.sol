// SPDX-License-Identifier: MIT
pragma solidity ^0.8.33;

interface IPaymentVault {
    enum Status { NONE, HELD, RELEASED, REFUNDED }

    struct Payment {
        address payer;
        uint64 createdAt;
        Status status;
        address merchantIdentity;
        address merchantPayout;
        uint256 amount;
        bytes32 categoryId;
    }

    event PaymentRefunded(
        bytes32 indexed packetId,
        address indexed payer,
        uint256 amount
    );

    function lockFunds(
        address _payer,
        uint256 _amount,
        address _merchantIdentity,
        address _merchantPayout,
        bytes32 _categoryId,
        bytes32 _packetId
    ) external;

    function prepareSettlementBulk(
        bytes32[] calldata _packetIds,
        address _settlementWallet
    ) external;


    function refund(bytes32 _packetId) external;

    function getPaymentDetails(bytes32 _packetId) external view returns (
        address payer,
        uint256 amount,
        uint64 createdAt,
        address merchantIdentity,
        bytes32 categoryId,
        address merchantPayout,
        Status status
    );
}