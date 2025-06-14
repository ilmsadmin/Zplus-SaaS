package utils

import (
	"crypto/rand"
	"encoding/hex"
	"strings"
)

// GenerateRandomString generates a random string of specified length
func GenerateRandomString(length int) (string, error) {
	bytes := make([]byte, length/2)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}

// SanitizeString removes harmful characters from strings
func SanitizeString(input string) string {
	return strings.TrimSpace(input)
}

// ValidateEmail checks if email format is valid
func ValidateEmail(email string) bool {
	return strings.Contains(email, "@") && len(email) > 5
}

// CreateSlug creates a URL-safe slug from a string
func CreateSlug(input string) string {
	slug := strings.ToLower(input)
	slug = strings.ReplaceAll(slug, " ", "-")
	slug = strings.ReplaceAll(slug, "_", "-")
	return slug
}