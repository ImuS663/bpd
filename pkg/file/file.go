package file

import (
	"os"
	"path/filepath"
)

func GetFilePathAndName(url string, outDir string) (string, string) {
	fileName := filepath.Base(url)

	return fileName, filepath.Join(outDir, fileName)
}

// OpenFile takes a file path and attempts to open the file for writing.
// If the file does not exist, it will be created.
// The function returns a pointer to the opened file and an error if any.
func OpenFile(filePath string) (*os.File, error) {
	return os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY, 0644)
}

// Exists takes a file path and returns true if the file exists, false otherwise.
func Exists(filePath string) bool {
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return false
	}
	return true
}
