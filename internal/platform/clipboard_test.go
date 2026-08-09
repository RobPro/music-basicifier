package platform

import (
	"testing"
)

func TestCopyTextToClipboardRequiresText(t *testing.T) {
	err := validateClipboardText("")
	if err == nil {
		t.Fatal("expected empty text to fail")
	}
}

func TestCopyTextToClipboardAcceptsWhitespace(t *testing.T) {
	err := validateClipboardText("  ")
	if err != nil {
		t.Fatalf("expected whitespace text to be accepted, got %v", err)
	}
}
