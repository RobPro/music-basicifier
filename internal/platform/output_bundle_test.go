package platform

import (
	"strings"
	"testing"
)

func TestGenerateOutputBundleRequiresInput(t *testing.T) {
	_, err := GenerateOutputBundle(nil)
	if err == nil {
		t.Fatal("expected error when conversion input is nil")
	}
}

func TestGenerateOutputBundleBuildsBothOutputs(t *testing.T) {
	input := &ConversionInput{Melody: &MelodyData{Notes: []MelodyNote{{Pitch: "C", Duration: 1}}}}
	bundle, err := GenerateOutputBundle(input)
	if err != nil {
		t.Fatalf("expected output bundle to be generated: %v", err)
	}
	if bundle == nil {
		t.Fatal("expected output bundle to be returned")
	}
	if bundle.QBASIC == "" || bundle.Adafruit == "" {
		t.Fatal("expected both outputs to be populated")
	}
	if !strings.Contains(bundle.Adafruit, "music.playTone(") {
		t.Fatalf("expected Adafruit output to use playTone, got %q", bundle.Adafruit)
	}
}

func TestGenerateOutputBundleWrapsQBasicError(t *testing.T) {
	input := &ConversionInput{Melody: &MelodyData{Notes: []MelodyNote{{Pitch: "", Duration: 1}}}}
	_, err := GenerateOutputBundle(input)
	if err == nil {
		t.Fatal("expected QBASIC generation to fail")
	}
	if !strings.Contains(err.Error(), "generate QBASIC output") {
		t.Fatalf("unexpected error: %v", err)
	}
}
