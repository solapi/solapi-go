package client

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/solapi/solapi-go/v2/storages"
)

func TestClient_Storages_Upload_WiresThrough(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		if r.URL.Path != "/storage/v1/files" {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(map[string]any{
			"type":         "DOCUMENT",
			"originalName": "a.txt",
			"fileId":       "fid-2",
			"name":         "a.txt",
			"url":          "http://x",
			"accountId":    "acc",
			"references":   []any{},
			"dateCreated":  "",
			"dateUpdated":  "",
		})
	}))
	defer ts.Close()

	c := newClientWithBaseURL(ts.URL, "k", "s")
	res, err := c.Storages.Upload(context.Background(), storages.UploadFileRequest{File: "QQ==", Name: "a.txt", Type: "DOCUMENT"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if res.FileID != "fid-2" {
		t.Fatalf("unexpected fileId: %s", res.FileID)
	}
}
