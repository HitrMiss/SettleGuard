const hre = require("hardhat");
const fs = require("fs");

async function main() {
    const configPath = "./backend_config.json";
    const config = JSON.parse(fs.readFileSync(configPath, "utf8"));
    const [admin] = await hre.ethers.getSigners();

    const registryAddress = config.contracts["SettleGuard#CategoryRegistry"];

    if (!registryAddress) {
        throw new Error("❌ Could not find SettleGuard#CategoryRegistry in backend_config.json");
    }

    console.log("Found CategoryRegistry at:", registryAddress);

    const registry = await hre.ethers.getContractAt("CategoryRegistry", registryAddress);

    const categories = [
        { name: "DIGITAL_GOODS", multiplier: 11000, buffer: 60, min: 300, max: 3600 },
        { name: "FOOD_BEVERAGE", multiplier: 10000, buffer: 30, min: 30, max: 720 }
    ];

    const ids = [];
    const rules = [];
    const goMapping = {};

    categories.forEach(c => {
        const id = hre.ethers.encodeBytes32String(c.name);
        ids.push(id);
        rules.push({
            existence: true,
            enabled: true,
            multiplierBps: c.multiplier,
            bufferSecs: c.buffer,
            minHoldSecs: c.min,
            maxHoldSecs: c.max,
            formulaParam: 0
        });
        goMapping[c.name] = id;
    });

    console.log("Seeding categories...");
    const tx = await registry.batchUpsertRules(ids, rules);
    await tx.wait();

    const categoryMapping = {};
    categories.forEach(c => {
        categoryMapping[c.name] = hre.ethers.encodeBytes32String(c.name);
    });

    config.categoryIds = categoryMapping;

    fs.writeFileSync(configPath, JSON.stringify(config, null, 2));
    console.log("✅ Successfully wrote Category IDs to backend_config.json");
}

main().catch(console.error);