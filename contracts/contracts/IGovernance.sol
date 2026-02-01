// SPDX-License-Identifier: MIT
pragma solidity ^0.8.33;

interface IGovernance {
    function GOVERNANCE_ADMIN_ROLE() external view returns (bytes32);
    function ARBITER_ROLE() external view returns (bytes32);

    function hasRole(bytes32 _role, address _account) external view returns (bool);
}