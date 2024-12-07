package utils

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"

	"regexp"
)

func GenerateApiKey() string {
	// Generate random bytes
	randomBytes := make([]byte, 32) // 32 bytes = 256 bits
	_, err := rand.Read(randomBytes)
	if err != nil {
		panic("Failed to generate random bytes")
	}

	// Hash the random bytes using SHA-256
	hash := sha256.Sum256(randomBytes)

	// Encode the hash as a hex string
	apiKey := hex.EncodeToString(hash[:])

	return apiKey
}

func IsValidEmail(email string) bool {
	// Regular expression to validate an email address
	regex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	re := regexp.MustCompile(regex)
	return re.MatchString(email)
}
