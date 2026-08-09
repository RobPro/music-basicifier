package platform

import (
	"fmt"
	"os/exec"
	"strings"
)

// CopyTextToClipboard copies a string onto the Windows clipboard using PowerShell.
func CopyTextToClipboard(text string) error {
	trimmed := strings.TrimSpace(text)
	if trimmed == "" {
		return fmt.Errorf("text is required")
	}

	cmd := exec.Command("powershell", "-NoProfile", "-Command", "$text = [Console]::In.ReadToEnd(); Set-Clipboard -Value $text")
	cmd.Stdin = strings.NewReader(trimmed)
	if output, err := cmd.CombinedOutput(); err != nil {
		return fmt.Errorf("copy to clipboard failed: %w: %s", err, strings.TrimSpace(string(output)))
	}
	return nil
}
