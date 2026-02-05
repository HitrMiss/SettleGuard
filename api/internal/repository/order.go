package repository

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

// OrderRecord Generated, not tested, I put the Z in lazy
type OrderRecord struct {
	PacketID string      `json:"packetId"`
	Cart     interface{} `json:"cart"`
	Amount   string      `json:"amount"`
	Status   string      `json:"status"`
}

// SaveOrderRecord writes orders to disk, I know I know but all IO is IO?
func SaveOrderRecord(packetID string, cart interface{}, amount string, createdAt uint64) error {
	dir := "./data/orders"

	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("could not create storage directory: %w", err)
	}

	record := map[string]interface{}{
		"packetId":  packetID,
		"cart":      cart,
		"amount":    amount,
		"createdAt": createdAt,
		"status":    "pending",
	}

	filePath := filepath.Join(dir, packetID+".json")
	data, err := json.MarshalIndent(record, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal order: %w", err)
	}

	return os.WriteFile(filePath, data, 0644)
}
