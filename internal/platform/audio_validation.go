package platform

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// ValidateAudioFilePath checks that a local audio path points to a readable file with a supported extension.
func ValidateAudioFilePath(path string) error {
	trimmed := strings.TrimSpace(path)
	if trimmed == "" {
		return fmt.Errorf("audio file path is required")
	}

	fileInfo, err := os.Stat(trimmed)
	if err != nil {
		return fmt.Errorf("audio file is not accessible: %w", err)
	}
	if fileInfo.IsDir() {
		return fmt.Errorf("audio file path points to a directory")
	}

	file, err := os.Open(trimmed)
	if err != nil {
		return fmt.Errorf("audio file is not readable: %w", err)
	}
	if err := file.Close(); err != nil {
		return fmt.Errorf("audio file is not readable: %w", err)
	}

	supportedExtensions := []string{".wav", ".m4u", ".mp3"}
	lowerPath := strings.ToLower(filepath.Ext(trimmed))
	for _, ext := range supportedExtensions {
		if lowerPath == ext {
			return nil
		}
	}

	return fmt.Errorf("unsupported audio format: expected .wav, .m4u, or .mp3")
}
