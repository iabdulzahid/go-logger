# go-logger

## **Overview**
The Go Logger package provides a customizable logging solution with support for JSON and plain-text formats, log rotation, and dynamic log level configuration. This package is built using Uber's Zap library for high performance and flexibility.

---

## **Features**
- Supports different log levels: `debug`, `info`, `warn`, `error`, `fatal`.
- Outputs logs to multiple destinations (`stdout`, `file`, etc.).
- Supports JSON and plain-text formats.
- Includes caller information (file, function, and line number).
- Customizable timestamp format.
- File rotation with options for size, backup count, and compression.

---

## **Installation**
Add the logger package to your project:

```bash
go get github.com/iabdulzahid/go-logger
```

## Usage
### Basic Example

```
package main

import (
	"fmt"
	"github.com/iabdulzahid/go-logger/logger"
	"go.uber.org/zap"
)

func main() {
	// Configure the logger
	config := logger.Config{
		AppName:            "my-app",
		LogLevel:           "info",
		EnableRotation:     true,
		LogFilePath:        "./logs/app.log",
		MaxSize:            50,
		MaxBackups:         5,
		MaxAge:             30,
		Compress:           true,
		TimeFormat:         "2006-01-02 15:04:05",
		EnableCallerInfo:   true,
		JSONFormat:         true,
		LogOutput:          []string{"stdout", "file"},
		LogFilePermissions: "0644",
	}

	// Initialize the logger
	log, err := logger.NewLogger(config)
	if err != nil {
		fmt.Println("Error creating logger:", err)
		return
	}

	// Log messages
	log.Info("This is an informational message", "request_id", "12345", "user_id", "admin")
	log.Warn("This is a warning message", "module", "auth")
	log.Error("This is an error message", fmt.Errorf("sample error"), "function", "main")

	// Context-specific logger
	ctxLogger := log.WithContext(
		zap.String("request_id", "12345"),
		zap.String("user_id", "admin"),
	)
	ctxLogger.Info("This is a context-specific info message")
}
```

## **Configuration**

The logger is highly customizable through the `Config` struct. Below is an overview of the configuration options available.

### **Config Struct**

The `Config` struct holds configuration parameters for the logger.

```
go
type Config struct {
    AppName            string   // Application name (required)
    LogLevel           string   // Log level (debug, info, warn, error, fatal) (required)
    EnableRotation     bool     // Enable log file rotation (default: false)
    LogFilePath        string   // Path to the log file (required if log rotation is enabled)
    MaxSize            int      // Max size (in MB) of the log file before rotation (default: 10 MB)
    MaxBackups         int      // Max number of old backups to keep (default: 3)
    MaxAge             int      // Max age (in days) of the log file before deletion (default: 7)
    Compress           bool     // Whether to compress old log files (default: false)
    TimeFormat         string   // Format for the timestamp (default: "RFC3339")
    EnableCallerInfo   bool     // Whether to include caller information in logs (default: false)
    JSONFormat         bool     // Whether to use JSON format for logs (default: false)
    LogOutput          []string // Output destinations for logs (stdout, file, syslog) (required)
    LogFilePermissions string   // File permissions for the log file (e.g., "0644") (default: "0644")
    LogFormat          string   // Format of logs, either "json" or "plain" (default: "plain")
}
```
## **Field Descriptions**

| **Field Name**       | **Type**      | **Description**                                                                                         | **Default**       | **Example**                    |
|----------------------|---------------|---------------------------------------------------------------------------------------------------------|-------------------|--------------------------------|
| `AppName`            | `string`      | The name of the application to include in each log entry.                                                 | Required          | `"my-app"`                     |
| `LogLevel`           | `string`      | The log level: `debug`, `info`, `warn`, `error`, or `fatal`.                                              | Required          | `"info"`                       |
| `EnableRotation`     | `bool`        | Enables log file rotation.                                                                                | `false`           | `true`                         |
| `LogFilePath`        | `string`      | The path to the log file. Required if `EnableRotation` is `true`.                                         | None              | `"./logs/app.log"`             |
| `MaxSize`            | `int`         | The maximum size (in MB) of the log file before rotation.                                                 | `10`              | `50`                           |
| `MaxBackups`         | `int`         | The maximum number of old backup log files to retain.                                                     | `3`               | `5`                            |
| `MaxAge`             | `int`         | The maximum age (in days) of log files before they are deleted.                                          | `7`               | `30`                           |
| `Compress`           | `bool`        | Whether to compress old log files after rotation.                                                        | `false`           | `true`                         |
| `TimeFormat`         | `string`      | The format for the timestamp in logs. Defaults to `RFC3339` if not provided.                             | `RFC3339`         | `"2006-01-02 15:04:05"`        |
| `EnableCallerInfo`   | `bool`        | Whether to include caller information (file, function, line number) in log entries.                      | `false`           | `true`                         |
| `JSONFormat`         | `bool`        | Whether to log in JSON format. If `false`, logs will be plain-text.                                        | `false`           | `true`                         |
| `LogOutput`          | `[]string`    | Specifies output destinations for logs. Options include `"stdout"`, `"file"`, and `"syslog"`.           | `["stdout"]`      | `["stdout", "file"]`           |
| `LogFilePermissions` | `string`      | The file permissions to apply when creating the log file (if `LogOutput` includes `"file"`).               | `"0644"`          | `"0600"`                       |
| `LogFormat`          | `string`      | The log format type, which can be `"json"` for structured logs or `"plain"` for plain text logs.         | `"plain"`         | `"json"`                       |


## Example Configuration
Here is an example of a complete configuration:

```
config := logger.Config{
    AppName:            "my-app",                  // Application name
    LogLevel:           "info",                    // Log level (debug, info, warn, error, fatal)
    EnableRotation:     true,                      // Enable log file rotation
    LogFilePath:        "./logs/app.log",          // Log file path
    MaxSize:            50,                        // Max log file size (MB)
    MaxBackups:         5,                         // Max backups to retain
    MaxAge:             30,                        // Max age of log files (days)
    Compress:           true,                      // Compress old log files
    TimeFormat:         "2006-01-02 15:04:05",    // Custom timestamp format
    EnableCallerInfo:   true,                      // Include caller information
    JSONFormat:         true,                      // Use JSON format for logs
    LogOutput:          []string{"stdout", "file"}, // Log to stdout and file
    LogFilePermissions: "0644",                    // File permissions for log file
    LogFormat:          "json",                    // Log format
}

```

## Log Methods
The logger provides several methods to log messages at different levels. Each method supports dynamic fields to add additional context to log entries.


## **Method Descriptions**

| **Method** | **Description**                                                | **Example Usage**                                             |
|------------|----------------------------------------------------------------|---------------------------------------------------------------|
| `Info`     | Logs an informational message.                                 | `log.Info("Application started", "version", "1.0.0")`         |
| `Debug`    | Logs a debug message.                                          | `log.Debug("Debugging user data", "user_id", "12345")`        |
| `Warn`     | Logs a warning message.                                        | `log.Warn("Disk space low", "disk", "/dev/sda")`               |
| `Error`    | Logs an error message, including error details.                | `log.Error("Error occurred", fmt.Errorf("timeout"), "task", "sync")` |


## Context-Specific Logging
You can create a logger with additional context-specific fields (e.g., request_id, user_id) by using the WithContext method. These fields will be included in every log entry.

```
go

ctxLogger := log.WithContext(
    zap.String("request_id", "12345"),
    zap.String("user_id", "admin"),
)
ctxLogger.Info("This is a context-specific log message")
```

## Closing the Logger
The logger should be closed when your application terminates to ensure all buffered log entries are written.

```
go

log.Close()
```

## File Rotation
To enable log file rotation, set EnableRotation to true in the configuration. You can also specify the maximum size of the log file, the number of backups to keep, and the maximum age of the log files.

## Rotation Configuration
## **Field Descriptions**

| **Field**    | **Description**                                                  |
|--------------|------------------------------------------------------------------|
| `MaxSize`    | The maximum size (in MB) of a log file before rotation.         |
| `MaxBackups` | The maximum number of backups to retain.                        |
| `MaxAge`     | The maximum age (in days) of log files before they are deleted. |
| `Compress`   | Whether to compress old log files.                              |


## Custom Timestamp Format
You can customize the timestamp format in logs by setting the TimeFormat field in the configuration. The default is RFC3339.

```
go
config := logger.Config{
    TimeFormat: "2006-01-02 15:04:05",  // Custom timestamp format
}
```

## Log Output Destinations
You can specify multiple output destinations for logs, such as stdout (console), file, and syslog. For example:

```
go

config := logger.Config{
    LogOutput: []string{"stdout", "file"},
}
```
