package main

import (
	"SettleGuard/internal/config"
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"SettleGuard/internal/provider"
	"SettleGuard/internal/worker"

	"github.com/ethereum/go-ethereum/common"
)

const (
	VaultKey      = "SettleGuard#PaymentVault"
	RegistryKey   = "SettleGuard#CategoryRegistry"
	SettlementKey = "SettleGuard#SettlementEngine"
	USDCKey       = "SettleGuard#MockUSDC"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	cfg, err := config.LoadConfig("cmd/api/backend_config.json")
	if err != nil {
		panic(err)
	}

	VaultAddr := cfg.Contracts[VaultKey]
	//RegistryAddr := cfg.Contracts[RegistryKey]
	//SettlementAddr := cfg.Contracts[SettlementKey]
	//USDCAddr := cfg.Contracts[USDCKey]

	blockchainProvider, err := provider.NewProvider()
	if err != nil {
		log.Fatalf("🚨 Worker failed to initialize blockchain provider: %v", err)
	}
	log.Printf("Connected to RPC: %s", blockchainProvider.MaskedURL())

	vaultAddress := common.HexToAddress(VaultAddr)

	settlementWorker := &worker.ValidationWorker{
		VaultAddress: vaultAddress,
		Provider:     blockchainProvider,
		StoragePath:  "./data/orders",
	}

	log.Printf("🚀 Settlement Worker active")
	log.Printf("📁 Monitoring: %s", settlementWorker.StoragePath)
	log.Printf("🔗 Targeting Vault: %s", vaultAddress.Hex())

	settlementWorker.Start(ctx)

	log.Println("👋 Worker shut down gracefully")
}
