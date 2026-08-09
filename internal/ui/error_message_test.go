package ui

import (
	"errors"
	"testing"
)

func TestBuildErrorMessage(t *testing.T) {
	if got := buildErrorMessage(nil); got != "" {
		t.Fatalf("expected empty message for nil error, got %q", got)
	}
	if got := buildErrorMessage(errors.New("boom")); got != "boom" {
		t.Fatalf("expected error text to be returned, got %q", got)
	}
}
