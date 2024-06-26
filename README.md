# goLibrary

This is a Go utility library designed to provide a comprehensive wrapper for various functionalities, making it easy to integrate logging services, middleware support, file management, and more into your projects.

## Logger

### Usage Example

```go
package main

import (
    "your_project_name/logger"
)

func main() {
    log := logger.NewLogger(logger.INFO)

    log.Info("This is an info message")
    log.Warning("This is a warning message")
    log.Error("This is an error message")
    log.Debug("This is a debug message") // Will not be printed if the logger level is INFO
}
´´´
## File Management

```go
    // Writing to a file
    err := fileutils.WriteFile("example.txt", "This is an example content")
    if err != nil {
        log.Error("Failed to write file: " + err.Error())
    } else {
        log.Info("File written successfully")
    }
´´´
