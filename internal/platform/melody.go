package platform

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"math"
	"os"
	"path/filepath"
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
// It currently supports PCM WAV decoding and returns clear errors for unsupported formats.
func ExtractMelodyFromAudioFile(path string) (*MelodyData, error) {
	if err := ValidateAudioFilePath(path); err != nil {
		return nil, fmt.Errorf("invalid audio file: %w", err)
	}

	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("read audio file: %w", err)
	}

	ext := filepath.Ext(path)
	if ext != ".wav" && ext != ".WAV" {
		return nil, fmt.Errorf("decode audio file: unsupported format for melody extraction (%s)", ext)
	}

	notes, err := extractNoteStreamFromWAV(data)
	if err != nil {
		return nil, fmt.Errorf("decode audio file: %w", err)
	}

	return &MelodyData{
		SourcePath: path,
		Notes:      notes,
	}, nil
}

func extractNoteStreamFromWAV(data []byte) ([]MelodyNote, error) {
	samples, sampleRate, err := extractPCMFrames(data)
	if err != nil {
		return nil, err
	}
	if len(samples) < 2 {
		return nil, fmt.Errorf("not enough audio samples")
	}

	windowSize := sampleRate / 8
	if windowSize < 2 {
		windowSize = 2
	}

	var notes []MelodyNote
	for start := 0; start < len(samples); start += windowSize {
		end := start + windowSize
		if end > len(samples) {
			end = len(samples)
		}
		if end-start < 2 {
			continue
		}

		frequency := estimateFrequency(samples[start:end], sampleRate)
		if frequency <= 0 {
			continue
		}

		duration := int(math.Ceil(float64(end-start) / float64(sampleRate) * 4))
		if duration < 1 {
			duration = 1
		}
		notes = append(notes, MelodyNote{Pitch: noteNameForFrequency(frequency), Duration: duration})
	}
	if len(notes) == 0 {
		return nil, fmt.Errorf("unable to estimate note stream")
	}
	return notes, nil
}

func extractPCMFrames(data []byte) ([]float64, int, error) {
	if len(data) < 44 {
		return nil, 0, fmt.Errorf("wav data is too small")
	}
	if !bytes.Equal(data[:4], []byte("RIFF")) || !bytes.Equal(data[8:12], []byte("WAVE")) {
		return nil, 0, fmt.Errorf("not a wav file")
	}

	audioFormat := int(binary.LittleEndian.Uint16(data[20:22]))
	if audioFormat != 1 {
		return nil, 0, fmt.Errorf("unsupported wav format")
	}

	numChannels := int(binary.LittleEndian.Uint16(data[22:24]))
	sampleRate := int(binary.LittleEndian.Uint32(data[24:28]))
	bitsPerSample := int(binary.LittleEndian.Uint16(data[34:36]))
	if numChannels <= 0 || sampleRate <= 0 || bitsPerSample <= 0 {
		return nil, 0, fmt.Errorf("invalid wav header")
	}
	if bitsPerSample != 8 && bitsPerSample != 16 {
		return nil, 0, fmt.Errorf("unsupported bit depth")
	}

	dataOffset := 12
	dataSize := 0
	for dataOffset+8 <= len(data) {
		chunkID := string(data[dataOffset : dataOffset+4])
		chunkSize := int(binary.LittleEndian.Uint32(data[dataOffset+4 : dataOffset+8]))
		if chunkID == "data" {
			dataOffset += 8
			dataSize = chunkSize
			break
		}
		dataOffset += 8 + chunkSize
		if dataOffset%2 != 0 {
			dataOffset++
		}
	}
	if dataSize == 0 || dataOffset+dataSize > len(data) {
		return nil, 0, fmt.Errorf("wav data chunk not found")
	}

	frameSize := (bitsPerSample / 8) * numChannels
	if frameSize <= 0 {
		return nil, 0, fmt.Errorf("invalid frame size")
	}

	samples := make([]float64, 0, dataSize/frameSize)
	for frameStart := dataOffset; frameStart+frameSize <= dataOffset+dataSize; frameStart += frameSize {
		sum := 0.0
		for channel := 0; channel < numChannels; channel++ {
			sampleStart := frameStart + channel*(bitsPerSample/8)
			if bitsPerSample == 8 {
				sum += float64(int(data[sampleStart]) - 128)
			} else {
				sum += float64(int16(binary.LittleEndian.Uint16(data[sampleStart : sampleStart+2])))
			}
		}
		samples = append(samples, sum/float64(numChannels))
	}
	if len(samples) < 2 {
		return nil, 0, fmt.Errorf("not enough audio samples")
	}
	return samples, sampleRate, nil
}

func estimateFrequency(samples []float64, sampleRate int) float64 {
	if len(samples) < 2 || sampleRate <= 0 {
		return 0
	}

	zeroCrossings := 0
	prev := samples[0]
	for _, sample := range samples[1:] {
		if (prev <= 0 && sample > 0) || (prev >= 0 && sample < 0) {
			zeroCrossings++
		}
		prev = sample
	}

	durationSeconds := float64(len(samples)) / float64(sampleRate)
	if durationSeconds <= 0 {
		return 0
	}
	return float64(zeroCrossings) / (2 * durationSeconds)
}

func noteNameForFrequency(frequency float64) string {
	const baseFrequency = 440.0
	const baseMidi = 69
	noteMidi := int(math.Round(12*math.Log2(frequency/baseFrequency) + baseMidi))
	noteNames := []string{"C", "C#", "D", "D#", "E", "F", "F#", "G", "G#", "A", "A#", "B"}
	noteIndex := (noteMidi%12 + 12) % 12
	return noteNames[noteIndex]
}
