package yellow

import (
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

// SignYellowMessage standard expected by clearing nodes.
func SignYellowMessage(payload interface{}, privateKeyHex string) (string, error) {
	data, err := json.Marshal(payload)
	if err != nil {
		return "", err
	}

	pKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		return "", fmt.Errorf("invalid private key: %w", err)
	}

	hash := crypto.Keccak256Hash(
		[]byte(fmt.Sprintf("\x19Ethereum Signed Message:\n%d%s", len(data), data)),
	)

	signature, err := crypto.Sign(hash.Bytes(), pKey)
	if err != nil {
		return "", err
	}

	signature[64] += 27

	return hexutil.Encode(signature), nil
}
