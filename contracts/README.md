# Sample Hardhat Project
### Generate Junk wallet for testing
```shell
node -e "const wallet = require('ethers').Wallet.createRandom(); console.log('Address:', wallet.address); console.log('Private Key:', wallet.privateKey)"
```


### Commands 
```shell
npx hardhat help
npx hardhat test
npx hardhat node
npx hardhat ignition deploy ./ignition/modules/SettleGuard.js
```
### Sample Environment file
```dotenv
# --- WALLET & PROVIDER ---
# Your wallet private key NO 0x
PRIVATE_KEY="NopeNotGonnaHappen"

# Your RPC URL (e.g., Alchemy or Infura Sepolia endpoint)
SEPOLIA_RPC_URL="IamBenderPleaseInsertGirder"

# This is mocked and not used anymore okay meow?
# Sepolia USDC: 0x1c7D4B196Cb0232b30447a592622705109140736
USDC_ADDRESS="0x1c7D4B196Cb0232b30447a592622705109140736"

# Should be the multisig address, however for testing YOLO
FEE_RECIPIENT=""

# Initial Governance Admins. Remember to use a comma for protection and multiple wallets
INITIAL_ADMINS=""

# Minimum age for a sole admin. Deploy = 604800 = 7 days
MIN_SOLE_ADMIN_AGE=0

# VERIFICATION
ETHERSCAN_API_KEY="InsertKey"
```

### Deploy & Export
```shell
npx hardhat ignition deploy ignition/modules/SettleGuard.js --network sepolia --parameters ignition/parameters.json && npx hardhat run scripts/exportAddresses.js --network sepolia
```
### Deploy Contacts
```shell
npx hardhat ignition deploy ignition/modules/SettleGuard.js --network sepolia --parameters ignition/parameters.json
```
### Export Addresses for API
```shell
npx hardhat run scripts/exportAddresses.js --network sepolia
```
### Run Once
Deposits fake USDC into wallets for admins {INITIAL_ADMINS}
Grants {INITIAL_ADMINS} 
```shell
npx hardhat run scripts/setupEnvironment.js --network sepolia
```

### Run Categories
```shell
npx hardhat run scripts/seedCategories.js --network sepolia
```
