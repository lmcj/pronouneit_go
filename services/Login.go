package services

import (
	"crypto/sha256"
	"encoding/hex"
)

func Hash256(password string) string {
	passwordBytes := []byte(password)

	hasher := sha256.New()
	hasher.Write(passwordBytes)

	hashSum := hasher.Sum(nil)

	hashString := hex.EncodeToString(hashSum)

	return hashString
}
