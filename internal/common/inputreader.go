package common

import (
	"os"
	"strings"
)

// ReadInput reads a file and returns its content as string
func ReadInput(path string) (string, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(content)), nil
}
