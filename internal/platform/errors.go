package platform

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

// LogError writes a simple error message to a temporary log file.
func LogError(message string) error {
	if message == "" {
		return fmt.Errorf("message is required")
	}

	dir := os.TempDir()
	logPath := filepath.Join(dir, "music-basicifier-errors.log")
	line := fmt.Sprintf("%s %s\n", time.Now().UTC().Format(time.RFC3339), message)
	f, err := os.OpenFile(logPath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0o644)
	if err != nil {
		return fmt.Errorf("write error log: %w", err)
	}
	defer f.Close()

	_, err = f.WriteString(line)
	if err != nil {
		return fmt.Errorf("write error log entry: %w", err)
	}
	return nil
}
