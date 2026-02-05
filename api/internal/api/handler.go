package api

import (
	"encoding/json"
	"fmt"
	"math/big"
	"net/http"
	"time"

	"SettleGuard/internal/contract"
	"SettleGuard/internal/repository"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

// CheckoutRequest defines the payload from the frontend cart
type CheckoutRequest struct {
	Items            []repository.CartItem `json:"items"`
	MerchantIdentity string                `json:"merchantIdentity"`
	MerchantPayout   string                `json:"merchantPayout"`
	CategoryId       string                `json:"categoryId"` // Expected as Hex "0x..."
}

// CheckoutResponse returns the data needed for MetaMask to execute the transaction
type CheckoutResponse struct {
	PacketId       string `json:"packetId"`
	Amount         string `json:"amount"` // Hex string for contract
	USDCAddress    string `json:"usdcAddress"`
	TargetContract string `json:"targetContract"`
	EncodedABI     string `json:"encodedAbi"`
}

// HandleCheckout processes the cart, validates on-chain rules, and returns encoded ABI
func HandleCheckout(
	w http.ResponseWriter,
	r *http.Request,
	client *ethclient.Client,
	registryAddr common.Address,
	engineAddr common.Address,
	usdcAddr common.Address,
) {
	var req CheckoutRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// We assume the user address is passed in the "X-User-Address" header
	payerAddrStr := r.Header.Get("X-User-Address")
	if payerAddrStr == "" {
		http.Error(w, "User address header missing", http.StatusUnauthorized)
		return
	}

	catId := common.HexToHash(req.CategoryId)
	isValid, err := repository.IsCategoryValid(client, registryAddr, catId)
	if err != nil || !isValid {
		http.Error(w, "Category is invalid or disabled in Registry", http.StatusForbidden)
		return
	}

	createdAt := uint64(time.Now().Unix())
	total := 0.0
	for _, item := range req.Items {
		total += item.Price
	}

	amountUSDC := new(big.Int)
	bigTotal := new(big.Float).SetFloat64(total)
	multiplier := new(big.Float).SetFloat64(1e6)
	bigTotal.Mul(bigTotal, multiplier).Int(amountUSDC)

	// Seed: Identity + Payout + Amount + Category + Timestamp
	packetSeed := fmt.Sprintf("%s-%s-%s-%s-%d",
		req.MerchantIdentity,
		req.MerchantPayout,
		amountUSDC.String(),
		catId.Hex(),
		createdAt,
	)
	packetId := crypto.Keccak256Hash([]byte(packetSeed))

	encodedData, err := contract.EncodePaymentCall(
		amountUSDC,
		common.HexToAddress(req.MerchantIdentity),
		common.HexToAddress(req.MerchantPayout),
		catId,
		packetId,
	)
	if err != nil {
		http.Error(w, "Failed to encode contract call: "+err.Error(), http.StatusInternalServerError)
		return
	}

	err = repository.SaveOrderRecord(
		packetId.Hex(),
		payerAddrStr,
		req.Items,
		amountUSDC.String(),
		createdAt,
		req.CategoryId,
	)
	if err != nil {
		http.Error(w, "Failed to persist order record", http.StatusInternalServerError)
		return
	}

	// 7. Success Response
	resp := CheckoutResponse{
		PacketId:       packetId.Hex(),
		Amount:         "0x" + amountUSDC.Text(16),
		USDCAddress:    usdcAddr.Hex(),
		TargetContract: engineAddr.Hex(),
		EncodedABI:     encodedData,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
