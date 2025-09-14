package auth

import (
	cr "crypto/rand"
)

const saltAlphabet = "1234567890abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

// GenerateSalt creates a random string with given length using [0-9a-zA-Z]
func GenerateSalt(length int) string {
	if length <= 0 {
		return ""
	}
	bytes := make([]byte, length)
	_, _ = cr.Read(bytes)
	out := make([]byte, length)
	for i := 0; i < length; i++ {
		out[i] = saltAlphabet[int(bytes[i])%len(saltAlphabet)]
	}
	return string(out)
}
