# Instructions
## Enviroment Setup
### Linux
If eth is not installed  
`sudo scripts/install_solc.sh`
### Windows
If eth is not installed  
`.\scripts\install_solc.ps1`

## Generating ABI GO files
Run from local terminal or shell from the API folder  
`go generate ./internal/contract/...`