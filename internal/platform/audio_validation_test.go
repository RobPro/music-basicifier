package platform

import (
	"os"
	"path/filepath"
	"testing"
)

func TestValidateAudioFilePathAcceptsSupportedFile(t *testing.T) {
	tempDir := t.TempDir()
	filePath := filepath.Join(tempDir, "sample.wav")
	if err := os.WriteFile(filePath, []byte("audio"), 0o644); err != nil {
		t.Fatalf("failed to create temp audio file: %v", err)
	}

	if err := ValidateAudioFilePath(filePath); err != nil {
		t.Fatalf("expected supported audio file to pass validation: %v", err)
	}
}

func TestValidateAudioFilePathRejectsUnsupportedExtension(t *testing.T) {
	tempDir := t.TempDir()
	filePath := filepath.Join(tempDir, "sample.txt")
	if err := os.WriteFile(filePath, []byte("audio"), 0o644); err != nil {
		t.Fatalf("failed to create temp file: %v", err)
	}

	err := ValidateAudioFilePath(filePath)
	if err == nil {
		t.Fatal("expected unsupported extension to fail validation")
	}
	if got := err.Error(); got != "unsupported audio format" {
		t.Fatalf("unexpected error: %s", got)
	}
}
