package utils

import (
	"bufio"
	"os"
	"strings"
)

// LoaderEnvFile reads an .env file line by line,
// extracts the key-value pairs, and sets them as
// environment variables in the current process.
// It handles comments, skips lines without key-value pairs,
// trims whitespace, and reports any errors that
// occur during the process.
func LoaderEnvFile(filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "#") || !strings.Contains(line, "=") {
			continue // Skip comments and lines without key-value pairs
		}

		parts := strings.SplitN(line, "=", 2)
		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])

		err := os.Setenv(key, value)
		if err != nil {
			return err
		}
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}
