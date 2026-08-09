package ui

import (
	"testing"

	"github.com/RobPro/music-basicifier/internal/platform"
)

func TestBuildPreviewTextHandlesNilBundle(t *testing.T) {
	if got := buildPreviewText(nil); got != "" {
		t.Fatalf("expected empty preview for nil bundle, got %q", got)
	}
}

func TestBuildPreviewTextFormatsBundle(t *testing.T) {
	bundle := &platform.OutputBundle{QBASIC: "10 REM test", Adafruit: "let melody = []"}
	got := buildPreviewText(bundle)
	want := "10 REM test\n\nlet melody = []"
	if got != want {
		t.Fatalf("unexpected preview text: %q", got)
	}
}
