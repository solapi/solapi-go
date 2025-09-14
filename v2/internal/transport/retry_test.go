package transport

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/solapi/solapi-go/v2/internal/auth"
)

type retryOK struct {
	Message string `json:"message"`
}

func TestFetchJSON_RetryOn503ThenSuccess(t *testing.T) {
	t.Parallel()

	attempts := 0
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		attempts++
		if attempts == 1 {
			http.Error(w, "service unavailable", http.StatusServiceUnavailable)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(retryOK{Message: "ok"})
	}))
	t.Cleanup(srv.Close)

	params := auth.AuthenticationParameter{ApiKey: "key", ApiSecret: "secret"}
	req := DefaultRequest{URL: srv.URL, Method: http.MethodGet}
	res, err := FetchJSON[struct{}, retryOK](context.Background(), params, req, nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if res.Message != "ok" {
		t.Fatalf("unexpected body: %+v", res)
	}
	if attempts < 2 {
		t.Fatalf("expected retry, attempts=%d", attempts)
	}
}
