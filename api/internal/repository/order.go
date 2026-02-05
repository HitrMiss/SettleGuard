package repository

import (
	"encoding/json"
	"os"
	"path/filepath"
)

// OrderRecord is what we save to the local disk
type OrderRecord struct {
	PacketID   string      `json:"packetId"`
	Payer      string      `json:"payer"`
	Items      interface{} `json:"items"`
	Amount     string      `json:"amount"`
	CreatedAt  uint64      `json:"createdAt"`
	CategoryID string      `json:"categoryId"`
	Status     string      `json:"status"`
}

type CartItem struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Price       float64 `json:"price"`
	Quantity    int     `json:"quantity"`
	Description string  `json:"description"`
	Image       string  `json:"image"`
}

// UIOrder is what we send to the frontend demo site
type UIOrder struct {
	PacketID  string      `json:"packetId"`
	Amount    string      `json:"amount"`
	CreatedAt uint64      `json:"createdAt"`
	Status    string      `json:"status"`
	Items     interface{} `json:"items"`
}

// SaveOrderRecord writes orders to disk, I know I know but all IO is IO?
func SaveOrderRecord(
	packetID string,
	payer string,
	items interface{},
	amount string,
	createdAt uint64,
	catID string, // The bytes32 hex
) error {
	dir := "./data/orders"
	os.MkdirAll(dir, 0755)

	record := OrderRecord{
		PacketID:   packetID,
		Payer:      payer,
		Items:      items,
		Amount:     amount,
		CreatedAt:  createdAt,
		CategoryID: catID,
		Status:     "pending",
	}

	filePath := filepath.Join(dir, packetID+".json")
	data, _ := json.MarshalIndent(record, "", "  ")
	return os.WriteFile(filePath, data, 0644)
}
