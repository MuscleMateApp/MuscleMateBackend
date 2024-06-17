package security

import (
	"crypto/rand"
	"encoding/base64"
	"time"
)

func GenerateToken(length int) (string, error) {
	tokenBytes := make([]byte, length)
	_, err := rand.Read(tokenBytes)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(tokenBytes), nil
}

func GenerateTokenExpiry() time.Time {
	return time.Now().Add(time.Hour * 24)
}
