package messages

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/solapi/solapi-go/v2/internal/auth"
)

func TestSendManyDetail_FastResponseWithoutCtaCti(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		if r.URL.Path != "/messages/v4/send-many/fast" {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(map[string]any{
			"groupInfo": map[string]any{
				"_id":    "G4V-fast-1",
				"status": "SEND-REQUESTED",
				"count": map[string]any{
					"total":             1,
					"registeredSuccess": 1,
				},
				"countForCharge": map[string]any{
					"sms": map[string]any{"82": 1},
					"ata": map[string]any{},
				},
				"app": map[string]any{
					"profit": map[string]any{"sms": 0, "lms": 0, "ata": 0},
				},
			},
			"failedMessageList": []any{},
			"messageList": []any{
				map[string]any{
					"messageId":     "M4V-1",
					"statusCode":    "2000",
					"statusMessage": "정상 접수(이통사로 접수 예정) ",
					"customFields":  map[string]any{"k": "v"},
				},
			},
		})
	}))
	defer ts.Close()

	svc := NewService(ts.URL, auth.AuthenticationParameter{ApiKey: "k", ApiSecret: "s"})
	show := true
	res, err := svc.Send(context.Background(), Message{To: "010", From: "010", Text: "hi"}, SendOptions{ShowMessageList: &show})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if res.Status != "SEND-REQUESTED" {
		t.Fatalf("unexpected status: %q", res.Status)
	}
	if res.GroupInfo.CountForCharge.CTA != nil || res.GroupInfo.CountForCharge.CTI != nil {
		t.Fatalf("expected cta/cti to be nil when absent")
	}
	if res.GroupInfo.App.Profit.CTA != 0 || res.GroupInfo.App.Profit.CTI != 0 {
		t.Fatalf("expected profit cta/cti zero when absent")
	}
	if len(res.MessageList) != 1 || res.MessageList[0].CustomFields["k"] != "v" {
		t.Fatalf("unexpected messageList: %+v", res.MessageList)
	}
}
