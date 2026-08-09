package platform

import "fmt"

// ConversionInput carries the melody data into the output conversion stage.
type ConversionInput struct {
	Melody *MelodyData
}

// BuildConversionInput creates a conversion input for downstream output generation.
func BuildConversionInput(melody *MelodyData) (*ConversionInput, error) {
	if melody == nil {
		return nil, fmt.Errorf("melody data is required")
	}
	if len(melody.Notes) == 0 {
		return nil, fmt.Errorf("melody data must contain at least one note")
	}
	for i, note := range melody.Notes {
		if note.Pitch == "" {
			return nil, fmt.Errorf("melody note %d has empty pitch", i)
		}
		if note.Duration <= 0 {
			return nil, fmt.Errorf("melody note %d has invalid duration", i)
		}
	}
	return &ConversionInput{Melody: melody}, nil
}
