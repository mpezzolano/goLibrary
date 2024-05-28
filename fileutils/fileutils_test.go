package fileutils

import (
	_ "io/ioutil"
	"os"
	"testing"
)

func TestWriteReadUpdateDeleteFile(t *testing.T) {
	filename := "test.txt"
	content := "This is a test content"

	// Test WriteFile
	err := WriteFile(filename, content)
	if err != nil {
		t.Fatalf("Failed to write file: %v", err)
	}

	// Test ReadFile
	readContent, err := ReadFile(filename)
	if err != nil {
		t.Fatalf("Failed to read file: %v", err)
	}
	if readContent != content {
		t.Errorf("Expected content %v, but got %v", content, readContent)
	}

	// Test UpdateFile
	updatedContent := "This is updated content"
	err = UpdateFile(filename, updatedContent)
	if err != nil {
		t.Fatalf("Failed to update file: %v", err)
	}

	readContent, err = ReadFile(filename)
	if err != nil {
		t.Fatalf("Failed to read file: %v", err)
	}
	if readContent != updatedContent {
		t.Errorf("Expected content %v, but got %v", updatedContent, readContent)
	}

	// Test DeleteFile
	err = DeleteFile(filename)
	if err != nil {
		t.Fatalf("Failed to delete file: %v", err)
	}
	if _, err := os.Stat(filename); !os.IsNotExist(err) {
		t.Errorf("Expected file to be deleted, but it still exists")
	}
}
