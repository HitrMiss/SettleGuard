package verifier

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"crypto/x509"
	"encoding/hex"
	"encoding/pem"
	"fmt"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/crypto"
)

// LoadPrivateKey reads a PEM file from disk and returns an ECDSA Private Key
func LoadPrivateKey(path string) (*ecdsa.PrivateKey, error) {
	keyData, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read pem file: %w", err)
	}

	block, _ := pem.Decode(keyData)
	if block == nil || block.Type != "EC PRIVATE KEY" {
		return nil, fmt.Errorf("failed to decode PEM block containing EC Private Key")
	}

	privKey, err := x509.ParseECPrivateKey(block.Bytes)
	if err != nil {
		return nil, fmt.Errorf("failed to parse EC private key: %w", err)
	}

	return privKey, nil
}

// LoadMerchantKey loads the P-256 private key
func LoadMerchantKey(path string) (*ecdsa.PrivateKey, error) {
	keyBytes, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("unable to read file: %v", err)
	}

	block, _ := pem.Decode(keyBytes)
	if block == nil || block.Type != "EC PRIVATE KEY" {
		return nil, fmt.Errorf("failed to decode PEM block containing EC Private Key")
	}

	privKey, err := x509.ParseECPrivateKey(block.Bytes)
	if err != nil {
		return nil, fmt.Errorf("failed to parse EC private key: %v", err)
	}

	return privKey, nil
}

func SignManifest(privKey *ecdsa.PrivateKey, manifest string) (string, error) {
	hash := sha256.Sum256([]byte(manifest))

	r, s, err := ecdsa.Sign(rand.Reader, privKey, hash[:])
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%x%x", r, s), nil
}

// SignPacket "Authorized PacketId"
func SignPacket(priv *ecdsa.PrivateKey, data string) (string, error) {
	hash := sha256.Sum256([]byte(data))
	r, s, err := ecdsa.Sign(rand.Reader, priv, hash[:])
	if err != nil {
		return "", err
	}
	// We represent the signature as a hex string: R + S
	return fmt.Sprintf("%x%x", r, s), nil
}

// SignMessage hashes the input and signs it using the ECDSA private key
func SignMessage(privKey *ecdsa.PrivateKey, message []byte) (r, s *big.Int, err error) {
	hash := crypto.Keccak256(message)
	return ecdsa.Sign(rand.Reader, privKey, hash)
}

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
