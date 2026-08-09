package platform

import (
	"testing"
)

func TestCopyTextToClipboardRequiresText(t *testing.T) {
	err := CopyTextToClipboard("")
	if err == nil {
		t.Fatal("expected empty text to fail")
	}
}
