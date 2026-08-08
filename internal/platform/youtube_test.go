package platform

import (
	"errors"
	"path/filepath"
	"testing"
)

func TestDownloadYouTubeAudioUsesExpectedCommand(t *testing.T) {
	var called bool
	var gotName string
	var gotArgs []string

	runner := func(name string, args ...string) ([]byte, error) {
		called = true
		gotName = name
		gotArgs = append([]string(nil), args...)
		return []byte("ok"), nil
	}

	destination := filepath.Join("tmp", "downloads")
	outputPath, err := downloadYouTubeAudio("https://example.com/video", destination, runner)
	if err != nil {
		t.Fatalf("downloadYouTubeAudio returned error: %v", err)
	}
	if !called {
		t.Fatal("expected command runner to be called")
	}
	if gotName != "yt-dlp" {
		t.Fatalf("unexpected command name: %s", gotName)
	}
	if len(gotArgs) != 6 {
		t.Fatalf("unexpected argument count: %d", len(gotArgs))
	}
	if gotArgs[0] != "--extract-audio" || gotArgs[1] != "--audio-format" || gotArgs[2] != "mp3" || gotArgs[3] != "--output" {
		t.Fatalf("unexpected args prefix: %v", gotArgs)
	}
	if gotArgs[4] != filepath.Join(destination, "youtube-audio.%(ext)s") {
		t.Fatalf("unexpected output path: %s", gotArgs[4])
	}
	if gotArgs[5] != "https://example.com/video" {
		t.Fatalf("unexpected URL arg: %s", gotArgs[5])
	}
	if outputPath != filepath.Join(destination, "youtube-audio.mp3") {
		t.Fatalf("unexpected output path returned: %s", outputPath)
	}
}

func TestDownloadYouTubeAudioReturnsHelpfulError(t *testing.T) {
	runner := func(name string, args ...string) ([]byte, error) {
		return nil, errors.New("executing yt-dlp failed")
	}

	_, err := downloadYouTubeAudio("https://example.com/video", "tmp", runner)
	if err == nil {
		t.Fatal("expected error for failing command")
	}
	if got := err.Error(); got != "download failed: executing yt-dlp failed" {
		t.Fatalf("unexpected error: %s", got)
	}
}
