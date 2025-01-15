package crypto

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"encoding/pem"
	"os"
)

// GenerateKeyPair generates a new ECDSA key pair
func GenerateKeyPair() (*ecdsa.PrivateKey, error) {
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		return nil, err
	}

	return privateKey, nil
}

// SavePrivateKey saves the private key to a PEM file
func SavePrivateKey(fileName string, privateKey *ecdsa.PrivateKey) error {
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()
	keyBytes, err := x509.MarshalECPrivateKey(privateKey)
	if err != nil {
		return err
	}
	pem.Encode(file, &pem.Block{Type: "EC PRIVATE KEY", Bytes: keyBytes})
	return nil
}
