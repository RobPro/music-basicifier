package platform

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// MelodyNote represents a simplified extracted note for downstream conversion.
type MelodyNote struct {
	Pitch    string
	Duration int
}

// MelodyData captures a simplified melody stream extracted from an audio file.
type MelodyData struct {
	SourcePath string
	Notes      []MelodyNote
}

// ExtractMelodyFromAudioFile creates a temporary melody representation from a local audio file.
// This is intentionally lightweight and serves as a placeholder for the future real audio decoder.
func ExtractMelodyFromAudioFile(path string) (*MelodyData, error) {
	if err := ValidateAudioFilePath(path); err != nil {
		return nil, fmt.Errorf("invalid audio file: %w", err)
	}

	info, err := os.Stat(path)
	if err != nil {
		return nil, fmt.Errorf("read audio file: %w", err)
	}
	if info.IsDir() {
		return nil, fmt.Errorf("audio file path points to a directory")
	}

	baseName := strings.TrimSuffix(filepath.Base(path), filepath.Ext(path))
	return &MelodyData{
		SourcePath: path,
		Notes: []MelodyNote{
			{Pitch: baseName, Duration: 1},
		},
	}, nil
}
