require("@nomicfoundation/hardhat-toolbox");
require("dotenv").config();
const fs = require("fs");

const rawAdmins = process.env.INITIAL_ADMINS ? process.env.INITIAL_ADMINS.split(",") : [];

const sanitizedAdmins = rawAdmins
    .map(a => a.trim())
    .filter(a => a.startsWith("0x"));

const params = {
  SettleGuard: {
    minAdminAge: 0,
    adminList: sanitizedAdmins
  }
};
fs.writeFileSync("ignition/parameters.json", JSON.stringify(params, null, 2));

module.exports = {
  solidity: "0.8.33",
  networks: {
    sepolia: {
      url: process.env.SEPOLIA_RPC_URL,
      accounts: process.env.PRIVATE_KEY ? [process.env.PRIVATE_KEY.trim()] : []
    }
  },
  verify: {
    etherscan: {
      apiKey: "process.env.ETHERSCAN_API_KEY",
    },
  }
};