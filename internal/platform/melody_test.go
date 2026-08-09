package platform

import (
	"os"
	"path/filepath"
	"testing"
)

func TestExtractMelodyFromAudioFileReturnsMelodyData(t *testing.T) {
	tempDir := t.TempDir()
	filePath := filepath.Join(tempDir, "sample.wav")
	if err := os.WriteFile(filePath, []byte("audio"), 0o644); err != nil {
		t.Fatalf("failed to create temp audio file: %v", err)
	}

	data, err := ExtractMelodyFromAudioFile(filePath)
	if err != nil {
		t.Fatalf("expected melody extraction to succeed: %v", err)
	}
	if data == nil {
		t.Fatal("expected melody data to be returned")
	}
	if data.SourcePath != filePath {
		t.Fatalf("unexpected source path: %s", data.SourcePath)
	}
	if len(data.Notes) != 1 {
		t.Fatalf("expected a single melody note, got %d", len(data.Notes))
	}
	if got := data.Notes[0].Pitch; got != "sample" {
		t.Fatalf("unexpected pitch: %s", got)
	}
}

func TestExtractMelodyFromAudioFileRejectsInvalidPath(t *testing.T) {
	_, err := ExtractMelodyFromAudioFile("missing.wav")
	if err == nil {
		t.Fatal("expected invalid path to fail")
	}
}
