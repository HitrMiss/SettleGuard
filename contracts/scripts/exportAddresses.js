const fs = require("fs");
const path = require("path");

async function main() {
    // Ignition stores deployment info in the ignition/deployments folder
    const deploymentPath = path.join(__dirname, "../ignition/deployments/chain-11155111/deployed_addresses.json");

    if (fs.existsSync(deploymentPath)) {
        const addresses = JSON.parse(fs.readFileSync(deploymentPath, "utf8"));

        // Format specifically for your Go backend config
        const goConfig = {
            network: "sepolia",
            contracts: addresses
        };

        fs.writeFileSync("backend_config.json", JSON.stringify(goConfig, null, 2));
        console.log("✅ Addresses exported to backend_config.json");
    }
}

main().catch((error) => {
    console.error(error);
    process.exit(1);
});