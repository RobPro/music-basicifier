package ui

import "testing"

func TestBuildCopyTextUsesPreviewContent(t *testing.T) {
	preview := "preview output"
	if got := buildCopyText(preview); got != preview {
		t.Fatalf("expected copy text to match preview content, got %q", got)
	}
}
