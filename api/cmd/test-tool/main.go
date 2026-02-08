package main

import (
	"SettleGuard/internal/config"
	"SettleGuard/internal/contract"
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"SettleGuard/internal/provider"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

// I need to stop copying and pasting this
const (
	VaultKey      = "SettleGuard#PaymentVault"
	RegistryKey   = "SettleGuard#CategoryRegistry"
	SettlementKey = "SettleGuard#SettlementEngine"
	USDCKey       = "SettleGuard#MockUSDC"
)

// This class has been generated with the assistance of AI
func main() {
	network := "sepolia" // Default
	p, err := provider.NewProvider()
	if err != nil {
		log.Fatalf("[-] Connection failed: %v", err)
	}

	cfg, err := config.LoadConfig("cmd/api/backend_config.json")
	if err != nil {
		panic(err)
	}

	VaultAddr := cfg.Contracts[VaultKey]
	//RegistryAddr := cfg.Contracts[RegistryKey]
	//SettlementAddr := cfg.Contracts[SettlementKey]
	//USDCAddr := cfg.Contracts[USDCKey]

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\n--- SettleGuard Sanity Tool ---")
		fmt.Printf("Connected to: %s [%s]\n", p.URL, network)
		fmt.Println("1. Test ENS Merchant Lookup")
		fmt.Printf("2. Check BondVault Balance (TODO)\n")
		fmt.Printf("3. Get Payment Packet\n")
		fmt.Println("4. Generate Merchant Keys")
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
			fmt.Println("Bound Vault soon!)...")

		case "3":
			fmt.Print("Enter Packet ID: ")
			packetIdStr, _ := reader.ReadString('\n')
			packetIdStr = strings.TrimSpace(packetIdStr)
			packetId := common.HexToHash(packetIdStr)

			vaultAddrHex := common.HexToAddress(VaultAddr)

			// Bind the contract
			vault, err := contract.NewPaymentVault(vaultAddrHex, p.Client)
			if err != nil {
				fmt.Printf("❌ Failed to bind Vault: %v\n", err)
				break
			}

			fmt.Printf("[*] Querying Vault at %s...\n", VaultAddr)
			details, err := vault.GetPaymentDetails(nil, packetId)
			if err != nil {
				fmt.Printf("❌ Call failed: %v (Is the Packet ID correct?)\n", err)
				break
			}

			// Display the Data
			fmt.Println("\n--- On-Chain Payment Record ---")
			fmt.Printf("  Payer:      %s\n", details.Payer.Hex())
			fmt.Printf("  Identity:   %s\n", details.MerchantIdentity.Hex())
			fmt.Printf("  Payout:     %s\n", details.MerchantPayout.Hex())
			fmt.Printf("  Amount:     %s units\n", details.Amount.String())
			fmt.Printf("  Category:   %s\n", common.BytesToHash(details.CategoryId[:]).Hex())
			fmt.Printf("  Timestamp:  %d\n", details.CreatedAt)
			fmt.Printf("  Status:     %d (0:None, 1:Held, 2:Released, 3:Refunded)\n", details.Status)
			fmt.Printf("  Signature:  R: %x\n", details.R)
			fmt.Printf("              S: %x\n", details.S)

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
