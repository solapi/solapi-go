package auth

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"time"
)

type AuthenticationParameter struct {
	ApiKey    string
	ApiSecret string
}

var (
	ErrInvalidAPIKey    = errors.New("invalid apiKey")
	ErrInvalidAPISecret = errors.New("invalid apiSecret")
)

// BuildAuthorizationHeader builds Authorization header string with random salt and current date
func BuildAuthorizationHeader(params AuthenticationParameter) (string, error) {
	if params.ApiKey == "" {
		return "", ErrInvalidAPIKey
	}
	if params.ApiSecret == "" {
		return "", ErrInvalidAPISecret
	}
	now := time.Now().UTC()
	salt := GenerateSalt(32)
	return buildWith(params, now, salt), nil
}

// BuildAuthorizationHeaderWith builds Authorization header with provided time and optional salt
// If salt is empty, a new random salt is generated.
func BuildAuthorizationHeaderWith(params AuthenticationParameter, t time.Time, salt string) (string, error) {
	if params.ApiKey == "" {
		return "", ErrInvalidAPIKey
	}
	if params.ApiSecret == "" {
		return "", ErrInvalidAPISecret
	}
	if salt == "" {
		salt = GenerateSalt(32)
	}
	return buildWith(params, t.UTC(), salt), nil
}

func buildWith(params AuthenticationParameter, t time.Time, salt string) string {
	date := t.Format(time.RFC3339)
	mac := hmac.New(sha256.New, []byte(params.ApiSecret))
	mac.Write([]byte(date + salt))
	signature := hex.EncodeToString(mac.Sum(nil))
	return fmt.Sprintf("HMAC-SHA256 apiKey=%s, date=%s, salt=%s, signature=%s", params.ApiKey, date, salt, signature)
}
