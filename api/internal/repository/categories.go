package repository

import (
	"SettleGuard/internal/contract"
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

// CategoryDetail Chain State
type CategoryDetail struct {
	ID            string `json:"id"`
	MultiplierBps uint16 `json:"multiplierBps"`
	BufferSecs    uint32 `json:"bufferSecs"`
	MinHoldSecs   uint32 `json:"minHoldSecs"`
	MaxHoldSecs   uint32 `json:"maxHoldSecs"`
	Enabled       bool   `json:"enabled"`
}

// GetAllCategories iteratered
func GetAllCategories(client *ethclient.Client, registryAddr common.Address) ([]CategoryDetail, error) {
	registry, err := contract.NewCategoryRegistry(registryAddr, client)
	if err != nil {
		return nil, fmt.Errorf("failed to bind registry: %w", err)
	}

	var categories []CategoryDetail

	for i := int64(0); ; i++ {
		id, err := registry.CategoryList(&bind.CallOpts{Context: context.Background()}, big.NewInt(i))
		if err != nil {
			break
		}

		rule, err := registry.GetRule(&bind.CallOpts{}, id)
		if err != nil {
			continue
		}

		categories = append(categories, CategoryDetail{
			ID:            common.Bytes2Hex(id[:]),
			MultiplierBps: rule.MultiplierBps,
			BufferSecs:    rule.BufferSecs,
			MinHoldSecs:   rule.MinHoldSecs,
			MaxHoldSecs:   rule.MaxHoldSecs,
			Enabled:       rule.Enabled,
		})
	}

	return categories, nil
}

// IsCategoryValid To Be or Not To Be
func IsCategoryValid(client *ethclient.Client, registryAddr common.Address, catId [32]byte) (bool, error) {
	registry, err := contract.NewCategoryRegistry(registryAddr, client)
	if err != nil {
		return false, err
	}

	rule, err := registry.GetRule(&bind.CallOpts{}, catId)
	if err != nil {
		return false, err
	}

	return rule.Enabled && rule.Existence, nil
}
