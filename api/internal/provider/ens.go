package provider

import (
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/wealdtech/go-ens/v3"
)

// SetMerchantPubKey updates the "settleGuard.p256" text record for an ENS name
// this is for testing only and not a production level function!!! requires wallet key DANGER WILL ROBINSON
func (p *Provider) SetMerchantPubKey(auth *bind.TransactOpts, domain string, compressedKey string) (common.Hash, error) {
	resolver, err := ens.NewResolver(p.Client, domain)
	if err != nil {
		return common.Hash{}, err
	}

	// We use the "SetText" method on the resolver
	tx, err := resolver.SetText(auth, "settleGuard.p256", compressedKey)
	if err != nil {
		return common.Hash{}, fmt.Errorf("failed to update ENS record: %v", err)
	}

	return tx.Hash(), nil
}

// GetMerchantPubKey fetch merchant public key text record from ENS.
func (p *Provider) GetMerchantPubKey(domain string) (string, error) {
	resolver, err := ens.NewResolver(p.Client, domain)
	if err != nil {
		return "", fmt.Errorf("failed to resolve ENS name %s: %w", domain, err)
	}

	key, err := resolver.Text("settleGuard.p256")
	if err != nil {
		return "", fmt.Errorf("failed to fetch text record for %s: %w", domain, err)
	}

	if key == "" {
		return "", fmt.Errorf("no 'settleGuard.p256' text record found for %s", domain)
	}

	return key, nil
}
