package provider

import (
	"bytes"
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

// TODO: Wicked migraine, really check this code
type LocalWallet struct {
	PrivateKey *ecdsa.PrivateKey
	Address    common.Address
}

// CreateWallet generates or loads a secp256k1 wallet for Ethereum gas
func CreateWallet() (*LocalWallet, error) {
	priv, err := crypto.GenerateKey()
	if err != nil {
		return nil, err
	}

	addr := crypto.PubkeyToAddress(priv.PublicKey)
	return &LocalWallet{PrivateKey: priv, Address: addr}, nil
}

// ExportHex in case you are under a curse, the curse of life
func (w *LocalWallet) ExportHex() string {
	return hex.EncodeToString(crypto.FromECDSA(w.PrivateKey))
}

// RequestAlchemyFaucet uses the API key already present in the Provider's URL
func (p *Provider) RequestAlchemyFaucet(address string) error {
	parts := strings.Split(p.URL, "/")
	if len(parts) < 1 {
		return fmt.Errorf("invalid provider URL format")
	}
	apiKey := parts[len(parts)-1]

	network := "sepolia"
	if strings.Contains(p.URL, "mainnet") {
		return fmt.Errorf("faucet not available for mainnet")
	}

	apiURL := "https://dashboard.alchemy.com/api/faucet/request-funds"

	payload := map[string]string{
		"address": address,
		"network": network,
		"apiKey":  apiKey,
	}

	jsonData, _ := json.Marshal(payload)
	req, err := http.NewRequest("POST", apiURL, bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("alchemy faucet error (status %d): %s", resp.StatusCode, string(body))
	}

	fmt.Printf("🌊 Processed for 💰 %s\n", address)
	return nil
}

// WaitForBalance I wonder what song would take this long?
func (p *Provider) WaitForBalance(address common.Address, timeout time.Duration) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()

	fmt.Printf("⏳ Waiting for 💰 to arrive at %s...\n", address.Hex())

	for {
		select {
		case <-ctx.Done():
			return fmt.Errorf("timed out waiting for 💰")
		case <-ticker.C:
			balance, err := p.Client.BalanceAt(context.Background(), address, nil)
			if err != nil {
				continue
			}
			if balance.Sign() > 0 {
				fmt.Printf("Take 💰 and run! Balance: %s wei\n", balance.String())
				return nil
			}
		}
	}
}
