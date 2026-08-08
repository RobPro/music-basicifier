package ui

import "testing"

func TestBuildConfirmationMessage(t *testing.T) {
	url := "https://example.com/video"
	audio := "C:/music/example.wav"

	message := buildConfirmationMessage(url, audio)

	if message == "" {
		t.Fatal("expected confirmation message to be populated")
	}

	if got := message; got != "Input received.\nYouTube URL: https://example.com/video\nAudio file: C:/music/example.wav" {
		t.Fatalf("unexpected confirmation message: %q", got)
	}
}
