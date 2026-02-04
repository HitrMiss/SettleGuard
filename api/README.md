# Instructions
## Environment Setup

### .ENV file
Create a new file called ".env" in root of this folder  
File contents: 
```
ALCHEMY_API_KEY={YOUR DEV KEY}  
BLOCKCHAIN=eth-mainnet
```
"BLOCKCHAIN" values used by project currently
```
eth-sepolia       - TestNet
eth-mainnet       - Ethereum Mainnet
polygon-amoy      - Polygon Mainnet
polygon-mainnet   - Mainnet
```

### Linux
If eth is not installed  
`sudo scripts/install_solc.sh`
### Windows
If eth is not installed  
`.\scripts\install_solc.ps1`

## Generating ABI GO files
Run from local terminal or shell from the API folder  
`go generate ./internal/contract/...`