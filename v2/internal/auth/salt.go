package auth

import (
	cr "crypto/rand"
	"math/big"
)

const saltAlphabet = "1234567890abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

// randRead is a seam for crypto/rand.Read to enable testing error paths.
var randRead = cr.Read

// randReader bridges randRead into an io.Reader so we can use crypto/rand.Int
// while still allowing tests to inject failures via randRead.
type randReader struct{}

func (randReader) Read(p []byte) (int, error) { return randRead(p) }

// GenerateSalt creates a random string with given length using [0-9a-zA-Z]
func GenerateSalt(length int) string {
	if length <= 0 {
		return ""
	}
	out := make([]byte, length)
	for i := range out {
		num, err := cr.Int(randReader{}, big.NewInt(int64(len(saltAlphabet))))
		if err != nil {
			// This is extremely unlikely to happen and indicates a problem with the OS's entropy source.
			panic("fatal: failed to read from crypto/rand: " + err.Error())
		}
		out[i] = saltAlphabet[num.Int64()]
	}
	return string(out)
}
