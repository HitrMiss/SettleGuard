package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"SettleGuard/api/internal/provider"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
)

// This class has been generated with the assistance of AI
func main() {
	// 1. Initial Setup
	network := "sepolia" // Default
	p, err := provider.NewProvider(network)
	if err != nil {
		log.Fatalf("[-] Connection failed: %v", err)
	}

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\n--- SettleGuard Sanity Tool ---")
		fmt.Printf("Connected to: %s [%s]\n", p.URL, network)
		fmt.Println("1. Test ENS Merchant Lookup")
		fmt.Printf("2. Check BondVault Balance (TODO)\n")
		fmt.Println("3. Change Network (Current: " + network + ")")
		fmt.Println("q. Quit")
		fmt.Print("Select an option: ")

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		switch input {
		case "1":
			fmt.Print("Enter ENS Name (e.g., merchant.eth): ")
			name, _ := reader.ReadString('\n')
			name = strings.TrimSpace(name)

			key, err := p.GetMerchantPubKey(name)
			if err != nil {
				fmt.Printf("❌ Error: %v\n", err)
			} else {
				fmt.Printf("✅ P-256 Public Key: %s\n", key)
			}

		case "2":
			fmt.Println("🏗️  BondVault logic is coming in the next file (internal/fraud/engine.go)...")

		case "3":
			fmt.Print("Enter network (mainnet/sepolia): ")
			net, _ := reader.ReadString('\n')
			network = strings.TrimSpace(net)
			p, _ = provider.NewProvider(network)
			fmt.Printf("🔄 Switched to %s\n", network)

			// Inside cmd/test-tool/main.go switch:

		case "4":
			fmt.Println("[*] Generating P-256 Identity...")
			priv, _ := provider.GenerateP256Key()

			// Save both files: merchant_test.pem and merchant_test.pub
			err := provider.SaveKeyPair("merchant_test", priv)
			if err != nil {
				fmt.Printf("❌ Error saving keys: %v\n", err)
				break
			}

			pubHex, _ := os.ReadFile("merchant_test.pub")
			fmt.Printf("💾 Keys saved: merchant_test.pem and merchant_test.pub\n")
			fmt.Printf("✨ Compressed Public Key: %s\n", string(pubHex))

		case "5": // Local Wallet & Faucet
			wallet, _ := provider.CreateWallet()
			fmt.Printf("📂 Local Eth Wallet Created!\nAddress: %s\n", wallet.Address.Hex())
			fmt.Printf("🔑 Private Key: %s\n", wallet.ExportHex())
			fmt.Printf("\n👉 Please go to a faucet (e.g., https://sepoliafaucet.com) and send funds to this address.\n")

		case "6": // ENS Update Logic
			fmt.Print("Enter ENS Domain: ")
			domain, _ := reader.ReadString('\n')
			domain = strings.TrimSpace(domain)

			fmt.Print("Enter Eth Private Key for Gas: ")
			pkHex, _ := reader.ReadString('\n')
			pkHex = strings.TrimSpace(pkHex)

			// Setup Auth
			privKey, _ := crypto.HexToECDSA(pkHex)
			chainID, _ := p.Client.ChainID(context.Background())
			auth, _ := bind.NewKeyedTransactorWithChainID(privKey, chainID)

			// Load the public key we generated in Case 4
			pubKeyBytes, _ := os.ReadFile("merchant_test.pub")

			fmt.Println("🚀 Sending ENS Update Transaction...")
			txHash, err := p.SetMerchantPubKey(auth, domain, string(pubKeyBytes))
			if err != nil {
				fmt.Printf("❌ TX Failed: %v\n", err)
			} else {
				fmt.Printf("✅ TX Sent! Hash: %s\n", txHash.Hex())
			}

		case "q":
			fmt.Println("Exiting...")
			return

		default:
			fmt.Println("Invalid selection, try again.")
		}
	}
}
