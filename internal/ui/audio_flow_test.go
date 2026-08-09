package ui

import (
	"os"
	"path/filepath"
	"testing"
)

func TestBuildConversionInputFromAudioFile(t *testing.T) {
	tempDir := t.TempDir()
	filePath := filepath.Join(tempDir, "tone.wav")
	if err := os.WriteFile(filePath, []byte("audio"), 0o644); err != nil {
		t.Fatalf("failed to create temp audio file: %v", err)
	}

	input, err := buildConversionInputFromAudioFile(filePath)
	if err != nil {
		t.Fatalf("expected conversion input to be built: %v", err)
	}
	if input == nil || input.Melody == nil {
		t.Fatal("expected conversion input with melody data")
	}
	if len(input.Melody.Notes) == 0 {
		t.Fatal("expected at least one melody note")
	}
}
