package crypto

import (
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/sha256"
	"math/big"
)

// SignData generates a digital signature for the given data
func SignData(privateKey *ecdsa.PrivateKey, data []byte) (r, s *big.Int, err error) {
	hash := sha256.Sum256(data)
	return ecdsa.Sign(rand.Reader, privateKey, hash[:])
}

// VerifySignature verifies a digital signature
func VerifySignature(publicKey *ecdsa.PublicKey, data []byte, r, s *big.Int) bool {
	hash := sha256.Sum256(data)
	return ecdsa.Verify(publicKey, hash[:], r, s)
}
