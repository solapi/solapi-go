package messages

import (
	"encoding/json"
	"testing"
)

func TestMessage_VoiceOptions_JSON(t *testing.T) {
	m := Message{
		To:   "01000000000",
		From: "0200000000",
		Text: "테스트",
		Type: "VOICE",
		VoiceOptions: &VoiceOptions{
			VoiceType:  "FEMALE",
			ReplyRange: 3,
		},
	}

	b, err := json.Marshal(m)
	if err != nil {
		t.Fatalf("marshal failed: %v", err)
	}

	var out map[string]any
	if err := json.Unmarshal(b, &out); err != nil {
		t.Fatalf("unmarshal failed: %v", err)
	}

	vo, ok := out["voiceOptions"].(map[string]any)
	if !ok {
		t.Fatalf("voiceOptions missing or wrong type: %T", out["voiceOptions"])
	}

	if v, _ := vo["voiceType"].(string); v != "FEMALE" {
		t.Fatalf("unexpected voiceType: %v", vo["voiceType"])
	}
}
