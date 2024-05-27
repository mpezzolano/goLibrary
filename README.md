# goLibrary


This is a Go utility library that provides logging services and middleware support.

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

