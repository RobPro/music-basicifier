package platform

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// DownloadYouTubeAudio downloads YouTube audio using yt-dlp and returns the local path to the result.
func DownloadYouTubeAudio(url, destinationDir string) (string, error) {
	return downloadYouTubeAudio(url, destinationDir, func(name string, args ...string) ([]byte, error) {
		cmd := exec.Command(name, args...)
		var stderr bytes.Buffer
		cmd.Stderr = &stderr
		output, err := cmd.CombinedOutput()
		if err != nil {
			return output, fmt.Errorf("%s: %w", strings.TrimSpace(stderr.String()), err)
		}
		return output, nil
	})
}

func downloadYouTubeAudio(url, destinationDir string, runner func(name string, args ...string) ([]byte, error)) (string, error) {
	if strings.TrimSpace(url) == "" {
		return "", fmt.Errorf("youtube URL is required")
	}
	if err := os.MkdirAll(destinationDir, 0o755); err != nil {
		return "", fmt.Errorf("create download directory: %w", err)
	}

	outputPattern := filepath.Join(destinationDir, "youtube-audio.%(ext)s")
	output, err := runner("yt-dlp", "--extract-audio", "--audio-format", "mp3", "--output", outputPattern, url)
	if err != nil {
		return "", fmt.Errorf("download failed: %w", err)
	}

	_ = output
	return filepath.Join(destinationDir, "youtube-audio.mp3"), nil
}
