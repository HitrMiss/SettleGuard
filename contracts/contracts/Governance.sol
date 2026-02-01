// SPDX-License-Identifier: MIT

pragma solidity ^0.8.33;

import "./IGovernance.sol";
import "@openzeppelin/contracts/access/extensions/AccessControlEnumerable.sol";
// Add baggage it's least it not emotional?
import "@openzeppelin/contracts/access/IAccessControl.sol";

contract Governance is IGovernance, AccessControlEnumerable {
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

    constructor(address[] memory _governanceAdmins, uint64 _minSoleAdminAgeSeconds) {
        if (_governanceAdmins.length == 0) revert EmptyGovernanceAdmins();

        // Default to 7 days if caller passes 0.
        uint64 minAge = _minSoleAdminAgeSeconds == 0 ? uint64(7 days) : _minSoleAdminAgeSeconds;
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
        for (uint256 i = 0; i < _governanceAdmins.length; i++) {
            address admin = _governanceAdmins[i];
            if (admin == address(0)) revert ZeroAddress();
            _grantRole(GOVERNANCE_ADMIN_ROLE, admin);
        }
    }

    // External Get-er Done
    function getAdminAddedAt(address _admin) external view returns (uint64) {
        return governanceAdminAddedAt[_admin];
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

    // Who needs sleep, never gonna get it
    // Full House was an okay show... now with interfaces *Spoke in Ron Popeil style delivery*
    // Wait theres more!!! Diamonds
    /**
     * @dev I blame Hard Hat!
     */
    function hasRole(bytes32 _role, address _account)
    public
    view
    virtual
    override(AccessControl, IAccessControl, IGovernance)
    returns (bool)
    {
        return super.hasRole(_role, _account);
    }
    /**
     * @dev See Above?
     */
    function supportsInterface(bytes4 _interfaceId)
    public
    view
    virtual
    override(AccessControlEnumerable)
    returns (bool)
    {
        return super.supportsInterface(_interfaceId);
    }

}