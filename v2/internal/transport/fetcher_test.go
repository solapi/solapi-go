package transport

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/solapi/solapi-go/v2/internal/auth"
)

type okResponse struct {
	Message string `json:"message"`
}

func TestFetchJSON_SuccessAndAuthorizationHeader(t *testing.T) {
	t.Parallel()

	var sawAuthorization bool

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Authorization") != "" {
			sawAuthorization = true
		}
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(okResponse{Message: "ok"})
	}))
	t.Cleanup(srv.Close)

	params := auth.AuthenticationParameter{ApiKey: "key", ApiSecret: "secret"}
	req := DefaultRequest{URL: srv.URL, Method: http.MethodGet}

	res, err := FetchJSON[struct{}, okResponse](context.Background(), params, req, nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if res.Message != "ok" {
		t.Fatalf("unexpected body: %+v", res)
	}
	if !sawAuthorization {
		t.Fatalf("Authorization header was not set")
	}
}

// Verify we can use a provided http.Client (with custom Transport) instead of http.DefaultClient
func TestFetchJSON_WithCustomClient(t *testing.T) {
	t.Parallel()

	// Test server to echo a header that our custom RoundTripper will set
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(okResponse{Message: r.Header.Get("X-Custom-From-Client")})
	}))
	t.Cleanup(srv.Close)

	// custom transport injects header
	rt := roundTripperFunc(func(req *http.Request) (*http.Response, error) {
		req.Header.Set("X-Custom-From-Client", "yes")
		return http.DefaultTransport.RoundTrip(req)
	})

	httpClient := &http.Client{Transport: rt}

	params := auth.AuthenticationParameter{ApiKey: "key", ApiSecret: "secret"}
	req := DefaultRequest{URL: srv.URL, Method: http.MethodGet}

	res, err := FetchJSONWithClient[struct{}, okResponse](context.Background(), httpClient, params, req, nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if res.Message != "yes" {
		t.Fatalf("custom client not used, message=%q", res.Message)
	}
}

// minimal roundTripperFunc helper
type roundTripperFunc func(*http.Request) (*http.Response, error)

func (f roundTripperFunc) RoundTrip(r *http.Request) (*http.Response, error) { return f(r) }
