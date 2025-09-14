package transport

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/solapi/solapi-go/v2/internal/auth"
)

func TestFetchJSON_4xxReturnsApiError(t *testing.T) {
	t.Parallel()

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(map[string]string{
			"errorCode":    "InvalidParam",
			"errorMessage": "bad",
		})
	}))
	t.Cleanup(srv.Close)

	params := auth.AuthenticationParameter{ApiKey: "key", ApiSecret: "secret"}
	req := DefaultRequest{URL: srv.URL, Method: http.MethodGet}
	var out map[string]any
	err := FetchJSON(context.Background(), params, req, nil, &out)
	var apiErr *ApiError
	if !errors.As(err, &apiErr) {
		t.Fatalf("expected ApiError, got %v", err)
	}
	if apiErr.HTTPStatus != 400 || apiErr.ErrorCode != "InvalidParam" {
		t.Fatalf("unexpected apiErr: %+v", apiErr)
	}
}

func TestFetchJSON_5xxReturnsDefaultError(t *testing.T) {
	t.Parallel()

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "boom", http.StatusInternalServerError)
	}))
	t.Cleanup(srv.Close)

	params := auth.AuthenticationParameter{ApiKey: "key", ApiSecret: "secret"}
	req := DefaultRequest{URL: srv.URL, Method: http.MethodGet}
	var out map[string]any
	err := FetchJSON(context.Background(), params, req, nil, &out)
	var defErr *DefaultError
	if !errors.As(err, &defErr) {
		t.Fatalf("expected DefaultError, got %v", err)
	}
}
