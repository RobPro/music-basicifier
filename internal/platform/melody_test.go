package platform

import (
	"bytes"
	"encoding/binary"
	"math"
	"os"
	"path/filepath"
	"testing"
)

func TestExtractMelodyFromAudioFileReturnsMelodyData(t *testing.T) {
	tempDir := t.TempDir()
	filePath := filepath.Join(tempDir, "sample.wav")
	if err := os.WriteFile(filePath, []byte("audio"), 0o644); err != nil {
		t.Fatalf("failed to create temp audio file: %v", err)
	}

	data, err := ExtractMelodyFromAudioFile(filePath)
	if err != nil {
		t.Fatalf("expected melody extraction to succeed: %v", err)
	}
	if data == nil {
		t.Fatal("expected melody data to be returned")
	}
	if data.SourcePath != filePath {
		t.Fatalf("unexpected source path: %s", data.SourcePath)
	}
	if len(data.Notes) != 1 {
		t.Fatalf("expected a single melody note, got %d", len(data.Notes))
	}
	if got := data.Notes[0].Pitch; got != "sample" {
		t.Fatalf("unexpected pitch: %s", got)
	}
}

func TestExtractMelodyFromAudioFileParsesSimpleWav(t *testing.T) {
	tempDir := t.TempDir()
	filePath := filepath.Join(tempDir, "tone.wav")
	if err := os.WriteFile(filePath, buildSimpleWavPCM(), 0o644); err != nil {
		t.Fatalf("failed to create temp audio file: %v", err)
	}

	data, err := ExtractMelodyFromAudioFile(filePath)
	if err != nil {
		t.Fatalf("expected melody extraction to succeed: %v", err)
	}
	if len(data.Notes) < 2 {
		t.Fatalf("expected a note stream with multiple notes, got %d", len(data.Notes))
	}
	if got := data.Notes[0].Pitch; got != "A" {
		t.Fatalf("expected extracted pitch A, got %s", got)
	}
	if data.Notes[0].Duration <= 0 {
		t.Fatal("expected a positive note duration")
	}
	if data.Notes[1].Duration <= 0 {
		t.Fatal("expected a positive duration for the second note")
	}
}

func TestExtractMelodyFromAudioFileRejectsInvalidPath(t *testing.T) {
	_, err := ExtractMelodyFromAudioFile("missing.wav")
	if err == nil {
		t.Fatal("expected invalid path to fail")
	}
}

func buildSimpleWavPCM() []byte {
	const sampleRate = 22050
	const frequency = 440
	const durationSeconds = 0.25
	const amplitude = 12000
	sampleCount := int(math.Round(float64(sampleRate) * durationSeconds))
	pcm := make([]int16, sampleCount)
	for i := range pcm {
		angle := 2 * math.Pi * float64(frequency) * float64(i) / float64(sampleRate)
		pcm[i] = int16(math.Sin(angle) * amplitude)
	}

	var buf bytes.Buffer
	writeUint32 := func(v uint32) {
		_ = binary.Write(&buf, binary.LittleEndian, v)
	}
	writeUint16 := func(v uint16) {
		_ = binary.Write(&buf, binary.LittleEndian, v)
	}

	buf.WriteString("RIFF")
	writeUint32(uint32(36 + len(pcm)*2))
	buf.WriteString("WAVE")
	buf.WriteString("fmt ")
	writeUint32(16)
	writeUint16(1)
	writeUint16(1)
	writeUint32(sampleRate)
	writeUint32(sampleRate * 2)
	writeUint16(2)
	writeUint16(16)
	buf.WriteString("data")
	writeUint32(uint32(len(pcm) * 2))
	for _, sample := range pcm {
		_ = binary.Write(&buf, binary.LittleEndian, sample)
	}

	return buf.Bytes()
}
