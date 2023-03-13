package utils

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
)

func GenerateToken(length int) string {
	randomBytes := make([]byte, length)
	rand.Read(randomBytes)
	return hex.EncodeToString(randomBytes)
}

func HashTokenString(str string) string {
	h := sha256.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}
