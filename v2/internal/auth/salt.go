package auth

import (
	cr "crypto/rand"
)

const saltAlphabet = "1234567890abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

// randRead is a seam for crypto/rand.Read to enable testing error paths.
var randRead = cr.Read

// GenerateSalt creates a random string with given length using [0-9a-zA-Z]
func GenerateSalt(length int) string {
	if length <= 0 {
		return ""
	}
	bytes := make([]byte, length)
	if _, err := randRead(bytes); err != nil {
		panic("fatal: failed to read from crypto/rand: " + err.Error())
	}
	out := make([]byte, length)
	for i := 0; i < length; i++ {
		out[i] = saltAlphabet[int(bytes[i])%len(saltAlphabet)]
	}
	return string(out)
}
