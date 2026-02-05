require("@nomicfoundation/hardhat-toolbox");
require("dotenv").config();

const rawAdmins = process.env.INITIAL_ADMINS ? process.env.INITIAL_ADMINS.split(",") : [];

const sanitizedAdmins = rawAdmins
    .map(a => a.trim())
    .filter(a => a.startsWith("0x"));

module.exports = {
  solidity: "0.8.33",
  networks: {
    sepolia: {
      url: process.env.SEPOLIA_RPC_URL,
      accounts: process.env.PRIVATE_KEY
    }
  },
  // define Params
  ignition: {
    parameters: {
      SettleGuardModule: {
        adminList: sanitizedAdmins,
        minAdminAge: 0n
      }
    }
  }
};