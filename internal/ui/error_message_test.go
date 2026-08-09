package ui

import (
	"errors"
	"testing"
)

func TestBuildErrorMessage(t *testing.T) {
	if got := buildErrorMessage(nil); got != "" {
		t.Fatalf("expected empty message for nil error, got %q", got)
	}
	if got := buildErrorMessage(errors.New("invalid audio file: missing")); got != "Please provide a readable .wav, .m4u, or .mp3 file." {
		t.Fatalf("unexpected user message: %q", got)
	}
	if got := buildErrorMessage(errors.New("could not extract melody: decode audio file")); got != "Could not extract melody from this audio file." {
		t.Fatalf("unexpected user message: %q", got)
	}
	if got := buildErrorMessage(errors.New("could not generate outputs: malformed melody")); got != "Could not generate output text from this melody." {
		t.Fatalf("unexpected user message: %q", got)
	}
	if got := buildErrorMessage(errors.New("could not copy output: clipboard failure")); got != "Could not copy output to the clipboard." {
		t.Fatalf("unexpected user message: %q", got)
	}
	if got := buildErrorMessage(errors.New("could not download YouTube audio: timeout")); got != "Could not download audio from the provided YouTube URL." {
		t.Fatalf("unexpected user message: %q", got)
	}
	if got := buildErrorMessage(errors.New("boom")); got != "Something went wrong. Please try again." {
		t.Fatalf("expected fallback user message, got %q", got)
	}
}
