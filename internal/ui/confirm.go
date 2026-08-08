package ui

import (
	"fmt"
	"strings"
)

func buildConfirmationMessage(url, audioFile string) string {
	parts := []string{
		"Input received.",
		fmt.Sprintf("YouTube URL: %s", strings.TrimSpace(url)),
		fmt.Sprintf("Audio file: %s", strings.TrimSpace(audioFile)),
	}
	return strings.Join(parts, "\n")
}
