package utils

import (
	"crypto/sha256"
	"encoding/hex"
)

func calculateHash(data string) string {
	h := sha256.New()
	h.Write([]byte(data))
	hashed := h.Sum(nil)
	return hex.EncodeToString(hashed)
}
