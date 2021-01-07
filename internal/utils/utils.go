package utils

import "strings"

// IsStringEmpty verify if a string is empty
func IsStringEmpty(text string) bool {
	return strings.Compare(text, "") == 0
}
