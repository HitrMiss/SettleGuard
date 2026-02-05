package provider

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
)

type Provider struct {
	Client *ethclient.Client
	URL    string
}

func (p *Provider) MaskedURL() string {
	if !strings.Contains(p.URL, ".alchemy.com/v2/") {
		return p.URL
	}

	parts := strings.Split(p.URL, "/v2/")
	if len(parts) < 2 {
		return "invalid-url"
	}
	// Mask up, Money Heist Style
	return fmt.Sprintf("%s/v2/********", parts[0])
}

func NewProvider() (*Provider, error) {
	err := godotenv.Load()
	if err != nil {
		log.Println("ℹ️ No .env file found; using system environment variables")
	}

	var url string

	network := strings.ToLower(strings.TrimSpace(os.Getenv("BLOCKCHAIN")))

	if network == "" {
		return nil, fmt.Errorf("🚨 BLOCKCHAIN environment variable is not set")
	}

	if network == "localhost" || network == "local" {
		url = os.Getenv("LOCAL_NODE_URL")
		if url == "" {
			url = "http://127.0.0.1:8545"
		}
	} else {
		key := os.Getenv("ALCHEMY_API_KEY")
		if key == "" {
			return nil, fmt.Errorf("🚨 ALCHEMY_API_KEY missing for network: %s", network)
		}

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
