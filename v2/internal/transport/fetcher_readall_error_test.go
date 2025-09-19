package transport

import (
	"context"
	"errors"
	"net/http"
	"strings"
	"testing"

	"github.com/solapi/solapi-go/v2/internal/auth"
)

// erroringBody is an io.ReadCloser that always fails on Read
type erroringBody struct{}

func (erroringBody) Read(p []byte) (int, error) { return 0, errors.New("forced read error") }
func (erroringBody) Close() error               { return nil }

// When server returns non-2xx and body read fails, ensure error surfaces meaningfully
func TestFetchJSON_Non2xxAndBodyReadError(t *testing.T) {
	t.Parallel()

	// RoundTripper returns a response with status 500 and an erroring body
	rt := roundTripperFunc(func(req *http.Request) (*http.Response, error) {
		return &http.Response{
			StatusCode: http.StatusInternalServerError,
			Body:       erroringBody{},
			Header:     make(http.Header),
			Request:    req,
		}, nil
	})

	httpClient := &http.Client{Transport: rt}

	params := auth.AuthenticationParameter{ApiKey: "key", ApiSecret: "secret"}
	req := DefaultRequest{URL: "http://example.invalid", Method: http.MethodGet}

	_, err := FetchJSONWithClient[struct{}, okResponse](context.Background(), httpClient, params, req, nil)
	if err == nil {
		t.Fatalf("expected error, got nil")
	}
	// We expect the error message to include the underlying read error
	if !strings.Contains(err.Error(), "forced read error") {
		t.Fatalf("expected error to mention read failure, got: %v", err)
	}
}
