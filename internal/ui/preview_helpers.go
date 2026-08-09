// Package ui contains Fyne application setup and top-level window wiring.
package ui

import (
	"fmt"
	"strings"

	"github.com/RobPro/music-basicifier/internal/platform"
)

func buildPreviewText(bundle *platform.OutputBundle) string {
	if bundle == nil {
		return ""
	}
	return bundle.QBASIC + "\n\n" + bundle.Adafruit
}

func buildCopyText(preview string) (string, error) {
	if strings.TrimSpace(preview) == "" || preview == buildEmptyPreviewText() {
		return "", fmt.Errorf("no output available to copy")
	}
	return preview, nil
}

func buildErrorMessage(err error) string {
	if err == nil {
		return ""
	}

	msg := err.Error()
	switch {
	case strings.Contains(msg, "invalid audio file"):
		return "Please provide a readable .wav, .m4u, or .mp3 file."
	case strings.Contains(msg, "could not extract melody"):
		return "Could not extract melody from this audio file."
	case strings.Contains(msg, "could not generate outputs"):
		return "Could not generate output text from this melody."
	case strings.Contains(msg, "could not copy output"):
		return "Could not copy output to the clipboard."
	case strings.Contains(msg, "could not download YouTube audio"):
		return "Could not download audio from the provided YouTube URL."
	default:
		return "Something went wrong. Please try again."
	}
}

func logErrorWithContext(context string, err error) {
	if err == nil {
		return
	}
	_ = platform.LogError(fmt.Sprintf("%s: %v", context, err))
}

func buildEmptyPreviewText() string {
	return "No output yet. Confirm an input to generate preview text."
}

func buildPreviewStatusText(hasOutput bool) string {
	if hasOutput {
		return "Preview ready"
	}
	return "Waiting for input"
}
