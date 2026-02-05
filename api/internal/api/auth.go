package api

import (
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"SettleGuard/internal/repository"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

func HandleGetNonce(w http.ResponseWriter, r *http.Request) {
	address := r.URL.Query().Get("address")
	if address == "" {
		http.Error(w, "address required", 400)
		return
	}

	b := make([]byte, 16)
	rand.Read(b)
	nonce := hex.EncodeToString(b)

	repository.SaveSession(repository.UserSession{
		Address:   strings.ToLower(address),
		Nonce:     nonce,
		ExpiresAt: time.Now().Add(10 * time.Minute).Unix(),
	})

	json.NewEncoder(w).Encode(map[string]string{"nonce": nonce})
}

func HandleVerifySignature(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Address   string `json:"address"`
		Signature string `json:"signature"`
	}
	json.NewDecoder(r.Body).Decode(&req)

	session, err := repository.GetSession(strings.ToLower(req.Address))
	if err != nil || time.Now().Unix() > session.ExpiresAt {
		http.Error(w, "Invalid or expired session", 401)
		return
	}

	// Check your self before you wreck yourself
	msg := fmt.Sprintf("\x19Ethereum Signed Message:\n%d%s", len(session.Nonce), session.Nonce)
	msgHash := crypto.Keccak256Hash([]byte(msg))

	sig, _ := hexutil.Decode(req.Signature)
	if sig[64] == 27 || sig[64] == 28 {
		sig[64] -= 27
	}

	pubKey, err := crypto.SigToPub(msgHash.Bytes(), sig)
	recoveredAddr := crypto.PubkeyToAddress(*pubKey)

	if strings.ToLower(recoveredAddr.Hex()) != strings.ToLower(req.Address) {
		http.Error(w, "Signature verification failed", 403)
		return
	}

	session.IsAuthed = true
	repository.SaveSession(*session)

	json.NewEncoder(w).Encode(map[string]string{"status": "authenticated"})
}
