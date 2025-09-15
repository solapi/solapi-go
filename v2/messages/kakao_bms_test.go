package messages

import (
	"encoding/json"
	"testing"
)

func TestKakaoOptions_BMS_JSON(t *testing.T) {
	m := Message{
		To: "01000000000",
		KakaoOptions: &KakaoOptions{
			PfID:       "pf",
			TemplateID: "tid",
			BMS: &KakaoBMSOptions{
				Targeting: "I",
			},
		},
	}

	b, err := json.Marshal(m)
	if err != nil {
		t.Fatalf("failed to marshal: %v", err)
	}

	var out map[string]any
	if err := json.Unmarshal(b, &out); err != nil {
		t.Fatalf("failed to unmarshal: %v", err)
	}

	ko, ok := out["kakaoOptions"].(map[string]any)
	if !ok {
		t.Fatalf("kakaoOptions missing or wrong type: %T", out["kakaoOptions"])
	}

	bms, ok := ko["bms"].(map[string]any)
	if !ok {
		t.Fatalf("bms missing or wrong type: %T", ko["bms"])
	}

	if v, ok := bms["targeting"].(string); !ok || v != "I" {
		t.Fatalf("unexpected targeting value: %v", bms["targeting"])
	}
}
