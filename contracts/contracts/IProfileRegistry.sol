// SPDX-License-Identifier: MIT
pragma solidity ^0.8.33;

interface IProfileRegistry {
    function getScore(address _user) external view returns (uint256);
}