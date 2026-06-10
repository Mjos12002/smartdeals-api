package utils

import (
	"crypto/sha256"
	"encoding/hex"
)

func EncryptAES(plainText string) (string, error) {
	data := []byte(plainText)
	hash := sha256.Sum256(data)
	return hex.EncodeToString(hash[:]), nil
}
