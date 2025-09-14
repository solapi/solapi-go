package messages

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/solapi/solapi-go/v2/internal/auth"
)

func TestSendManyDetail_AllFailedReturnsError(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		if r.URL.Path != "/messages/v4/send-many/detail" {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(map[string]any{
			"groupInfo": map[string]any{
				"count": map[string]any{
					"total":            2,
					"registeredFailed": 2,
				},
			},
			"failedMessageList": []any{map[string]any{}, map[string]any{}},
		})
	}))
	defer ts.Close()

	svc := NewService(ts.URL, auth.AuthenticationParameter{ApiKey: "k", ApiSecret: "s"})
	_, err := svc.Send(context.Background(), SendRequest{Messages: []Message{{To: "010"}}})
	if err == nil {
		t.Fatalf("expected error, got nil")
	}
	var mnre *MessageNotReceivedError
	if _, ok := err.(*MessageNotReceivedError); !ok {
		if errors.As(err, &mnre) {
			// ok
		} else {
			t.Fatalf("expected MessageNotReceivedError, got %T", err)
		}
	}
}

func TestSendManyDetail_PartialSuccessReturnsResponse(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(map[string]any{
			"groupInfo": map[string]any{
				"count": map[string]any{
					"total":            3,
					"registeredFailed": 1,
				},
			},
			"failedMessageList": []any{map[string]any{}},
		})
	}))
	defer ts.Close()

	svc := NewService(ts.URL, auth.AuthenticationParameter{ApiKey: "k", ApiSecret: "s"})
	res, err := svc.Send(context.Background(), SendRequest{Messages: []Message{{To: "010"}}})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if res.GroupInfo.Count.Total != 3 {
		t.Fatalf("unexpected total: %d", res.GroupInfo.Count.Total)
	}
}
