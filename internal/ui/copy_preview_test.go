package ui

import "testing"

func TestBuildCopyTextUsesPreviewContent(t *testing.T) {
	preview := "preview output"
	got, err := buildCopyText(preview)
	if err != nil {
		t.Fatalf("expected copy text to succeed: %v", err)
	}
	if got != preview {
		t.Fatalf("expected copy text to match preview content, got %q", got)
	}
}

func TestBuildCopyTextRejectsEmptyPreview(t *testing.T) {
	_, err := buildCopyText(buildEmptyPreviewText())
	if err == nil {
		t.Fatal("expected empty preview to fail")
	}
}
