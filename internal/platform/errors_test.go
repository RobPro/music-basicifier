package platform

import "testing"

func TestLogErrorRequiresMessage(t *testing.T) {
	if err := LogError(""); err == nil {
		t.Fatal("expected empty message to fail")
	}
}
