package platform

import (
	"strings"
	"testing"
)

func TestBuildConversionInputRequiresMelodyData(t *testing.T) {
	_, err := BuildConversionInput(nil)
	if err == nil {
		t.Fatal("expected error when melody data is nil")
	}
}

func TestBuildConversionInputAcceptsMelodyData(t *testing.T) {
	melody := &MelodyData{Notes: []MelodyNote{{Pitch: "C", Duration: 1}}}
	input, err := BuildConversionInput(melody)
	if err != nil {
		t.Fatalf("expected conversion input to be built: %v", err)
	}
	if input == nil || input.Melody != melody {
		t.Fatal("expected conversion input to retain melody data")
	}
}

func TestBuildConversionInputRejectsEmptyPitch(t *testing.T) {
	_, err := BuildConversionInput(&MelodyData{Notes: []MelodyNote{{Pitch: "", Duration: 1}}})
	if err == nil {
		t.Fatal("expected error for empty pitch")
	}
	if !strings.Contains(err.Error(), "empty pitch") {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestBuildConversionInputRejectsInvalidDuration(t *testing.T) {
	_, err := BuildConversionInput(&MelodyData{Notes: []MelodyNote{{Pitch: "C", Duration: 0}}})
	if err == nil {
		t.Fatal("expected error for invalid duration")
	}
	if !strings.Contains(err.Error(), "invalid duration") {
		t.Fatalf("unexpected error: %v", err)
	}
}
