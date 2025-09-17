package groups

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/solapi/solapi-go/v2/internal/auth"
	"github.com/solapi/solapi-go/v2/messages"
)

func TestGroups_AddMessages_BasicSuccess(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPut {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		if r.URL.Path != "/messages/v4/groups/g1/messages" {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		var body map[string]any
		_ = json.NewDecoder(r.Body).Decode(&body)
		msgs, ok := body["messages"].([]any)
		if !ok || len(msgs) != 2 {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(map[string]any{
			"groupInfo": map[string]any{
				"count": map[string]any{"total": 2, "registeredFailed": 0},
			},
			"failedMessageList": []any{},
		})
	}))
	defer ts.Close()

	svc := NewService(ts.URL, auth.AuthenticationParameter{ApiKey: "k", ApiSecret: "s"})
	res, err := svc.AddMessages(context.Background(), "g1", AddGroupMessagesRequest{
		Messages: []messages.Message{{To: "010"}, {To: "011"}},
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if res.GroupInfo.Count.Total != 2 {
		t.Fatalf("unexpected total: %d", res.GroupInfo.Count.Total)
	}
}

