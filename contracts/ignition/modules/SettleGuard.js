const { buildModule } = require("@nomicfoundation/hardhat-ignition/modules");

module.exports = buildModule("SettleGuard", (m) => {
    const deployer = m.getAccount(0);
    const MIN_SOLE_ADMIN_AGE = m.getParameter("minAdminAge", 0n);
    const admins = m.getParameter("adminList");

    //Mock USDC
    const mockUsdc = m.contract("MockUSDC");

    // Core Governance
    const governance =
        m.contract("Governance",
            [admins, MIN_SOLE_ADMIN_AGE]);

    // Deploy Registries
    const profileRegistry =
        m.contract("ProfileRegistry",
            [governance]);
    const categoryRegistry =
        m.contract("CategoryRegistry",
            [governance]);

    // Deploy Vaults
    const bondVault = m.contract("BondVault", [
        governance,
        mockUsdc,          // Dynamic address from the deployed MockUSDC
        deployer,
        50n,
        1000000n
    ]);

    const paymentVault = m.contract("PaymentVault", [
        mockUsdc,          // Dynamic address from the deployed MockUSDC
        governance
    ]);

    // SettlementEngine
    const settlementEngine = m.contract("SettlementEngine", [
        paymentVault,
        categoryRegistry,
        bondVault,
        profileRegistry,
        governance
    ]);

    //Linking & Role Setup
    m.call(bondVault, "setSettlementEngine", [settlementEngine]);
    m.call(paymentVault, "setEngine", [settlementEngine]);

    // Preset trust score
    m.call(profileRegistry, "setScore", [deployer, 10]);

    return {
        mockUsdc,
        governance,
        profileRegistry,
        categoryRegistry,
        bondVault,
        paymentVault,
        settlementEngine
    };
});