// package main

// import (
// 	"fmt"

// 	"github.com/iabdulzahid/go-logger/logger"
// )

// func main() {
// 	// Logger configuration
// 	config := logger.Config{
// 		FileName:         "app.log",
// 		AppName:          "MyApp",
// 		LogLevel:         "debug",
// 		JSONFormat:       false, // Set to true for JSON format logs
// 		MaxSize:          10,    // Rotate logs after 10MB
// 		MaxBackups:       3,     // Keep 3 old log backups
// 		MaxAge:           7,     // Delete logs older than 7 days
// 		Compress:         true,  // Compress old log files
// 		LogRotation:      true,  // Enable log rotation
// 		EnableCallerInfo: true,  // Enable caller info in logs
// 	}

// 	// Create a logger instance
// 	log, err := logger.NewLogger(config)
// 	if err != nil {
// 		fmt.Println("Error creating logger:", err)
// 		return
// 	}
// 	defer log.Close()

// 	// Example log entries with different levels
// 	log.Info("Application started", "version", "1.0.0", "env", "production")
// 	log.Warn("Disk space low", "disk", "/dev/sda1", "free_space", 10)

// 	// Dynamically change log level
// 	err = log.SetLogLevel("debug")
// 	if err != nil {
// 		fmt.Println("Error setting log level:", err)
// 	}

// 	log.Debug("Debugging message", "user", "john_doe", "operation", "start")
// 	// log.Error("Disk error", "disk", "/dev/sda1", "free_space", -1)

// 	// Example Fatal log (this will terminate the application)
// 	// log.Fatal("Application crashed", "error", "disk failure")
// }

package main

import (
	"fmt"

	"github.com/iabdulzahid/go-logger/logger"
)

func main() {
	config := logger.Config{
		AppName:            "MyApp",
		LogLevel:           "debug",
		LogFormat:          "json", // "json" or "plain"
		EnableCallerInfo:   true,
		JSONFormat:         true, // Use JSON format
		LogOutput:          []string{"stdout", "file"},
		LogFilePath:        "./logs/app.log",
		LogFilePermissions: "0644",
		TimeFormat:         "2006-01-02 15:04:05", // Custom timestamp format
	}

	log, err := logger.NewLogger(config)
	if err != nil {
		fmt.Println("Error creating logger:", err)
		return
	}

	// Log some messages
	log.LogInfo("This is an info message", "request_id", "12345", "user_id", "abcde")
	log.LogDebug("This is a debug message", "session_id", "xyz789")
	log.LogWarn("This is a warning message", "module", "auth")
	log.LogError("This is an error message", fmt.Errorf("sample error"), "function", "main")
}
