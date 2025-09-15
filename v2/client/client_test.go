package client

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/solapi/solapi-go/v2/messages"
)

func TestClient_MessagesSend_WiresThrough(t *testing.T) {
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
					"total":            1,
					"registeredFailed": 0,
				},
			},
			"failedMessageList": []any{},
		})
	}))
	defer ts.Close()

	c := newClientWithBaseURL(ts.URL, "k", "s")
	if c.Messages == nil {
		t.Fatalf("Messages service is nil")
	}

	res, err := c.Messages.Send(context.Background(), messages.SendRequest{Messages: []messages.Message{{To: "010"}}})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if res.GroupInfo.Count.Total != 1 {
		t.Fatalf("unexpected total: %d", res.GroupInfo.Count.Total)
	}
}

func TestClient_Send_Shortcut(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(map[string]any{
			"groupInfo": map[string]any{
				"count": map[string]any{
					"total":            1,
					"registeredFailed": 0,
				},
			},
			"failedMessageList": []any{},
		})
	}))
	defer ts.Close()

	c := newClientWithBaseURL(ts.URL, "k", "s")
	res, err := c.Send(messages.Message{To: "010"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if res.GroupInfo.Count.Total != 1 {
		t.Fatalf("unexpected total: %d", res.GroupInfo.Count.Total)
	}
}

func TestClient_Send_ErrOnEmptyRecipientInList(t *testing.T) {
	calls := 0
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		calls++
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer ts.Close()

	c := newClientWithBaseURL(ts.URL, "k", "s")
	_, err := c.Send(messages.Message{ToList: []string{"010", ""}})
	if err == nil {
		t.Fatalf("expected error for empty recipient in list, got nil")
	}
	if err.Error() != "recipient contains empty value" {
		t.Fatalf("unexpected error: %v", err)
	}
	if calls != 0 {
		t.Fatalf("request should not be sent when recipient invalid; calls=%d", calls)
	}
}

func TestClient_List_Shortcut(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(map[string]any{
			"messageList": map[string]any{},
			"limit":       5,
		})
	}))
	defer ts.Close()

	c := newClientWithBaseURL(ts.URL, "k", "s")
	res, err := c.List(messages.ListQuery{Limit: 5})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if res.Limit != 5 {
		t.Fatalf("unexpected limit: %d", res.Limit)
	}
}
