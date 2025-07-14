package util

import (
	"crypto/rand"
	"encoding/hex"
)

func GenerateAPIKey() (string, error) {
	bytes := make([]byte, 16) // 16 bytes = 32 hex characters
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}
