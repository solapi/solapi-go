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

func TestClient_WithHTTPClient_PropagatesToTransport(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("X-From-Custom") != "1" {
			w.WriteHeader(http.StatusBadRequest)
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

	// custom client that injects header via RoundTripper
	rt := roundTripperFunc(func(req *http.Request) (*http.Response, error) {
		req.Header.Set("X-From-Custom", "1")
		return http.DefaultTransport.RoundTrip(req)
	})
	hc := &http.Client{Transport: rt}

	c := newClientWithBaseURL(ts.URL, "k", "s").WithHTTPClient(hc)
	_, err := c.Send(messages.Message{To: "010"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

type roundTripperFunc func(*http.Request) (*http.Response, error)

func (f roundTripperFunc) RoundTrip(r *http.Request) (*http.Response, error) { return f(r) }
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
	if calls != 0 {
		t.Fatalf("request should not be sent when recipient invalid; calls=%d", calls)
	}
}

// storages wiring test moved to storages_wiring_test.go
// messages shortcuts moved to messages_shortcuts_test.go
