package ui

import "testing"

func TestBuildCopyText(t *testing.T) {
	preview := "demo preview"
	if got := buildCopyText(preview); got != preview {
		t.Fatalf("expected copy text to match preview, got %q", got)
	}
}
