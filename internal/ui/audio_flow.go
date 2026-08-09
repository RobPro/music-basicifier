// Package ui contains Fyne application setup and top-level window wiring.
package ui

import "github.com/RobPro/music-basicifier/internal/platform"

func buildConversionInputFromAudioFile(path string) (*platform.ConversionInput, error) {
	melody, err := platform.ExtractMelodyFromAudioFile(path)
	if err != nil {
		return nil, err
	}
	return platform.BuildConversionInput(melody)
}
