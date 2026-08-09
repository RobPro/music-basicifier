package platform

import "fmt"

// OutputBundle contains the generated text outputs for the supported targets.
type OutputBundle struct {
	QBASIC   string
	Adafruit string
}

// GenerateOutputBundle creates both text outputs from a conversion input.
func GenerateOutputBundle(input *ConversionInput) (*OutputBundle, error) {
	if input == nil || input.Melody == nil {
		return nil, fmt.Errorf("conversion input is required")
	}

	qbasic, err := GenerateQBASICProgram(input)
	if err != nil {
		return nil, fmt.Errorf("generate QBASIC output: %w", err)
	}
	adafruit, err := GenerateAdafruitJavaScript(input)
	if err != nil {
		return nil, fmt.Errorf("generate Adafruit output: %w", err)
	}

	return &OutputBundle{QBASIC: qbasic, Adafruit: adafruit}, nil
}
