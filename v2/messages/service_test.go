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

func TestSendManyDetail_FullResponseDeserialization(t *testing.T) {
	// Mock a full API response structure to test deserialization
	fullResponse := map[string]any{
		"groupInfo": map[string]any{
			"count": map[string]any{
				"total":             2,
				"sentTotal":         2,
				"sentFailed":        0,
				"sentSuccess":       2,
				"sentPending":       0,
				"sentReplacement":   0,
				"refund":            0,
				"registeredFailed":  0,
				"registeredSuccess": 2,
			},
			"countForCharge": map[string]any{
				"sms": map[string]any{
					"82": 2,
				},
			},
			"balance": map[string]any{
				"requested":   2.0,
				"replacement": 0.0,
				"additional":  0.0,
				"refund":      0.0,
				"sum":         2.0,
			},
			"app": map[string]any{
				"app":     "test-app",
				"version": "1.0.0",
				"profit": map[string]any{
					"sms": 0.1,
				},
			},
		},
		"failedMessageList": []any{},
		"messageList": []any{
			map[string]any{
				"messageId":     "test-message-id-1",
				"statusCode":    "2000",
				"statusMessage": "Success",
				"customFields": map[string]any{
					"field1": "value1",
				},
			},
		},
		"status":          "COMPLETE",
		"dateSent":        "2024-01-01T00:00:00.000Z",
		"scheduledDate":   "",
		"dateCompleted":   "2024-01-01T00:00:01.000Z",
		"isRefunded":      false,
		"flagUpdated":     false,
		"prepaid":         true,
		"strict":          false,
		"masterAccountId": "test-master-account",
		"allowDuplicates": false,
		"_id":             "test-group-id",
		"accountId":       "test-account-id",
		"apiVersion":      "v4",
		"customFields": map[string]any{
			"groupField": "groupValue",
		},
		"hint":    "",
		"groupId": "test-group-id",
		"price": map[string]any{
			"KR": map[string]any{
				"sms": 0.1,
			},
		},
		"serviceMethod": "send-many/detail",
		"sdkVersion":    "go/2.0.0",
		"osPlatform":    "darwin | go1.21.0",
		"log": []any{
			map[string]any{
				"createAt":   "2024-01-01T00:00:00.000Z",
				"message":    "Message sent successfully",
				"oldBalance": 100.0,
				"newBalance": 98.0,
				"oldPoint":   0.0,
				"newPoint":   0.0,
				"totalPrice": 2.0,
			},
		},
		"dateCreated": "2024-01-01T00:00:00.000Z",
		"dateUpdated": "2024-01-01T00:00:01.000Z",
	}

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(fullResponse)
	}))
	defer ts.Close()

	svc := NewService(ts.URL, auth.AuthenticationParameter{ApiKey: "k", ApiSecret: "s"})
	res, err := svc.Send(context.Background(), SendRequest{Messages: []Message{{To: "010", Text: "test"}}})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Test key fields to verify deserialization worked
	if res.Status != "COMPLETE" {
		t.Errorf("expected status COMPLETE, got %s", res.Status)
	}
	if res.GroupID != "test-group-id" {
		t.Errorf("expected groupId test-group-id, got %s", res.GroupID)
	}
	if res.AccountID != "test-account-id" {
		t.Errorf("expected accountId test-account-id, got %s", res.AccountID)
	}
	if res.GroupInfo.Count.Total != 2 {
		t.Errorf("expected total 2, got %d", res.GroupInfo.Count.Total)
	}
	if len(res.MessageList) != 1 {
		t.Errorf("expected 1 message in list, got %d", len(res.MessageList))
	}
	if res.MessageList[0].MessageID != "test-message-id-1" {
		t.Errorf("expected messageId test-message-id-1, got %s", res.MessageList[0].MessageID)
	}

	// Print the response to see what was actually deserialized
	b, _ := json.MarshalIndent(res, "", "  ")
	t.Logf("Deserialized response:\n%s", string(b))
}

func TestSendManyDetail_IDFieldIssue(t *testing.T) {
	// Test potential issue with _id field vs id field
	responseWithId := map[string]any{
		"groupInfo": map[string]any{
			"count": map[string]any{
				"total": 1,
			},
		},
		"failedMessageList": []any{},
		"id":                "test-group-id-with-regular-id", // Using "id" instead of "_id"
		"accountId":         "test-account-id",
		"groupId":           "test-group-id",
	}

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(responseWithId)
	}))
	defer ts.Close()

	svc := NewService(ts.URL, auth.AuthenticationParameter{ApiKey: "k", ApiSecret: "s"})
	res, err := svc.Send(context.Background(), SendRequest{Messages: []Message{{To: "010", Text: "test"}}})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if res.ID == "" {
		t.Errorf("ID field not deserialized, expected some value, got empty string")
	}

	b, _ := json.MarshalIndent(res, "", "  ")
	t.Logf("Response with 'id' field:\n%s", string(b))
}

func TestSendManyDetail_PriceFieldInGroupInfo(t *testing.T) {
	// Test price field when it's in groupInfo (actual API structure)
	responseWithPriceInGroupInfo := map[string]any{
		"groupInfo": map[string]any{
			"count": map[string]any{
				"total": 1,
			},
			"price": map[string]any{
				"82": map[string]any{
					"sms": 10.5,
					"lms": 20.0,
				},
			},
		},
		"failedMessageList": []any{},
	}

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(responseWithPriceInGroupInfo)
	}))
	defer ts.Close()

	svc := NewService(ts.URL, auth.AuthenticationParameter{ApiKey: "k", ApiSecret: "s"})
	res, err := svc.Send(context.Background(), SendRequest{Messages: []Message{{To: "010", Text: "test"}}})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Price should be properly deserialized from groupInfo
	if len(res.Price) == 0 {
		t.Errorf("Expected Price field to contain value from groupInfo")
	}
	if country82Price, exists := res.Price["82"]; !exists {
		t.Errorf("Expected 82 key in Price map")
	} else if country82Price.SMS != 10.5 {
		t.Errorf("Expected SMS price to be 10.5, got %f", country82Price.SMS)
	}

	b, _ := json.MarshalIndent(res, "", "  ")
	t.Logf("Response with price in groupInfo:\n%s", string(b))
}

func TestSendManyDetail_RealAPIResponse(t *testing.T) {
	// Test with actual API response structure that was seen in terminal
	realAPIResponse := map[string]any{
		"groupInfo": map[string]any{
			"count": map[string]any{
				"total":             1,
				"sentTotal":         0,
				"sentFailed":        0,
				"sentSuccess":       0,
				"sentPending":       0,
				"sentReplacement":   0,
				"refund":            0,
				"registeredFailed":  0,
				"registeredSuccess": 1,
			},
			"countForCharge": map[string]any{
				"sms": map[string]any{
					"82": 1,
				},
			},
			"balance": map[string]any{
				"requested":   7.8,
				"replacement": 0.0,
				"additional":  0.0,
				"refund":      0.0,
				"sum":         7.8,
			},
			"point": map[string]any{
				"requested":   0.0,
				"replacement": 0.0,
				"additional":  0.0,
				"refund":      0.0,
				"sum":         0.0,
			},
			"app": map[string]any{
				"profit": map[string]any{
					"sms":                   0.0,
					"lms":                   0.0,
					"mms":                   0.0,
					"ata":                   0.0,
					"cta":                   0.0,
					"cti":                   0.0,
					"nsa":                   0.0,
					"rcs_sms":               0.0,
					"rcs_lms":               0.0,
					"rcs_mms":               0.0,
					"rcs_tpl":               0.0,
					"rcs_itpl":              0.0,
					"rcs_ltpl":              0.0,
					"fax":                   0.0,
					"voice":                 0.0,
					"bms_text":              0.0,
					"bms_image":             0.0,
					"bms_wide":              0.0,
					"bms_wide_item_list":    0.0,
					"bms_carousel_feed":     0.0,
					"bms_premium_video":     0.0,
					"bms_commerce":          0.0,
					"bms_carousel_commerce": 0.0,
				},
				"app":     "",
				"version": "",
			},
		},
		"failedMessageList": []any{},
		"status":            "",
		"dateSent":          "",
		"scheduledDate":     "",
		"dateCompleted":     "",
		"isRefunded":        false,
		"flagUpdated":       false,
		"prepaid":           false,
		"strict":            false,
		"masterAccountId":   "",
		"allowDuplicates":   false,
		"accountId":         "",
		"apiVersion":        "",
		"customFields":      nil,
		"hint":              "",
		"groupId":           "",
		"price":             nil,
		"dateCreated":       "",
		"dateUpdated":       "",
	}

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(realAPIResponse)
	}))
	defer ts.Close()

	svc := NewService(ts.URL, auth.AuthenticationParameter{ApiKey: "k", ApiSecret: "s"})
	res, err := svc.Send(context.Background(), SendRequest{Messages: []Message{{To: "010", Text: "test"}}})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Test that all the fields from the real API response are properly deserialized
	if res.GroupInfo.Count.Total != 1 {
		t.Errorf("Expected total 1, got %d", res.GroupInfo.Count.Total)
	}
	if res.GroupInfo.Count.RegisteredSuccess != 1 {
		t.Errorf("Expected registeredSuccess 1, got %d", res.GroupInfo.Count.RegisteredSuccess)
	}
	if res.GroupInfo.Balance.Requested != 7.8 {
		t.Errorf("Expected balance.requested 7.8, got %f", res.GroupInfo.Balance.Requested)
	}

	// These fields should be empty strings or default values as shown in the terminal output
	if res.Status != "" {
		t.Errorf("Expected status to be empty string, got %s", res.Status)
	}
	if res.AccountID != "" {
		t.Errorf("Expected accountId to be empty string, got %s", res.AccountID)
	}

	b, _ := json.MarshalIndent(res, "", "  ")
	t.Logf("Real API response deserialization:\n%s", string(b))
}
