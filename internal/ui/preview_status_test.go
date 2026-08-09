package ui

import "testing"

func TestBuildPreviewStatusText(t *testing.T) {
	if got := buildPreviewStatusText(false); got != "Waiting for input" {
		t.Fatalf("unexpected waiting text: %q", got)
	}
	if got := buildPreviewStatusText(true); got != "Preview ready" {
		t.Fatalf("unexpected ready text: %q", got)
	}
}
