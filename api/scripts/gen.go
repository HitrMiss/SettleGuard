//go:build ignore

package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
)

// Tested on Windows have not yet tested on Linux my normal OS
func main() {
	_, filename, _, _ := runtime.Caller(0)
	scriptDir := filepath.Dir(filename)
	apiDir := filepath.Join(scriptDir, "..")

	contractsRoot := filepath.Join(scriptDir, "..", "contracts") // Points to /contracts
	apiDir = filepath.Join(scriptDir, "..")                      // Points to /api
	outDir := filepath.Join(apiDir, "internal", "contract")

	configFileName := "backend_config.json"
	srcConfig := filepath.Join(contractsRoot, configFileName)
	dstConfig := filepath.Join(outDir, configFileName)

	if _, err := os.Stat(srcConfig); err == nil {
		fmt.Printf("[+] Syncing %s to API internal...\n", configFileName)

		// Read source
		input, err := os.ReadFile(srcConfig)
		if err != nil {
			fmt.Printf("    - Error reading config: %v\n", err)
		} else {
			// Write to destination
			err = os.WriteFile(dstConfig, input, 0644)
			if err != nil {
				fmt.Printf("    - Error writing config: %v\n", err)
			} else {
				fmt.Println("    - Sync Success.")
			}
		}
	} else {
		fmt.Printf("[!] Warning: %s not found in contracts root. Deploy first!\n", configFileName)
	}

	solcName := "solc"
	if runtime.GOOS == "windows" {
		solcName = "solc.exe"
	}
	compilerPath := filepath.Join(apiDir, "bin", solcName)
	contractsRoot = filepath.Join(apiDir, "..", "contracts")
	solSourceDir := filepath.Join(contractsRoot, "contracts")
	outDir = filepath.Join(apiDir, "internal", "contract")

	defer cleanupArtifacts(outDir)

	if _, err := os.Stat(compilerPath); os.IsNotExist(err) {
		fmt.Println("Error: local solc not found. Run install_solc.ps1 first.")
		return
	}

	files, _ := filepath.Glob(filepath.Join(solSourceDir, "*.sol"))

	for _, file := range files {
		contractName := strings.TrimSuffix(filepath.Base(file), ".sol")
		if len(contractName) > 1 && contractName[0] == 'I' && contractName[1] >= 'A' && contractName[1] <= 'Z' {
			continue
		}

		fmt.Printf("[+] Compiling & Binding: %s\n", contractName)

		// A. Compile (Working directory set to 'contracts')
		relFile := filepath.Join("contracts", filepath.Base(file))
		compCmd := exec.Command(compilerPath,
			"--abi", "--bin", "--overwrite",
			"--base-path", ".",
			"--include-path", "node_modules",
			"@openzeppelin/=node_modules/@openzeppelin/",
			relFile,
			"-o", filepath.Join("..", "api", "internal", "contract"),
		)
		compCmd.Dir = contractsRoot

		if out, err := compCmd.CombinedOutput(); err != nil {
			fmt.Printf("    - Solc Error: %s\n", string(out))
			continue
		}

		abiPath := filepath.Join(outDir, contractName+".abi")
		binPath := filepath.Join(outDir, contractName+".bin")
		goOut := filepath.Join(outDir, strings.ToLower(contractName)+".go")

		abigenCmd := exec.Command("go", "run", "github.com/ethereum/go-ethereum/cmd/abigen",
			"--abi", abiPath, "--bin", binPath, "--pkg", "contract", "--type", contractName, "--out", goOut)

		if out, err := abigenCmd.CombinedOutput(); err != nil {
			fmt.Printf("    - Abigen Error: %s\n", string(out))
		} else {
			fmt.Printf("    - Success: %s\n", goOut)
		}
	}
}

// Clean up, clean up. Everybody, let's clean up. Clean up, clean up. Put your things away.
func cleanupArtifacts(dir string) {
	fmt.Println("[*] Cleaning up temporary ABI/BIN artifacts...")
	extensions := []string{"*.abi", "*.bin"}

	for _, ext := range extensions {
		matches, _ := filepath.Glob(filepath.Join(dir, ext))
		for _, match := range matches {
			if err := os.Remove(match); err != nil {
				fmt.Printf("    - Failed to remove %s: %v\n", filepath.Base(match), err)
			}
		}
	}
}
