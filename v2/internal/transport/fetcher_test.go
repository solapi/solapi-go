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
