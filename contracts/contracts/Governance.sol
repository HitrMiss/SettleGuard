// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.33;

import "@openzeppelin/contracts/access/AccessControl.sol";

contract Governance is AccessControl {
    bytes32 public constant GOVERNANCE_ADMIN_ROLE = keccak256("GOVERNANCE_ADMIN_ROLE");
    bytes32 public constant ARBITER_ROLE          = keccak256("ARBITER_ROLE");

    /// @dev Slow down any possible take over
    uint64 public immutable MIN_SOLE_ADMIN_AGE;

    /// @dev Mapping for min time admin existed
    mapping(address => uint64) public governanceAdminAddedAt;

    error EmptyGovernanceAdmins();
    error ZeroAddress();
    error CannotRemoveLastGovernanceAdmin();
    error RemainingAdminTooNew(uint64 addedAt, uint64 nowTs, uint64 minAge);
    error MinAgeTooLarge(uint64 provided);

    constructor(address[] memory coldWalletAdmins) {
        if (coldWalletAdmins.length == 0) revert EmptyColdWalletAdmins();

        // Default to 7 days if caller passes 0.
        uint64 minAge = minSoleAdminAgeSeconds == 0 ? uint64(7 days) : minSoleAdminAgeSeconds;
        // Just in case someone gets cheeky
        if (minAge > uint64(365 days)) revert MinAgeTooLarge(minAge);

        MIN_SOLE_ADMIN_AGE = minAge;

        //Ensure at least one admin always exists
        //TODO: hardcode company multisig wallet
        _grantRole(DEFAULT_ADMIN_ROLE, msg.sender);

        // Role admin graph
        _setRoleAdmin(GOVERNANCE_ADMIN_ROLE, GOVERNANCE_ADMIN_ROLE);
        _setRoleAdmin(ARBITER_ROLE, GOVERNANCE_ADMIN_ROLE);

        // In practice this should be a Safe multi-sig wallet
        for (uint256 i = 0; i < coldWalletAdmins.length; i++) {
            address admin = coldWalletAdmins[i];
            if (admin == address(0)) revert ZeroAddress();
            _grantRole(GOVERNANCE_ADMIN_ROLE, admin);
        }
    }

    // Arbiter list (Only changes via in practice multisig governance role)
    function addArbiter(address arbiter) external onlyRole(GOVERNANCE_ADMIN_ROLE) {
        if (arbiter == address(0)) revert ZeroAddress();
        _grantRole(ARBITER_ROLE, arbiter);
    }

    function removeArbiter(address arbiter) external onlyRole(GOVERNANCE_ADMIN_ROLE) {
        _revokeRole(ARBITER_ROLE, arbiter);
    }

    // Optional: manage cold-wallet role holders (self-admin)
    function addGovernanceAdmin(address admin) external onlyRole(GOVERNANCE_ADMIN_ROLE) {
        if (admin == address(0)) revert ZeroAddress();
        _grantRole(GOVERNANCE_ADMIN_ROLE, admin);

        // New you New age
        governanceAdminAddedAt[admin] = uint64(block.timestamp);
    }

    function removeGovernanceAdmin(address admin) external onlyRole(GOVERNANCE_ADMIN_ROLE) {
        uint256 count = getRoleMemberCount(GOVERNANCE_ADMIN_ROLE);
        if (count <= 1) revert CannotRemoveLastGovernanceAdmin();

        // If this removal leaves exactly one admin, the remaining admin must be old enough.
        if (count == 2) {
            address m0 = getRoleMember(GOVERNANCE_ADMIN_ROLE, 0);
            address m1 = getRoleMember(GOVERNANCE_ADMIN_ROLE, 1);
            address remaining = (admin == m0) ? m1 : m0;

            uint64 addedAt = governanceAdminAddedAt[remaining];
            uint64 nowTs   = uint64(block.timestamp);

            if (nowTs - addedAt < MIN_SOLE_ADMIN_AGE) {
                revert RemainingAdminTooNew(addedAt, nowTs, MIN_SOLE_ADMIN_AGE);
            }
        }

        _revokeRole(GOVERNANCE_ADMIN_ROLE, admin);
    }
}