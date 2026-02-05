// SPDX-License-Identifier: MIT
pragma solidity ^0.8.33;

import "./IPaymentVault.sol";
import "./ICategoryRegistry.sol";
import "./IProfileRegistry.sol";
import "./IGovernance.sol";
import "./IBondVault.sol";
import "@openzeppelin/contracts/utils/ReentrancyGuard.sol";

contract SettlementEngine is ReentrancyGuard {
    IPaymentVault public immutable vault;
    ICategoryRegistry public categories;
    IProfileRegistry public profiles;
    IBondVault public immutable bondVault;
    IGovernance public immutable governance;

    error PreflightFailed(string reason);
    error Unauthorized();

    constructor(address _vault, address _categories, address _bondVault, address _profiles, address _gov) {
        vault = IPaymentVault(_vault);
        categories = ICategoryRegistry(_categories);
        bondVault = IBondVault(_bondVault);
        profiles = IProfileRegistry(_profiles);
        governance = IGovernance(_gov);
    }

    modifier onlyArbiter() {
        if (!governance.hasRole(governance.ARBITER_ROLE(), msg.sender)) {
            revert Unauthorized();
        }
        _;
    }

    function updateRegistries(address _categories, address _profiles) external {
        if (!governance.hasRole(governance.GOVERNANCE_ADMIN_ROLE(), msg.sender)) revert Unauthorized();
        categories = ICategoryRegistry(_categories);
        profiles = IProfileRegistry(_profiles);
    }


    function deposit(
        uint256 _amount,
        address _merchantIdentity,
        address _merchantPayout,
        bytes32 _categoryId,
        bytes32 _packetId
    ) external nonReentrant {
        // merchant
        if (profiles.getScore(_merchantIdentity) == 0) {
            bondVault.lockRisk(_merchantIdentity, _amount);
        }
        // customer
        if (profiles.getScore(msg.sender) == 0) {
            bondVault.lockRisk(msg.sender, _amount);
        }

        vault.lockFunds(msg.sender, _amount, _merchantIdentity, _merchantPayout, _categoryId, _packetId);
    }

    function triggerSettlement(bytes32 _packetId) external onlyArbiter {
        // Destructure the tuple from the vault
        (
            address payer,
            uint256 amount,
            , // createdAt
            address merchantIdentity,
            , // categoryId
            address merchantPayout,
            IPaymentVault.Status status
        ) = vault.getPaymentDetails(_packetId);

        if (profiles.getScore(merchantIdentity) == 0) {
            bondVault.unlockRisk(merchantIdentity, amount);
        }
        if (profiles.getScore(payer) == 0) {
            bondVault.unlockRisk(payer, amount);
        }
        bytes32[] memory ids = new bytes32[](1);
        ids[0] = _packetId;
        vault.prepareSettlementBulk(ids, merchantPayout);
    }
}