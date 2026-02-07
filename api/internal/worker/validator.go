package worker

import (
	"context"
	"crypto/sha256"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"SettleGuard/internal/contract"
	"SettleGuard/internal/provider"
	"SettleGuard/internal/verifier"

	"github.com/ethereum/go-ethereum/common"
)

type ValidationWorker struct {
	VaultAddress common.Address
	Provider     *provider.Provider
	StoragePath  string // e.g., "./storage/orders" Get Shorty https://www.youtube.com/watch?v=HHSAml1BAR4
}

func (w *ValidationWorker) Start(ctx context.Context) {
	ticker := time.NewTicker(20 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			w.processDiskQueue()
		}
	}
}

func (w *ValidationWorker) processDiskQueue() {
	vault, err := contract.NewPaymentVault(w.VaultAddress, w.Provider.Client)
	if err != nil {
		log.Printf("Failed to bind vault: %v", err)
		return
	}

	files, _ := filepath.Glob(filepath.Join(w.StoragePath, "*.json"))
	for _, file := range files {
		packetIdStr := filepath.Base(file)
		packetIdStr = packetIdStr[:len(packetIdStr)-len(filepath.Ext(packetIdStr))] // remove .json
		packetId := common.HexToHash(packetIdStr)

		// Query getPaymentDetails returns the struct
		details, err := vault.GetPaymentDetails(nil, packetId)
		if err != nil {
			log.Printf("Failed to fetch on-chain details for %s: %v", packetIdStr, err)
			continue
		}

		if details.Status != 1 {
			continue
		}

		pubKeyHex, err := w.Provider.GetMerchantPubKeyByAddress(details.MerchantIdentity)
		if err != nil {
			log.Printf("ENS Resolution failed for merchant %s: %v", details.MerchantIdentity.Hex(), err)
			continue
		}

		// Reconstruct Manifest
		manifest := fmt.Sprintf("%s:%s:%s:%s:%d",
			details.MerchantIdentity.Hex(),
			details.MerchantPayout.Hex(),
			details.Amount.String(),
			common.BytesToHash(details.CategoryId[:]).Hex(),
			details.CreatedAt,
		)
		hash := sha256.Sum256([]byte(manifest))

		isValid, err := verifier.VerifyP256Signature(pubKeyHex, hash[:], details.R, details.S)

		if err != nil || !isValid {
			log.Printf("❌ SIGNATURE FRAUD: Packet %s", packetIdStr)
			w.moveToFolder(file, "failed", "Invalid Signature / Fraud Detected")
		} else {
			log.Printf("✅ VALIDATED: Packet %s", packetIdStr)
			// TODO: Add the code
		}
	}
}

// moveToFolder shifts the file to a sub-directory and logs the reason
func (w *ValidationWorker) moveToFolder(originalPath, subFolder, reason string) {
	fileName := filepath.Base(originalPath)
	newPath := filepath.Join(w.StoragePath, subFolder, fileName)

	err := os.Rename(originalPath, newPath)
	if err != nil {
		log.Printf("❌ Failed to move file %s: %v", fileName, err)
		return
	}

	if subFolder == "failed" {
		log.Printf("🚨 REJECTED: %s | Reason: %s", fileName, reason)
	} else {
		log.Printf("✅ SUCCESS: %s | %s", fileName, reason)
	}
}
