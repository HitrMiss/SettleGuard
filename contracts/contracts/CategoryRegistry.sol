// SPDX-License-Identifier: MIT

pragma solidity ^0.8.33;

import "./Governance.sol";

contract CategoryRegistry {
    IGovernance public immutable governance;

    struct CategoryRule {
        // aligned for storage
        bool enabled;               // Add only, disable do not remove
        uint16 multiplierBps;       // Basis points; Examples: 10000 = 1x, 20000 = 2x
        uint32 minHoldSecs;         // Minimum hold time; e.g 30 minutes for food (not less)
        uint32 maxHoldSecs;         // Maximum hold time; e.g 12 hours for food (catering? extreme farm to table?)
        uint32 bufferSecs;          // grace period
        uint32 formulaParam;        // Not used at release, added since we have the space
    }

    mapping(bytes32 => CategoryRule) private _rules;
    bytes32[] public categoryList;

    event RuleUpdated(bytes32 indexed id, uint16 multiplier, uint32 buffer);

    error Unauthorized();
    error CategoryDisabled();

    constructor(address _gov) {
        governance = IGovernance(_gov);
    }

    function _validate(CategoryRule calldata r) internal pure {
        if (r.multiplierBps == 0) revert BadRule();
        if (r.minHoldSecs == 0) revert BadRule();
        if (r.maxHoldSecs < r.minHoldSecs) revert BadRule();
    }

    /**
     * @notice Calculates final window with overflow protection and strict clamping.
     * @dev Uses Basis Points (BPS) for precision math.
     * Formula: Clamp((Expected * Multiplier) + Buffer, Min, Max)
     */
    function calculateExecutionWindow(bytes32 _id, uint32 _expectedSecs)
    external
    view
    returns (uint32)
    {
        CategoryRule storage r = _rules[_id];
        if (!r.enabled) revert CategoryDisabled();
        // Technically an overflow would revert anyways however added to avoid possible DoS
        uint256 calculated = (uint256(_expectedSecs) * uint256(r.multiplierBps)) / 10000;

        calculated += uint256(r.bufferSecs);

        if (calculated < uint256(r.minHoldSecs)) {
            return r.minHoldSecs;
        }

        if (calculated > uint256(r.maxHoldSecs)) {
            return r.maxHoldSecs;
        }

        return uint32(calculated);
    }

    function upsertRule(bytes32 _id, CategoryRule calldata _rule) external {
        if (!governance.hasRole(governance.GOVERNANCE_ADMIN_ROLE(), msg.sender)) {
            revert Unauthorized();
        }

        if (_id == bytes32(0))
        {
            revert ZeroId();
        }

        _validate(_rule);

        if (!_rules[_id]) {
            _rules[_id] = true;
            categoryList.push(_id);
        }

        _rules[_id] = _rule;
        emit RuleUpdated(_id, _rule.multiplierBps, _rule.bufferSecs);
    }

    function getRule(bytes32 _id) external view returns (CategoryRule memory) {
        return _rules[_id];
    }
}