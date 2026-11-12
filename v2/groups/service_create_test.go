package groups

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/solapi/solapi-go/v2/internal/auth"
)

func TestGroups_Create_BasicSuccess(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		if r.URL.Path != "/messages/v4/groups" {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		var body map[string]any
		_ = json.NewDecoder(r.Body).Decode(&body)
		if v, _ := body["allowDuplicates"].(bool); !v {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(map[string]any{
			"groupId": "g1",
		})
	}))
	defer ts.Close()

	svc := NewService(ts.URL, auth.AuthenticationParameter{ApiKey: "k", ApiSecret: "s"})
	res, err := svc.Create(context.Background(), CreateGroupOptions{AllowDuplicates: true, AppId: "app"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if res.GroupID != "g1" {
		t.Fatalf("unexpected groupId: %s", res.GroupID)
	}
}

func TestGroups_Create_FlatGroupDocument(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(flatCreateGroupBody))
	}))
	defer ts.Close()

	svc := NewService(ts.URL, auth.AuthenticationParameter{ApiKey: "k", ApiSecret: "s"})
	res, err := svc.Create(context.Background(), CreateGroupOptions{})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if res.GroupID != "g3" || res.GroupInfo.GroupID != "g3" {
		t.Fatalf("unexpected groupId: %q / %q", res.GroupID, res.GroupInfo.GroupID)
	}
	if res.GroupInfo.Status != "PENDING" {
		t.Fatalf("unexpected status: %q", res.GroupInfo.Status)
	}
	if res.GroupInfo.APIVersion != "4" {
		t.Fatalf("unexpected apiVersion: %q", res.GroupInfo.APIVersion)
	}
	if res.GroupInfo.Count.Total != 0 || res.GroupInfo.AccountID != "26041322754777" {
		t.Fatalf("unexpected decode: %+v", res.GroupInfo)
	}
}

const flatCreateGroupBody = `{"groupId":"g3","serviceMethod":"MT","accountId":"26041322754777","apiVersion":"4","sdkVersion":"go/2.0.0","osPlatform":"darwin | go1.27.3","count":{"total":0,"sentTotal":0,"sentFailed":0,"sentSuccess":0,"sentPending":0,"sentReplacement":0,"refund":0,"registeredFailed":0,"registeredSuccess":0},"countForCharge":{"sms":{},"lms":{},"mms":{},"ata":{},"cta":{},"cti":{},"nsa":{},"rcs_sms":{},"rcs_lms":{},"rcs_mms":{},"rcs_tpl":{},"rcs_itpl":{},"rcs_ltpl":{},"fax":{},"voice":{},"bms_text":{},"bms_image":{},"bms_wide":{},"bms_wide_item_list":{},"bms_carousel_feed":{},"bms_premium_video":{},"bms_commerce":{},"bms_carousel_commerce":{},"bms_free":{}},"log":[{"createAt":"2026-11-12T06:33:15.388Z","message":"[52.79.119.88 / AS16509 Amazon.com, Inc.] 메시지 그룹이 생성되었습니다."}],"status":"PENDING","dateSent":null,"scheduledDate":null,"dateCompleted":null,"balance":{"requested":0,"replacement":0,"additional":0,"refund":0,"sum":0},"point":{"requested":0,"replacement":0,"additional":0,"refund":0,"sum":0},"deposit":{"requested":0,"replacement":0,"additional":0,"refund":0,"sum":0},"isRefunded":false,"app":{"profit":{"sms":0,"lms":0,"mms":0,"ata":0,"cta":0,"cti":0,"nsa":0,"rcs_sms":0,"rcs_lms":0,"rcs_mms":0,"rcs_tpl":0,"rcs_itpl":0,"rcs_ltpl":0,"fax":0,"voice":0,"bms_text":0,"bms_image":0,"bms_wide":0,"bms_wide_item_list":0,"bms_carousel_feed":0,"bms_premium_video":0,"bms_commerce":0,"bms_carousel_commerce":0,"bms_free":0},"appId":null,"version":null},"flagUpdated":false,"customFields":{},"prepaid":true,"strict":false,"hint":null,"masterAccountId":null,"allowDuplicates":false,"originCode":null,"workspaceId":null,"price":{},"dateCreated":"2026-11-12T06:33:15.389Z","dateUpdated":"2026-11-12T06:33:15.389Z","_id":"g3"}`
