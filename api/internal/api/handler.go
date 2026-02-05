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
)

// TODO: head really hurts, I am starting to lose track of the code base in my head... This should be ENV or a config
const (
	SepoliaUSDCAddress      = "0x3793D61a6db2Da1fCEe93616A558332A2fC05A4A"
	SettlementEngineAddress = "0x"
)

type CartItem struct {
	ID    string  `json:"id"`
	Price float64 `json:"price"`
	Name  string  `json:"name"`
}

type CheckoutRequest struct {
	Items            []CartItem `json:"items"`
	MerchantIdentity string     `json:"merchantIdentity"`
	MerchantPayout   string     `json:"merchantPayout"`
	CategoryId       string     `json:"categoryId"`
}

type CheckoutResponse struct {
	PacketId       string `json:"packetId"`
	Amount         string `json:"amount"`
	USDCAddress    string `json:"usdcAddress"`
	TargetContract string `json:"targetContract"`
	EncodedABI     string `json:"encodedABI"`
}

func HandleCheckout(w http.ResponseWriter, r *http.Request) {
	var req CheckoutRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
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

	// comes from the chain NOT TO BE PASSED IN!
	catId := crypto.Keccak256Hash([]byte(req.CategoryId))

	packetSeed := fmt.Sprintf("%s-%s-%s-%s-%d",
		req.MerchantIdentity,
		req.MerchantPayout,
		amountUSDC.String(),
		req.CategoryId,
		createdAt,
	)
	packetId := crypto.Keccak256Hash([]byte(packetSeed))

	identityAddr := common.HexToAddress(req.MerchantIdentity)
	payoutAddr := common.HexToAddress(req.MerchantPayout)

	encodedData, err := contract.EncodePaymentCall(
		amountUSDC,
		identityAddr,
		payoutAddr,
		catId,
		packetId,
	)
	if err != nil {
		http.Error(w, "Encoding failed: "+err.Error(), http.StatusInternalServerError)
		return
	}

	err = repository.SaveOrderRecord(packetId.Hex(), req.Items, amountUSDC.String(), createdAt)
	if err != nil {
		http.Error(w, "Persistence failed", http.StatusInternalServerError)
		return
	}

	resp := CheckoutResponse{
		PacketId:       packetId.Hex(),
		Amount:         "0x" + amountUSDC.Text(16),
		USDCAddress:    SepoliaUSDCAddress,
		TargetContract: SettlementEngineAddress,
		EncodedABI:     encodedData,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
