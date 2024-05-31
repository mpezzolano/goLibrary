package main

import (
	"goLibrary/fileutils"
	_ "goLibrary/fileutils"
	"goLibrary/logger"
	_ "goLibrary/logger"
)

func runFileUtilsExamples() {
	log := logger.NewLogger(logger.INFO)

	err := fileutils.WriteFile("example.txt", "This is an example content")
	if err != nil {
		log.Error("Failed to write file: " + err.Error())
	} else {
		log.Info("File written successfully")
	}

	content, err := fileutils.ReadFile("example.txt")
	if err != nil {
		log.Error("Failed to read file: " + err.Error())
	} else {
		log.Info("File content: " + content)
	}

	err = fileutils.UpdateFile("example.txt", "This is updated content")
	if err != nil {
		log.Error("Failed to update file: " + err.Error())
	} else {
		log.Info("File updated successfully")
	}

	err = fileutils.DeleteFile("example.txt")
	if err != nil {
		log.Error("Failed to delete file: " + err.Error())
	} else {
		log.Info("File deleted successfully")
	}
}
