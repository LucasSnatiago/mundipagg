package utils

import (
	"encoding/base64"
	"strings"
)

// IsStringEmpty verify if a string is empty
func IsStringEmpty(text string) bool {
	return strings.Compare(text, "") == 0
}

// ToBase64 returns a base64 of the given text
func ToBase64(text string) string {
	return base64.StdEncoding.EncodeToString([]byte(text))
}
