package logger

// Config struct holds configuration parameters for the logger
type Config struct {
	AppName            string
	LogLevel           string
	LogFormat          string
	EnableRotation     bool
	LogFilePath        string
	MaxSize            int      // Maximum size of a log file (in MB) before rotation
	MaxBackups         int      // Maximum number of old backups to keep
	MaxAge             int      // Maximum age (in days) of a log file before it is deleted
	Compress           bool     // Whether to compress old log files
	TimeFormat         string   // Format for timestamp (default is "RFC3339")
	EnableCallerInfo   bool     // Whether to include caller information in logs
	JSONFormat         bool     // Whether to use JSON format for logs (default is false)
	LogOutput          []string // Output destinations for logs (e.g., "stdout", "file", "syslog")
	LogFilePermissions string   // File permissions for log file (e.g., "0644")
}
