package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	Network     string            `json:"network"`
	Contracts   map[string]string `json:"contracts"`
	CategoryIDs map[string]string `json:"categoryIds"`
}

func LoadConfig(path string) (*Config, error) {
	file, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read config: %w", err)
	}

	var cfg Config
	if err := json.Unmarshal(file, &cfg); err != nil {
		return nil, fmt.Errorf("failed to parse config: %w", err)
	}

	return &cfg, nil
}
