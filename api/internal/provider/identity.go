package provider

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"
)

// GenerateP256Key creates a new secp256r1 private key
func GenerateP256Key() (*ecdsa.PrivateKey, error) {
	return ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
}

// SavePrivateKey saves the P-256 private key to a local PEM file
func SavePrivateKey(filename string, key *ecdsa.PrivateKey) error {
	marshaledKey, err := x509.MarshalECPrivateKey(key)
	if err != nil {
		return err
	}

	pemBlock := &pem.Block{
		Type:  "EC PRIVATE KEY",
		Bytes: marshaledKey,
	}

	f, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		return err
	}
	defer f.Close()

	return pem.Encode(f, pemBlock)
}

// SaveKeyPair saves the Private Key (PEM) and the Compressed Public Key (hex)
// don't ask why the key is stored compressed it just is okay!
// truth is I know why it's for the CLI tool bwhahahaha
func SaveKeyPair(baseName string, key *ecdsa.PrivateKey) error {
	marshaledPriv, err := x509.MarshalECPrivateKey(key)
	if err != nil {
		return err
	}
	privBlock := &pem.Block{Type: "EC PRIVATE KEY", Bytes: marshaledPriv}
	err = os.WriteFile(baseName+".pem", pem.EncodeToMemory(privBlock), 0600)
	if err != nil {
		return fmt.Errorf("failed to save private key: %w", err)
	}

	compressed := EncodeCompressedPubKey(&key.PublicKey)
	err = os.WriteFile(baseName+".pub", []byte(compressed), 0644)
	if err != nil {
		return fmt.Errorf("failed to save public key: %w", err)
	}

	return nil
}

// EncodeCompressedPubKey returns the compressed hex string of a P-256 public key
func EncodeCompressedPubKey(pub *ecdsa.PublicKey) string {
	format := byte(0x02)
	if pub.Y.Bit(0) != 0 {
		format = 0x03
	}
	// Use %064x to pad to 32 bytes (64 hex chars)
	return fmt.Sprintf("%02x%064x", format, pub.X)
}

//Comments be like  https://www.reddit.com/r/ProgrammerHumor/comments/8w54mx/code_comments_be_like/
