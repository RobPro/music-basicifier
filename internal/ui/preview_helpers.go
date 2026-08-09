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
	return err.Error()
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
