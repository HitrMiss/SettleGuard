package api

import (
	"SettleGuard/internal/config"
	"SettleGuard/internal/provider"
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/sha256"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	"net/http"
	"time"

	"SettleGuard/internal/contract"
	"SettleGuard/internal/repository"
	"SettleGuard/internal/verifier"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"

	// "github.com/ethereum/go-ethereum/crypto" switched signing code
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

// I need to stop copying and pasting this
const (
	VaultKey      = "SettleGuard#PaymentVault"
	RegistryKey   = "SettleGuard#CategoryRegistry"
	SettlementKey = "SettleGuard#SettlementEngine"
	USDCKey       = "SettleGuard#MockUSDC"
)

var merchantKey *ecdsa.PrivateKey

func InitMerchant(pemPath string) {
	key, err := verifier.LoadPrivateKey(pemPath)
	if err != nil {
		log.Fatal("Critical: Could not load merchant PEM", err)
	}
	merchantKey = key
}

func HandleApprovalCheck(w http.ResponseWriter, r *http.Request, client *ethclient.Client) {
	userAddr := r.URL.Query().Get("address")

	cfg, err := config.LoadConfig("cmd/api/backend_config.json")
	if err != nil {
		panic(err)
	}

	vaultAddrHex := common.HexToAddress(cfg.Contracts[VaultKey])
	usdcHex := common.HexToAddress(cfg.Contracts[USDCKey])
	p, err := provider.NewProvider()

	if err != nil {
		log.Fatalf("[-] Connection failed: %v", err)
	}
	//vault, err := contract.NewPaymentVault(vaultAddrHex, p.Client)
	//if err != nil {
	//	fmt.Printf("❌ Failed to bind Vault: %v\n", err)
	//	break
	//}

	usdc, err := contract.NewMockUSDC(usdcHex, p.Client)
	if err != nil {
		fmt.Printf("❌ Failed to bind mUSDC: %v\n", err)
	}
	allowance, err := usdc.Allowance(nil, common.HexToAddress(userAddr), vaultAddrHex)

	isAuthorized := err == nil && allowance.Cmp(big.NewInt(0)) > 0

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]bool{"authorized": isAuthorized})
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
	//catId := common.HexToHash(req.CategoryId)
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

	timestamp := uint64(time.Now().Unix())

	// Seed: Identity + Payout + Amount + Category + Timestamp
	manifest := fmt.Sprintf("%s:%s:%s:%s:%d",
		req.MerchantIdentity,
		req.MerchantPayout,
		amountUSDC.String(),
		req.CategoryId,
		timestamp,
	)
	merchantPrivKey, err := verifier.LoadMerchantKey("merchant_test.pem")
	if err != nil {
		http.Error(w, "Server configuration error: Key not found", 500)
		return
	}
	//signedPacketId, err := verifier.SignPacket(merchantPrivKey, manifest)
	//if err != nil {
	//	http.Error(w, "Signing failed", http.StatusInternalServerError)
	//	return
	//}

	// We get R and S values
	hash := sha256.Sum256([]byte(manifest))
	rVal, sVal, err := ecdsa.Sign(rand.Reader, merchantPrivKey, hash[:])
	if err != nil {
		http.Error(w, "Signing failed", 500)
		return
	}
	//packetBytes, err := hex.DecodeString(strings.TrimPrefix(signedPacketId, "0x"))
	//if err != nil {
	//	http.Error(w, "Invalid packet ID format", http.StatusBadRequest)
	//	return
	//}
	//var packetIdFixed [32]byte
	//
	//copy(packetIdFixed[:], packetBytes)

	rBytes := common.LeftPadBytes(rVal.Bytes(), 32)
	sBytes := common.LeftPadBytes(sVal.Bytes(), 32)
	tBytes := make([]byte, 8)
	binary.BigEndian.PutUint64(tBytes, timestamp)

	// Concatenate r + s + timestamp
	packed := append(rBytes, sBytes...)
	packed = append(packed, tBytes...)

	// Hash it (Keccak256, NOT SHA256)
	packetIdHash := crypto.Keccak256Hash(packed)

	encodedData, err := contract.EncodeDepositCall(
		amountUSDC,
		common.HexToAddress(req.MerchantIdentity),
		common.HexToAddress(req.MerchantPayout),
		catId,
		packetIdHash,
		timestamp,
		rVal,
		sVal,
	)
	if err != nil {
		http.Error(w, "Failed to encode contract call: "+err.Error(), http.StatusInternalServerError)
		return
	}

	//err = repository.SaveOrderRecord(
	//	packetId.Hex(),
	//	payerAddrStr,
	//	req.Items,
	//	amountUSDC.String(),
	//	createdAt,
	//	req.CategoryId,
	//)
	err = repository.SaveOrderRecord(
		packetIdHash.String(),
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

	resp := CheckoutResponse{
		PacketId:       packetIdHash.Hex(),
		Amount:         "0x" + amountUSDC.Text(16),
		USDCAddress:    usdcAddr.Hex(),
		TargetContract: engineAddr.Hex(),
		EncodedABI:     encodedData,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
