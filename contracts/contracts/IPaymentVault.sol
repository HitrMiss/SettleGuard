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
        uint256 r;
        uint256 s;
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
        bytes32 _packetId,
        uint64 _createdAt,
        uint256 r,
        uint256 s
    ) external;

    function prepareSettlementBulk(
        bytes32[] calldata _packetIds,
        address _settlementWallet
    ) external;


    function refund(bytes32 _packetId) external;

    function getPaymentDetails(bytes32 _packetId) external view returns (
        address payer,            // 1
        uint64 createdAt,         // 2
        Status status,            // 3
        address merchantIdentity, // 4
        address merchantPayout,   // 5
        uint256 amount,           // 6
        bytes32 categoryId,       // 7
        uint256 r,                // 8
        uint256 s                 // 9
    );
}