package client

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/solapi/solapi-go/v2/groups"
)

func TestClient_WithHTTPClient_PropagatesToGroupsDirectCalls(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("X-From-Custom") != "1" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(groups.CreateGroupResponse{})
	}))
	defer ts.Close()

	rt := roundTripperFunc(func(req *http.Request) (*http.Response, error) {
		req.Header.Set("X-From-Custom", "1")
		return http.DefaultTransport.RoundTrip(req)
	})
	hc := &http.Client{Transport: rt}

	c := newClientWithBaseURL(ts.URL, "k", "s").WithHTTPClient(hc)

	// Direct service call should use custom client
	if _, err := c.Groups.Create(context.Background(), groups.CreateGroupOptions{}); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}
