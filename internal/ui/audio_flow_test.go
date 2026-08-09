package ui

import (
	"bytes"
	"encoding/binary"
	"math"
	"os"
	"path/filepath"
	"testing"
)

func TestBuildConversionInputFromAudioFile(t *testing.T) {
	tempDir := t.TempDir()
	filePath := filepath.Join(tempDir, "tone.wav")
	if err := os.WriteFile(filePath, buildSimpleWavPCM(), 0o644); err != nil {
		t.Fatalf("failed to create temp audio file: %v", err)
	}

	input, err := buildConversionInputFromAudioFile(filePath)
	if err != nil {
		t.Fatalf("expected conversion input to be built: %v", err)
	}
	if input == nil || input.Melody == nil {
		t.Fatal("expected conversion input with melody data")
	}
	if len(input.Melody.Notes) == 0 {
		t.Fatal("expected at least one melody note")
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
