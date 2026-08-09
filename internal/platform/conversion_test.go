package platform

import "testing"

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
