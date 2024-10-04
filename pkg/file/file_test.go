package file

import (
	"os"
	"testing"
)

func TestGetFilePathAndName(t *testing.T) {
	fileName, filePath := GetFilePathAndName("https://example.com/test.zip", "/tmp")

	if fileName != "test.zip" {
		t.Errorf("expected 'test.zip', got %s", fileName)
	}

	if filePath != "/tmp/test.zip" {
		t.Errorf("expected '/tmp/test.zip', got %s", filePath)
	}
}

func TestOpenFileCreateFile(t *testing.T) {
	tmpFile, err := OpenFile("test.txt")
	if err != nil {
		t.Error(err)
	}
	defer os.Remove(tmpFile.Name())

	if _, err := os.Stat(tmpFile.Name()); os.IsNotExist(err) {
		t.Error("file not created")
	}
}

func TestOpenFileOpenExistFile(t *testing.T) {
	tmpFile, err := os.CreateTemp("", "test.txt")
	if err != nil {
		t.Error(err)
	}
	defer os.Remove(tmpFile.Name())

	if _, err := OpenFile(tmpFile.Name()); err != nil {
		t.Error(err)
	}
}

func TestExistsTrue(t *testing.T) {
	tmpFile, err := os.CreateTemp("", "test")
	if err != nil {
		t.Error(err)
	}
	defer os.Remove(tmpFile.Name())

	if !Exists(tmpFile.Name()) {
		t.Error("expected true, got false")
	}
}
