package file

import (
	"os"
	"path/filepath"
)

func GetFilePathAndName(url string, outDir string) (string, string) {
	fileName := filepath.Base(url)

	return fileName, filepath.Join(outDir, fileName)
}

func OpenFile(filePath string) (*os.File, error) {
	return os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY, 0644)
}

func Exists(filePath string) bool {
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return false
	}
	return true
}
