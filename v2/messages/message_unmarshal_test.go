package messages

import (
	"encoding/json"
	"testing"
)

// Red: to가 배열일 때도 Message로 역직렬화가 가능해야 한다
func TestMessage_Unmarshal_ToArray(t *testing.T) {
	data := []byte(`{"to":["01011112222","01033334444"],"from":"029","text":"hi"}`)
	var m Message
	if err := json.Unmarshal(data, &m); err != nil {
		t.Fatalf("unexpected unmarshal error: %v", err)
	}
	if m.To != "" {
		t.Fatalf("expected To to be empty when to is array, got %q", m.To)
	}
	if len(m.ToList) != 2 || m.ToList[0] != "01011112222" || m.ToList[1] != "01033334444" {
		t.Fatalf("unexpected ToList: %+v", m.ToList)
	}
}

// Safety: 기존 단일 문자열 케이스도 유지되어야 한다
func TestMessage_Unmarshal_ToString(t *testing.T) {
	data := []byte(`{"to":"01055556666","from":"029","text":"hi"}`)
	var m Message
	if err := json.Unmarshal(data, &m); err != nil {
		t.Fatalf("unexpected unmarshal error: %v", err)
	}
	if m.To != "01055556666" {
		t.Fatalf("unexpected To: %q", m.To)
	}
	if len(m.ToList) != 0 {
		t.Fatalf("expected empty ToList, got: %+v", m.ToList)
	}
}
