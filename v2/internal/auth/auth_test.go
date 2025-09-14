package auth

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"regexp"
	"strings"
	"testing"
	"time"
)

func TestBuildAuthorizationHeaderWith_Deterministic(t *testing.T) {
	params := AuthenticationParameter{ApiKey: "testKey", ApiSecret: "testSecret"}
	fixed := time.Date(2025, 9, 14, 12, 34, 56, 0, time.UTC)
	salt := "Abc123XYZ789abcdEFGHijklMNOPqr"
	got, err := BuildAuthorizationHeaderWith(params, fixed, salt)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	date := fixed.Format(time.RFC3339)
	mac := hmac.New(sha256.New, []byte(params.ApiSecret))
	mac.Write([]byte(date + salt))
	sig := hex.EncodeToString(mac.Sum(nil))
	want := "HMAC-SHA256 apiKey=" + params.ApiKey + ", date=" + date + ", salt=" + salt + ", signature=" + sig
	if got != want {
		t.Fatalf("mismatch.\nwant: %s\n got: %s", want, got)
	}
}

func TestBuildAuthorizationHeader_InvalidKeys(t *testing.T) {
	_, err := BuildAuthorizationHeader(AuthenticationParameter{ApiKey: "", ApiSecret: "x"})
	if err == nil {
		t.Fatalf("expected error for empty apiKey")
	}
	_, err = BuildAuthorizationHeader(AuthenticationParameter{ApiKey: "x", ApiSecret: ""})
	if err == nil {
		t.Fatalf("expected error for empty apiSecret")
	}
}

func TestBuildAuthorizationHeader_SaltFormat(t *testing.T) {
	params := AuthenticationParameter{ApiKey: "k", ApiSecret: "s"}
	got, err := BuildAuthorizationHeaderWith(params, time.Now().UTC(), "")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	// extract salt
	parts := strings.Split(got, ", ")
	var salt string
	for _, p := range parts {
		if strings.HasPrefix(p, "salt=") {
			salt = strings.TrimPrefix(p, "salt=")
		}
	}
	if salt == "" {
		t.Fatalf("salt not found in header: %s", got)
	}
	if len(salt) != 32 {
		t.Fatalf("salt length = %d, want 32", len(salt))
	}
	if ok, _ := regexp.MatchString("^[0-9A-Za-z]{32}$", salt); !ok {
		t.Fatalf("salt contains invalid charset: %s", salt)
	}
}
