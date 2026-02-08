package repository

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"

	"SettleGuard/internal/contract"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

func GetUserHistory(vault *contract.PaymentVault, userAddress string) ([]UIOrder, error) {
	var history []UIOrder
	files, _ := os.ReadDir("./data/orders")

	for _, file := range files {
		if filepath.Ext(file.Name()) != ".json" {
			continue
		}

		data, err := os.ReadFile(filepath.Join("./data/orders", file.Name()))
		if err != nil {
			continue
		}

		var record OrderRecord
		json.Unmarshal(data, &record)

		if strings.EqualFold(record.Payer, userAddress) {
			var pID [32]byte
			copy(pID[:], common.FromHex(record.PacketID))

			// Check current status on-chain
			details, err := vault.GetPaymentDetails(&bind.CallOpts{}, pID)
			statusStr := "PENDING_ON_CHAIN"
			if err == nil {
				statusStr = mapStatus(details.Status)
			}

			history = append(history, UIOrder{
				PacketID:  record.PacketID,
				Amount:    record.Amount,
				CreatedAt: record.CreatedAt,
				Status:    statusStr,
				Items:     record.Items,
			})
		}
	}
	return history, nil
}

func mapStatus(s uint8) string {
	statuses := []string{"NONE", "HELD", "RELEASED", "REFUNDED"}
	if int(s) >= len(statuses) {
		return "UNKNOWN"
	}
	return statuses[s]
}
