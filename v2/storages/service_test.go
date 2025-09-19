package storages

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/solapi/solapi-go/v2/internal/auth"
)

func TestService_UploadFile_BasicSuccess(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		if r.URL.Path != "/storage/v1/files" {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		var body map[string]any
		_ = json.NewDecoder(r.Body).Decode(&body)
		if body["file"] != "Zm9vYmFy" { // base64("foobar")
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(map[string]any{
			"type":         "DOCUMENT",
			"originalName": "x.txt",
			"fileId":       "fid-1",
			"name":         "x.txt",
			"url":          "http://example/test",
			"accountId":    "acc-1",
			"references":   []any{},
			"dateCreated":  "2025-09-15T00:00:00Z",
			"dateUpdated":  "2025-09-15T00:00:00Z",
		})
	}))
	defer ts.Close()

	svc := NewService(ts.URL, auth.AuthenticationParameter{ApiKey: "k", ApiSecret: "s"})
	res, err := svc.Upload(context.Background(), UploadFileRequest{File: "Zm9vYmFy", Name: "x.txt", Type: "DOCUMENT"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if res.FileID != "fid-1" {
		t.Fatalf("unexpected fileId: %s", res.FileID)
	}
	if res.Name != "x.txt" {
		t.Fatalf("unexpected name: %s", res.Name)
	}
}

func TestService_UploadFile_ServerError(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer ts.Close()

	svc := NewService(ts.URL, auth.AuthenticationParameter{ApiKey: "k", ApiSecret: "s"})
	_, err := svc.Upload(context.Background(), UploadFileRequest{File: "AAA="})
	if err == nil {
		t.Fatalf("expected error, got nil")
	}
}
