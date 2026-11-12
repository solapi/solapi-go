package groups

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/solapi/solapi-go/v2/internal/auth"
)

func TestGroups_GetGroup_FastPath(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		if r.URL.Path != "/messages/v4/groups/fast/g1" {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(flatGroupBody))
	}))
	defer ts.Close()

	svc := NewService(ts.URL, auth.AuthenticationParameter{ApiKey: "k", ApiSecret: "s"})
	res, err := svc.GetGroup(context.Background(), "g1")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if res.GroupID != "g1" {
		t.Fatalf("unexpected groupId: %q", res.GroupID)
	}
	if res.GroupInfo.Status != "SEND-REQUESTED" {
		t.Fatalf("unexpected status: %q", res.GroupInfo.Status)
	}
	if res.GroupInfo.Count.Total != 1 || res.GroupInfo.Count.RegisteredSuccess != 1 {
		t.Fatalf("unexpected count: %+v", res.GroupInfo.Count)
	}
	if res.GroupInfo.APIVersion != "v4-fast" {
		t.Fatalf("unexpected apiVersion: %q", res.GroupInfo.APIVersion)
	}
	if res.GroupInfo.CountForCharge.SMS["82"] != 1 {
		t.Fatalf("unexpected countForCharge: %+v", res.GroupInfo.CountForCharge)
	}
	if res.GroupInfo.DateCreated != "2026-11-12T06:19:02.078Z" {
		t.Fatalf("unexpected dateCreated: %q", res.GroupInfo.DateCreated)
	}
}

func TestGroupResponse_UnmarshalWrappedGroupInfo(t *testing.T) {
	var res GroupResponse
	if err := json.Unmarshal([]byte(`{"groupId":"g2","groupInfo":{"status":"SENDING","count":{"total":3}}}`), &res); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if res.GroupID != "g2" || res.GroupInfo.Status != "SENDING" || res.GroupInfo.Count.Total != 3 {
		t.Fatalf("unexpected decode: %+v", res)
	}
}

const flatGroupBody = `{"_id":"g1","groupId":"g1","serviceMethod":"MT","accountId":"26041322754777","apiVersion":"v4-fast","sdkVersion":"go/2.0.0","osPlatform":"darwin | go1.27.3","count":{"total":1,"sentTotal":0,"sentFailed":0,"sentSuccess":0,"sentPending":0,"sentReplacement":0,"refund":0,"registeredFailed":0,"registeredSuccess":1},"countForCharge":{"sms":{"82":1},"lms":{},"mms":{},"ata":{},"nsa":{},"rcs_sms":{},"rcs_lms":{},"rcs_mms":{},"rcs_tpl":{},"rcs_itpl":{},"rcs_ltpl":{},"fax":{},"voice":{},"bms_text":{},"bms_image":{},"bms_wide":{},"bms_wide_item_list":{},"bms_carousel_feed":{},"bms_premium_video":{},"bms_commerce":{},"bms_carousel_commerce":{},"bms_free":{},"cta":{},"cti":{}},"log":[{"createAt":"2026-11-12T06:19:02.078Z","message":"메시지 그룹이 생성되었습니다."},{"createAt":"2026-11-12T06:19:02.078Z","message":"메시지를 발송했습니다.","oldBalance":99988496,"newBalance":99988496,"oldPoint":980514624,"newPoint":980514624,"oldDeposit":254509.09375,"newDeposit":254509.09375,"totalPrice":13}],"status":"SEND-REQUESTED","dateSent":"2026-11-12T06:19:02.078Z","scheduledDate":null,"dateCompleted":null,"price":{"82":{"sms":13,"lms":29,"mms":60,"ata":8,"cta":14,"cti":22,"nsa":6,"fax":100,"voice":200}},"balance":{"requested":0,"replacement":0,"additional":0,"refund":0,"sum":0},"point":{"requested":13,"replacement":0,"additional":0,"refund":0,"sum":13},"deposit":{"requested":0,"replacement":0,"additional":0,"refund":0,"sum":0},"app":{"appId":null,"version":null,"profit":{"sms":0,"lms":0,"mms":0,"ata":0,"nsa":0,"fax":0,"voice":0,"cta":0,"cti":0}},"customFields":{},"prepaid":true,"strict":false,"allowDuplicates":false,"originCode":null,"masterAccountId":null,"workspaceId":null,"isRefunded":false,"flagUpdated":false,"hint":{},"dateCreated":"2026-11-12T06:19:02.078Z","dateUpdated":"2026-11-12T06:19:02.078Z"}`
