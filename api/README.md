# Instructions
## Environment Setup

### .ENV file
Create a new file called ".env" in root of this folder  
File contents: 
```shell
ALCHEMY_API_KEY={YOUR DEV KEY}  
BLOCKCHAIN=eth-mainnet
```
"BLOCKCHAIN" values used by project currently
```
eth-sepolia       - TestNet
eth-mainnet       - Ethereum Mainnet
polygon-amoy      - Polygon Testnet
polygon-mainnet   - Polygon Mainnet
```

### Linux
If eth is not installed  
```shell
sudo scripts/install_solc.sh
```

### Windows
If eth is not installed  
```shell
.\scripts\install_solc.ps1
````

## Generating ABI GO files
Run from local terminal or shell from the API folder  
```shell
go generate ./internal/contract/...
```

## Building 
```shell
# Build the API
go build -o settleguard-api ./cmd/api

# Build the CLI
go build -o settleguard-cli ./cmd/test-tool
```
## Running
API Server  
```shell
go run ./cmd/api
```

CLI Test Tool
```shell
go run ./cmd/test-tool
```