package messages

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/solapi/solapi-go/v2/internal/auth"
)

func TestService_List_QueryCombinations(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		q := r.URL.Query()
		if q.Get("to") != "010" || q.Get("from") != "029" || q.Get("type") != "SMS,LMS" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if q.Get("startDate") != "2025-01-01" || q.Get("endDate") != "2025-02-01" || q.Get("dateType") != "CREATED" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(map[string]any{"messageList": map[string]any{}, "limit": 20})
	}))
	defer ts.Close()

	svc := NewService(ts.URL, auth.AuthenticationParameter{ApiKey: "k", ApiSecret: "s"})
	_, err := svc.List(context.Background(), ListQuery{
		To:        "010",
		From:      "029",
		TypeIn:    []string{"SMS", "LMS"},
		DateType:  "CREATED",
		StartDate: "2025-01-01",
		EndDate:   "2025-02-01",
	})
	if err != nil {
		t.Fatalf("unexpected err: %v", err)
	}
}

func TestService_List_ServerErrorBubbles(t *testing.T) {
	calls := 0
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		calls++
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer ts.Close()

	svc := NewService(ts.URL, auth.AuthenticationParameter{ApiKey: "k", ApiSecret: "s"})
	_, err := svc.List(context.Background(), ListQuery{Limit: 1})
	if err == nil {
		t.Fatalf("expected error, got nil")
	}
	if calls == 0 {
		t.Fatalf("server was not called")
	}
}
