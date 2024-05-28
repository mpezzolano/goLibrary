package fileutils

import (
	"io/ioutil"
	"os"
)

func ReadFile(filepath string) (string, error) {
	content, err := ioutil.ReadFile(filepath)
	if err != nil {
		return "", err
	}
	return string(content), nil
}

func WriteFile(filepath string, content string) error {
	return os.WriteFile(filepath, []byte(content), 0644)
}

func DeleteFile(filepath string) error {
	return os.Remove(filepath)
}

func UpdateFile(filepath string, newContent string) error {
	return os.WriteFile(filepath, []byte(newContent), 0644)
}
