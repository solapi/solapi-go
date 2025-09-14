package apirequest

import (
	"encoding/hex"
	"os"
	"testing"
)

func TestRandomStringLength(t *testing.T) {
	n := 20
	s := RandomString(n)
	if len(s) != n*2 {
		t.Fatalf("expected hex string length %d, got %d", n*2, len(s))
	}
	if _, err := hex.DecodeString(s); err != nil {
		t.Fatalf("expected valid hex string, decode failed: %v", err)
	}
}

func TestExistsFalseForMissingFile(t *testing.T) {
	ok, err := exists("definitely-not-present-12345.txt")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if ok {
		t.Fatalf("expected file to not exist")
	}
}

func TestExistsTrueForExistingFile(t *testing.T) {
	f, err := os.Create("temp_test_file.txt")
	if err != nil {
		t.Fatalf("failed to create temp file: %v", err)
	}
	_ = f.Close()
	defer os.Remove("temp_test_file.txt")

	ok, err := exists("temp_test_file.txt")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !ok {
		t.Fatalf("expected file to exist")
	}
}
