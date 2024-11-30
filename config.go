package main

// Config holds the configuration options for the logger.
//
// This struct is used to configure various aspects of the logger, such as log
// level, output destinations, rotation settings, and the format of the logs.
type Config struct {
	// AppName is the name of the application to include in each log entry.
	// Default: Required field.
	AppName string

	// LogLevel sets the log level. Options are: "debug", "info", "warn", "error", "fatal".
	// Default: "info"
	LogLevel string

	// EnableRotation enables log file rotation when the log file reaches the specified size.
	// Default: false
	EnableRotation bool

	// LogFilePath specifies the path where the log file is stored. Required if EnableRotation is true.
	// Example: "./logs/app.log"
	LogFilePath string

	// MaxSize is the maximum size (in MB) of a log file before rotation occurs.
	// Default: 10 MB
	MaxSize int

	// MaxBackups specifies the maximum number of backup log files to retain.
	// Default: 3 backups
	MaxBackups int

	// MaxAge is the maximum age (in days) of log files before they are deleted.
	// Default: 7 days
	MaxAge int

	// Compress enables compression for old log files after rotation.
	// Default: false
	Compress bool

	// TimeFormat specifies the format for the timestamp in log entries.
	// Default: RFC3339
	TimeFormat string

	// EnableCallerInfo includes caller information in the logs, such as file, function, and line number.
	// Default: false
	EnableCallerInfo bool

	// JSONFormat determines whether the log output should be in JSON format.
	// Default: false (plain-text logs)
	JSONFormat bool

	// LogOutput specifies the log output destinations. Possible values include "stdout", "file", and "syslog".
	// Default: ["stdout"]
	LogOutput []string

	// LogFilePermissions sets the file permissions for the log file (if LogOutput includes "file").
	// Example: "0644"
	LogFilePermissions string

	// LogFormat specifies the log format type: "json" for structured logs or "plain" for plain-text logs.
	// Default: "plain"
	LogFormat string
}

// Config struct holds configuration parameters for the logger
// type Config struct {
// 	AppName            string
// 	LogLevel           string
// 	EnableRotation     bool
// 	LogFilePath        string
// 	MaxSize            int      // Maximum size of a log file (in MB) before rotation
// 	MaxBackups         int      // Maximum number of old backups to keep
// 	MaxAge             int      // Maximum age (in days) of a log file before it is deleted
// 	Compress           bool     // Whether to compress old log files
// 	TimeFormat         string   // Format for timestamp (default is "RFC3339")
// 	EnableCallerInfo   bool     // Whether to include caller information in logs
// 	JSONFormat         bool     // Whether to use JSON format for logs (default is false)
// 	LogOutput          []string // Output destinations for logs (e.g., "stdout", "file", "syslog")
// 	LogFilePermissions string   // File permissions for log file (e.g., "0644")
//     LogFormat          string
// }
