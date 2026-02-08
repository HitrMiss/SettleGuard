package api

import (
	"encoding/json"
	"net/http"

	"SettleGuard/internal/repository"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func HandleListCategories(w http.ResponseWriter, r *http.Request, client *ethclient.Client, registryAddr common.Address) {
	cats, err := repository.GetAllCategories(client, registryAddr)
	if err != nil {
		http.Error(w, "Failed to fetch categories", 500)
		return
	}
	json.NewEncoder(w).Encode(cats)
}
