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
	return &ConversionInput{Melody: melody}, nil
}
