package crypto

import (
	"crypto/sha256"
	"encoding/hex"
)

// ComputeSHA256 computes the SHA-256 hash of input data
func ComputeSHA256(data []byte) string {
	hash := sha256.Sum256(data)
	return hex.EncodeToString(hash[:])
}
