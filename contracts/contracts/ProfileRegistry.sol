// SPDX-License-Identifier: MIT
pragma solidity ^0.8.33;

import "./IGovernance.sol";
//TODO: DO NOT USE, I REALLY NEED SLEEP
contract ProfileRegistry {
    IGovernance public immutable governance;

    //TODO: Check the gas add up
    struct WalletProfile {
        uint8 trustScore;            // 0-10 (0=Noid, 10=Little Caesar) (0=Hamburglar, 10=Ronald McDonald.)
        uint64 lastReversalAt;       // Timestamp of the last bad behavior
        uint256 totalVolume;         // Lifetime successful throughput
        uint256 totalReversedVolume; // The "Bad Debt" accumulated
    }

    error Unauthorized();

    mapping(address => WalletProfile) public profiles;

    constructor(address _gov) { governance = IGovernance(_gov); }

    function updateScore(address _user, uint8 _score) external {
        if (!governance.hasRole(governance.ARBITER_ROLE(), msg.sender)) revert Unauthorized();
        profiles[_user].trustScore = _score;
    }
}