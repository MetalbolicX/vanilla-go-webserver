package utils

import (
	"log"
	"os"
)

// GetRootDir returns the root path of project.
func GetRootDir() string {
	executablePath, err := os.Getwd()
	if err != nil {
		log.Fatal("Failed to get executable")
	}
	return executablePath
}
