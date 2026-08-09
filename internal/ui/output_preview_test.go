package ui

import (
	"testing"

	"github.com/RobPro/music-basicifier/internal/platform"
)

func TestBuildPreviewText(t *testing.T) {
	input := &platform.ConversionInput{Melody: &platform.MelodyData{Notes: []platform.MelodyNote{{Pitch: "C", Duration: 1}}}}
	bundle, err := platform.GenerateOutputBundle(input)
	if err != nil {
		t.Fatalf("expected output bundle to build: %v", err)
	}

	preview := buildPreviewText(bundle)
	if preview == "" {
		t.Fatal("expected preview text to be populated")
	}
	if got := preview; got != bundle.QBASIC+"\n\n"+bundle.Adafruit {
		t.Fatalf("unexpected preview text: %q", got)
	}
}
