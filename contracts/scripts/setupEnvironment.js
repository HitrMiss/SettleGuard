const fs = require("fs");
const hre = require("hardhat");
require("dotenv").config();

async function main() {
    const configPath = "./backend_config.json";
    if (!fs.existsSync(configPath)) {
        throw new Error("❌ backend_config.json not found. Run exportAddresses.js first!");
    }
    const { contracts } = JSON.parse(fs.readFileSync(configPath, "utf8"));

    const adminAddresses = process.env.INITIAL_ADMINS
        .split(",")
        .map(a => a.trim())
        .filter(a => a !== "");

    const mockUsdc = await hre.ethers.getContractAt("MockUSDC", contracts["SettleGuard#MockUSDC"]);
    const profileRegistry = await hre.ethers.getContractAt("ProfileRegistry", contracts["SettleGuard#ProfileRegistry"]);

    console.log("🚀 Starting environment setup...");

    for (const admin of adminAddresses) {
        console.log(`\nProcessing Admin: ${admin}`);

        const mintTx = await mockUsdc.mint(admin, 1000_000000n);
        await mintTx.wait();
        console.log(`  ✅ Minted 1k USDC`);

        const scoreTx = await profileRegistry.setScore(admin, 10);
        await scoreTx.wait();
        console.log(`  ✅ Trust Score set to 10`);
    }

    console.log("\n🎉 Setup complete. Admins are funded and trusted!");
}

main().catch((error) => {
    console.error(error);
    process.exit(1);
});