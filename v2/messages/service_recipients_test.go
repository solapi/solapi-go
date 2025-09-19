package messages

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/solapi/solapi-go/v2/internal/auth"
)

func TestSendManyDetail_ToListMarshalsAsArray(t *testing.T) {
	var seen any
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var body map[string]any
		_ = json.NewDecoder(r.Body).Decode(&body)
		msgs, _ := body["messages"].([]any)
		if len(msgs) == 0 {
			t.Fatalf("no messages in payload")
		}
		first, _ := msgs[0].(map[string]any)
		seen = first["to"]
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(map[string]any{
			"groupInfo": map[string]any{
				"count": map[string]any{
					"total":            2,
					"registeredFailed": 0,
				},
			},
			"failedMessageList": []any{},
		})
	}))
	defer ts.Close()

	svc := NewService(ts.URL, auth.AuthenticationParameter{ApiKey: "k", ApiSecret: "s"})
	_, err := svc.Send(context.Background(), SendRequest{Messages: []Message{{ToList: []string{"010", "011"}}}})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if _, ok := seen.([]any); !ok {
		t.Fatalf("expected to to be array, got %T", seen)
	}
}

func TestSendManyDetail_EmptyRecipientReturnsError(t *testing.T) {
	svc := NewService("http://invalid", auth.AuthenticationParameter{ApiKey: "k", ApiSecret: "s"})
	_, err := svc.Send(context.Background(), SendRequest{Messages: []Message{{To: ""}}})
	if err == nil {
		t.Fatalf("expected error, got nil")
	}
}
