package platform

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestLogErrorRequiresMessage(t *testing.T) {
	if err := LogError(""); err == nil {
		t.Fatal("expected empty message to fail")
	}
}

func TestLogErrorWritesMessage(t *testing.T) {
	logPath := filepath.Join(os.TempDir(), "music-basicifier-errors.log")
	_ = os.Remove(logPath)

	message := "test error message"
	if err := LogError(message); err != nil {
		t.Fatalf("expected log write to succeed: %v", err)
	}

	content, err := os.ReadFile(logPath)
	if err != nil {
		t.Fatalf("expected log file to exist: %v", err)
	}
	if !strings.Contains(string(content), message) {
		t.Fatalf("expected log content to contain message %q, got %q", message, string(content))
	}
}
