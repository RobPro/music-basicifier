package ui

import "testing"

func TestBuildEmptyPreviewText(t *testing.T) {
	got := buildEmptyPreviewText()
	want := "No output yet. Confirm an input to generate preview text."
	if got != want {
		t.Fatalf("unexpected empty preview text: %q", got)
	}
}
