package groups

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/solapi/solapi-go/v2/internal/auth"
)

func TestGroups_ListMessages_BasicSuccess(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		if r.URL.Path != "/messages/v4/groups/g1/messages" {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(map[string]any{
			"messageList": map[string]any{},
			"limit":       10,
			"startKey":    "",
			"nextKey":     "",
		})
	}))
	defer ts.Close()

	svc := NewService(ts.URL, auth.AuthenticationParameter{ApiKey: "k", ApiSecret: "s"})
	res, err := svc.ListMessages(context.Background(), "g1", ListMessagesQuery{Limit: 10})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if res.Limit != 10 {
		t.Fatalf("unexpected limit: %d", res.Limit)
	}
}

