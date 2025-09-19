package messages

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/solapi/solapi-go/v2/internal/auth"
)

func TestService_List_BasicSuccess(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		if r.URL.Path != "/messages/v4/list" {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		if r.URL.Query().Get("limit") == "" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(map[string]any{
			"messageList": map[string]any{
				"m4v": map[string]any{"messageId": "m4v", "to": "01000000000", "from": "01000000000"},
			},
			"limit":    1,
			"startKey": "",
			"nextKey":  "",
		})
	}))
	defer ts.Close()

	svc := NewService(ts.URL, auth.AuthenticationParameter{ApiKey: "k", ApiSecret: "s"})
	res, err := svc.List(context.Background(), ListQuery{Limit: 1})
	if err != nil {
		t.Fatalf("unexpected err: %v", err)
	}
	if res.Limit != 1 {
		t.Fatalf("unexpected limit: %d", res.Limit)
	}
	if len(res.MessageList) != 1 {
		t.Fatalf("unexpected messageList size: %d", len(res.MessageList))
	}
	if m, ok := res.MessageList["m4v"]; !ok || m.MessageID != "m4v" {
		t.Fatalf("unexpected messageList: %+v", res.MessageList)
	}
}
