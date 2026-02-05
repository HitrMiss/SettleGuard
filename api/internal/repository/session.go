package repository

import (
	"encoding/json"
	"os"
	"path/filepath"
)

type UserSession struct {
	Address   string `json:"address"`
	Nonce     string `json:"nonce"`
	ExpiresAt int64  `json:"expiresAt"`
	IsAuthed  bool   `json:"isAuthed"`
}

func SaveSession(session UserSession) error {
	dir := "./data/sessions"
	os.MkdirAll(dir, 0755)

	filePath := filepath.Join(dir, session.Address+".json")
	data, _ := json.MarshalIndent(session, "", "  ")
	return os.WriteFile(filePath, data, 0644)
}

func GetSession(address string) (*UserSession, error) {
	filePath := filepath.Join("./data/sessions", address+".json")
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var session UserSession
	json.Unmarshal(data, &session)
	return &session, nil
}
