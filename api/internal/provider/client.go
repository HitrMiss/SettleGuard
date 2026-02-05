package provider

import (
	"fmt"
	"os"

	"github.com/ethereum/go-ethereum/ethclient"
)

type Provider struct {
	Client *ethclient.Client
	URL    string
}

// NewProvider connection to local first alchemy worse.
func NewProvider(network string) (*Provider, error) {
	var url string

	if local := os.Getenv("LOCAL_NODE_URL"); local != "" {
		url = local
	} else {
		key := os.Getenv("ALCHEMY_API_KEY")
		url = fmt.Sprintf("https://%s.g.alchemy.com/v2/%s", network, key)
	}

	client, err := ethclient.Dial(url)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to %s: %v", url, err)
	}

	return &Provider{
		Client: client,
		URL:    url,
	}, nil
}
