package main

import (
	"SettleGuard/internal/api"
	"SettleGuard/internal/config"
	"SettleGuard/internal/contract"
	"SettleGuard/internal/provider"
	"SettleGuard/internal/repository"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/ethereum/go-ethereum/common"
)

const (
	VaultKey       = "SettleGuard#PaymentVault"
	RegistryAddr   = "SettleGuard#CategoryRegistry"
	SettlementAddr = "0x"
	USDCAddr       = "0x"
)

func main() {
	// Start me up
	cfg, err := config.LoadConfig("cmd/api/backend_config.json")
	if err != nil {
		panic(err)
	}

	VaultAddr := cfg.Contracts[VaultKey]
	registryAddr := cfg.Contracts["SettleGuard#CategoryRegistry"]

	// Pulling your Category IDs
	digitalGoodsID := cfg.CategoryIDs["DIGITAL_GOODS"]

	fmt.Printf("Connected to %s\n", cfg.Network)
	fmt.Printf("Registry Address: %s\n", registryAddr)
	fmt.Printf("Digital Goods ID: %s\n", digitalGoodsID)

	p, err := provider.NewProvider()
	if err != nil {
		log.Fatalf("🚨 Provider Error: %v", err)
	}

	ethClient := p.Client

	// Auth Routes
	http.HandleFunc("/api/get-nonce", enableCORS(api.HandleGetNonce))
	http.HandleFunc("/api/verify", enableCORS(api.HandleVerifySignature))

	http.HandleFunc("/api/checkout", enableCORS(func(w http.ResponseWriter, r *http.Request) {
		api.HandleCheckout(
			w,
			r,
			ethClient,
			common.HexToAddress(RegistryAddr),
			common.HexToAddress(SettlementAddr),
			common.HexToAddress(USDCAddr),
		)
	}))

	// History Route
	http.HandleFunc("/api/history", enableCORS(func(w http.ResponseWriter, r *http.Request) {
		address := r.URL.Query().Get("address")

		vault, err := contract.NewPaymentVault(common.HexToAddress(VaultAddr), ethClient)
		if err != nil {
			http.Error(w, "Failed to connect to Vault contract", http.StatusInternalServerError)
			return
		}

		history, err := repository.GetUserHistory(vault, address)
		if err != nil {
			http.Error(w, "Failed to retrieve history", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(history)
	}))

	port := ":8080"
	fmt.Printf("🚀 (to the moon) SettleGuard API Server running on http://localhost%s\n", port)
	fmt.Printf("🛒 (need a quarter?) Checkout endpoint ready for demo site: http://localhost%s/api/checkout\n", port)
	fmt.Printf("⛓️  Target Blockchain: %s\n", os.Getenv("BLOCKCHAIN"))
	fmt.Printf("📡 RPC URL: %s\n", p.MaskedURL())
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal(err)
	}
}

// enableCORS: Because browsers are picky eaters
func enableCORS(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, X-User-Address") // Added X-User-Address for auth

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next(w, r)
	}
}
