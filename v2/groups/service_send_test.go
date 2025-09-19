package groups

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/solapi/solapi-go/v2/internal/auth"
)

func TestGroups_Send_BasicSuccess(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		if r.URL.Path != "/messages/v4/groups/g1/send" {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(map[string]any{
			"groupInfo": map[string]any{
				"count": map[string]any{"total": 1, "registeredFailed": 0},
			},
			"failedMessageList": []any{},
		})
	}))
	defer ts.Close()

	svc := NewService(ts.URL, auth.AuthenticationParameter{ApiKey: "k", ApiSecret: "s"})
	res, err := svc.Send(context.Background(), "g1")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if res.GroupInfo.Count.Total != 1 {
		t.Fatalf("unexpected total: %d", res.GroupInfo.Count.Total)
	}
}

