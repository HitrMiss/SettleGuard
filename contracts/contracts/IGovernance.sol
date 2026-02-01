// SPDX-License-Identifier: MIT
pragma solidity ^0.8.33;

interface IGovernance {
    function GOVERNANCE_ADMIN_ROLE() external view returns (bytes32);
    function ARBITER_ROLE() external view returns (bytes32);

    function hasRole(bytes32 role, address account) external view returns (bool);
    function MIN_SOLE_ADMIN_AGE() external view returns (uint64);

    function getAdminAddedAt(address admin) external view returns (uint64);
}