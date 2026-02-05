package verifier

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"encoding/hex"
	"fmt"
	"math/big"
)

// VerifyP256Signature checks a signature against a compressed P-256 public key.
func VerifyP256Signature(compressedPubKeyHex string, messageHash []byte, r, s *big.Int) (bool, error) {
	pubKeyBytes, err := hex.DecodeString(compressedPubKeyHex)
	if err != nil || len(pubKeyBytes) != 33 {
		return false, fmt.Errorf("invalid compressed pubkey: %v", err)
	}

	x, y := elliptic.UnmarshalCompressed(elliptic.P256(), pubKeyBytes)
	if x == nil {
		return false, fmt.Errorf("failed to decompress P-256 public key")
	}

	pubKey := &ecdsa.PublicKey{
		Curve: elliptic.P256(),
		X:     x,
		Y:     y,
	}

	return ecdsa.Verify(pubKey, messageHash, r, s), nil
}

// VerifyMerchantSignature takes the compressed key from ENS and validates the signature
func VerifyMerchantSignature(compressedKeyHex string, messageHash []byte, rHex, sHex string) (bool, error) {
	pubKeyBytes, err := hex.DecodeString(compressedKeyHex)
	if err != nil {
		return false, fmt.Errorf("invalid pubkey hex: %v", err)
	}

	x, y := elliptic.UnmarshalCompressed(elliptic.P256(), pubKeyBytes)
	if x == nil {
		return false, fmt.Errorf("failed to decompress P-256 key")
	}

	pubKey := &ecdsa.PublicKey{
		Curve: elliptic.P256(),
		X:     x,
		Y:     y,
	}

	r := new(big.Int)
	s := new(big.Int)
	r.SetString(rHex, 16)
	s.SetString(sHex, 16)

	return ecdsa.Verify(pubKey, messageHash, r, s), nil
}
