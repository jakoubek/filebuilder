package filebuilder

import (
	"os"
	"path/filepath"
)

func OptionalPathPrepend(filename string, optionalDir string) string {
	dir, file := filepath.Split(filename)
	if dir == "" {
		dir = optionalDir
	}
	return filepath.Join(dir, file)
}

func readFile(path string) (string, error) {
	fileContent, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(fileContent), nil
}

func doesFileExist(path string) bool {
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
