package messages

import (
	"encoding/json"
	"testing"
)

func TestMessage_FaxOptions_JSON(t *testing.T) {
	m := Message{
		To:   "01000000000",
		From: "020000000",
		Type: "FAX",
		FaxOptions: &FaxOptions{
			FileIDs: []string{"fid1", "fid2"},
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

	fo, ok := out["faxOptions"].(map[string]any)
	if !ok {
		t.Fatalf("faxOptions missing or wrong type: %T", out["faxOptions"])
	}

	ids, ok := fo["fileIds"].([]any)
	if !ok || len(ids) != 2 {
		t.Fatalf("unexpected fileIds: %v", fo["fileIds"])
	}
}
