// SPDX-License-Identifier: MIT
pragma solidity ^0.8.33;

interface ICategoryRegistry {
    struct CategoryRule {
        // aligned for storage
        bool existence;             // If only life was this easy to learn of existence
        bool enabled;               // Add only, disable do not remove
        uint16 multiplierBps;       // Basis points; Examples: 10000 = 1x, 20000 = 2x
        uint32 minHoldSecs;         // Minimum hold time; e.g 30 minutes for food (not less)
        uint32 maxHoldSecs;         // Maximum hold time; e.g 12 hours for food (catering? extreme farm to table?)
        uint32 bufferSecs;          // grace period, in case ❄️ Fern ❄️ comes back
        uint32 formulaParam;        // Not used at release, added since we have the space
    }
    function getRule(bytes32 _id) external view returns (CategoryRule memory);
    function calculateExecutionWindow(bytes32 _id, uint32 _expectedSecs) external view returns (uint32);
}