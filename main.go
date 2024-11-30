package main

import (
	"fmt"

	"github.com/iabdulzahid/go-logger/logger"
	"go.uber.org/zap"
)

func main() {
	config := logger.Config{
		AppName:            "go-logger",
		LogLevel:           "info",
		LogFormat:          "json", // "json" or "plain"
		EnableCallerInfo:   true,
		JSONFormat:         true,               // Use JSON format
		LogOutput:          []string{"stdout"}, //[]string{"stdout", "file"}
		LogFilePath:        "./logs/app.log",
		LogFilePermissions: "0644",
		TimeFormat:         "2006-01-02 15:04:05", // Custom timestamp format
		EnableRotation:     true,
		MaxSize:            50,
		MaxBackups:         10,
		MaxAge:             28,
		Compress:           true,
	}

	log, err := logger.NewLogger(config)
	if err != nil {
		fmt.Println("Error creating logger:", err)
		return
	}
	ctxLogger := log.WithContext(
		zap.String("request_id", "12345"),
		zap.String("user_id", "admin"),
	)

	// Log a message with the context
	ctxLogger.Info("This is a context-specific info message")

	// Log some messages
	ctxLogger.Info("This is an info message")
	ctxLogger.Debug("This is a debug message", "session_id", "xyz789")
	log.Warn("This is a warning message", "module", "auth")
	log.Error("This is an error message", fmt.Errorf("sample error"), "function", "main")
}
