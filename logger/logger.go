// package logger

// package logger

// import (
// 	"fmt"
// 	"runtime"
// 	"strings"

// 	"github.com/natefinch/lumberjack"
// 	"go.uber.org/zap"
// 	"go.uber.org/zap/zapcore"
// )

// // Config struct holds configuration parameters for the logger
// type Logger struct {
// 	logger *zap.Logger
// }

// // CustomEncoder configures the encoder to use RFC3339 timestamp format and colored log levels
// func customEncoderConfig(isConsole bool, config Config) zapcore.EncoderConfig {
// 	encoderConfig := zap.NewProductionEncoderConfig()
// 	encoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout(config.TimeFormat) // time.RFC3339 timestamp format
// 	if isConsole {
// 		encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder // Colored log level for console output
// 	} else {
// 		encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder // Plain log level for file output
// 	}
// 	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
// 	return encoderConfig
// }

// // NewLogger creates a new logger instance with the provided configuration
// func NewLogger(config Config) (*Logger, error) {
// 	var level zapcore.Level
// 	switch strings.ToLower(config.LogLevel) {
// 	case "debug":
// 		level = zapcore.DebugLevel
// 	case "info":
// 		level = zapcore.InfoLevel
// 	case "warn":
// 		level = zapcore.WarnLevel
// 	case "error":
// 		level = zapcore.ErrorLevel
// 	case "fatal":
// 		level = zapcore.FatalLevel
// 	default:
// 		return nil, fmt.Errorf("invalid log level: %s", config.LogLevel)
// 	}

// 	// Set log format
// 	var consoleEncoder zapcore.Encoder
// 	var fileEncoder zapcore.Encoder
// 	consoleEncoder = zapcore.NewConsoleEncoder(customEncoderConfig(true, config)) // Console with colors
// 	fileEncoder = zapcore.NewJSONEncoder(customEncoderConfig(true, config))       // File without colors

// 	// Log rotation configuration
// 	var writer zapcore.WriteSyncer
// 	if config.EnableRotation {
// 		writer = zapcore.AddSync(&lumberjack.Logger{
// 			Filename:   config.LogFilePath,
// 			MaxSize:    config.MaxSize,    // 50MB
// 			MaxBackups: config.MaxBackups, //5
// 			MaxAge:     config.MaxAge,     // 7days
// 			Compress:   config.Compress,   // default true
// 		})
// 	} else {
// 		writer = zapcore.AddSync(&lumberjack.Logger{
// 			Filename: config.LogFilePath,
// 			MaxSize:  config.MaxSize, // 50MB
// 		})
// 	}

// 	// Creating two cores: one for console and one for file
// 	consoleCore := zapcore.NewCore(consoleEncoder, zapcore.Lock(zapcore.AddSync(&lumberjack.Logger{})), level)
// 	fileCore := zapcore.NewCore(fileEncoder, writer, level)

// 	// Combine the two cores into one logger
// 	core := zapcore.NewTee(consoleCore, fileCore)

// 	// Create logger instance
// 	logger := zap.New(core)

// 	return &Logger{logger: logger}, nil
// }

// // LogInfo logs an informational message with timestamp and optional context
// // Now accepts key-value pairs of any type
// func (l *Logger) LogInfo(msg string, args ...interface{}) {
// 	// Convert the key-value pairs (args) into zap.Field
// 	var fields []zap.Field
// 	for i := 0; i < len(args); i += 2 {
// 		// Ensure that args are in key-value pairs
// 		if i+1 < len(args) {
// 			key, ok := args[i].(string)
// 			if !ok {
// 				l.logger.Error("Invalid key type, key must be a string", zap.Any("key", args[i]))
// 				continue
// 			}
// 			value := args[i+1]
// 			// Add the key-value pair as zap.Field
// 			fields = append(fields, zap.Any(key, value))
// 		}
// 	}
// 	// Log the message with the fields
// 	l.logger.Info(msg, fields...)
// }

// // LogDebug logs a debug message with timestamp and optional context
// func (l *Logger) LogDebug(msg string, args ...interface{}) {
// 	var fields []zap.Field
// 	for i := 0; i < len(args); i += 2 {
// 		if i+1 < len(args) {
// 			key, ok := args[i].(string)
// 			if !ok {
// 				l.logger.Error("Invalid key type, key must be a string", zap.Any("key", args[i]))
// 				continue
// 			}
// 			value := args[i+1]
// 			fields = append(fields, zap.Any(key, value))
// 		}
// 	}
// 	l.logger.Debug(msg, fields...)
// }

// // LogWarn logs a warning message with timestamp and optional context
// func (l *Logger) LogWarn(msg string, args ...interface{}) {
// 	var fields []zap.Field
// 	for i := 0; i < len(args); i += 2 {
// 		if i+1 < len(args) {
// 			key, ok := args[i].(string)
// 			if !ok {
// 				l.logger.Error("Invalid key type, key must be a string", zap.Any("key", args[i]))
// 				continue
// 			}
// 			value := args[i+1]
// 			fields = append(fields, zap.Any(key, value))
// 		}
// 	}
// 	l.logger.Warn(msg, fields...)
// }

// // LogError logs an error message with timestamp, details like function, filename, and line number
// func (l *Logger) LogError(msg string, err error, args ...interface{}) {
// 	pc, file, line, _ := runtime.Caller(1)
// 	funcName := runtime.FuncForPC(pc).Name()
// 	funcNameParts := strings.Split(funcName, ".")
// 	funcName = funcNameParts[len(funcNameParts)-1]

// 	var fields []zap.Field
// 	for i := 0; i < len(args); i += 2 {
// 		if i+1 < len(args) {
// 			key, ok := args[i].(string)
// 			if !ok {
// 				l.logger.Error("Invalid key type, key must be a string", zap.Any("key", args[i]))
// 				continue
// 			}
// 			value := args[i+1]
// 			fields = append(fields, zap.Any(key, value))
// 		}
// 	}

// 	l.logger.Error(msg,
// 		append(fields,
// 			zap.String("function", funcName),
// 			zap.String("file", file),
// 			zap.Int("line", line),
// 			zap.Error(err),
// 		)...,
// 	)
// }

// // WithContext adds context-specific fields (like request ID or session ID) to the logger
// func (l *Logger) WithContext(fields ...zap.Field) *Logger {
// 	newLogger := l.logger.With(fields...)
// 	return &Logger{logger: newLogger}
// }

// // Close the logger (flush any buffered log entries)
// func (l *Logger) Close() {
// 	_ = l.logger.Sync()
// }

package logger

import (
	"fmt"
	"os"
	"runtime"
	"strings"
	"time"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Logger struct holds the Zap logger instance
type Logger struct {
	logger *zap.Logger
}

// CustomEncoder configures the encoder to use RFC3339 timestamp format and colored log levels
func customEncoderConfig(isConsole bool, config Config) zapcore.EncoderConfig {
	encoderConfig := zap.NewProductionEncoderConfig()
	timeFormat := config.TimeFormat
	// Use provided time format or default to RFC3339 if not specified
	if timeFormat == "" {
		timeFormat = time.RFC3339
	}
	encoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout(timeFormat) // Custom timestamp format
	if isConsole {
		encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder // Colored log level for console output
	} else {
		encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder // Plain log level for file output
	}
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder // Include caller info (file and line)
	return encoderConfig
}

// NewLogger creates a new logger instance with the provided configuration
func NewLogger(config Config) (*Logger, error) {
	var level zapcore.Level
	switch strings.ToLower(config.LogLevel) {
	case "debug":
		level = zapcore.DebugLevel
	case "info":
		level = zapcore.InfoLevel
	case "warn":
		level = zapcore.WarnLevel
	case "error":
		level = zapcore.ErrorLevel
	case "fatal":
		level = zapcore.FatalLevel
	default:
		return nil, fmt.Errorf("invalid log level: %s", config.LogLevel)
	}

	// Set the log format (either JSON or plain text)
	var consoleEncoder zapcore.Encoder
	var fileEncoder zapcore.Encoder
	if config.JSONFormat {
		consoleEncoder = zapcore.NewJSONEncoder(customEncoderConfig(true, config)) // JSON format for console
		fileEncoder = zapcore.NewJSONEncoder(customEncoderConfig(false, config))   // JSON format for file
	} else {
		consoleEncoder = zapcore.NewConsoleEncoder(customEncoderConfig(true, config)) // Plain text format for console
		fileEncoder = zapcore.NewConsoleEncoder(customEncoderConfig(false, config))   // Plain text format for file
	}

	// Log rotation configuration
	var writer zapcore.WriteSyncer
	if config.EnableRotation {
		writer = zapcore.AddSync(&lumberjack.Logger{
			Filename:   config.LogFilePath,
			MaxSize:    config.MaxSize,    // 50MB
			MaxBackups: config.MaxBackups, //5
			MaxAge:     config.MaxAge,     // 7days
			Compress:   config.Compress,   // default true
		})
	} else {
		writer = zapcore.AddSync(&lumberjack.Logger{
			Filename: config.LogFilePath,
			MaxSize:  config.MaxSize, // 50MB
		})
	}

	// Create cores based on the LogOutput configuration
	var cores []zapcore.Core
	for _, output := range config.LogOutput {
		switch output {
		case "stdout":
			consoleCore := zapcore.NewCore(consoleEncoder, zapcore.Lock(os.Stdout), level)
			cores = append(cores, consoleCore)
		case "file":
			fileCore := zapcore.NewCore(fileEncoder, writer, level)
			cores = append(cores, fileCore)
		default:
			return nil, fmt.Errorf("invalid log output: %s", output)
		}
	}

	// Combine the cores into one logger
	core := zapcore.NewTee(cores...)

	// Create the logger instance
	logger := zap.New(core)

	// // If caller information is enabled, create a new logger with it
	// // if config.EnableCallerInfo {
	// // 	logger = logger.WithOptions(zap.AddCaller())
	// // }

	// // Apply the app name to the logger context
	// logger = logger.With(zap.String("app", config.AppName))
	// Add the app field to the logger context only once at initialization
	logger = logger.With(zap.String("app", config.AppName))

	// If caller information is enabled, create a new logger with it
	if config.EnableCallerInfo {
		logger = logger.WithOptions(zap.AddCaller())
	}

	return &Logger{logger: logger}, nil
}

// LogInfo logs an informational message with timestamp and optional context
func (l *Logger) LogInfo(msg string, args ...interface{}) {
	var fields []zap.Field
	for i := 0; i < len(args); i += 2 {
		if i+1 < len(args) {
			key, ok := args[i].(string)
			if !ok {
				l.logger.Error("Invalid key type, key must be a string", zap.Any("key", args[i]))
				continue
			}
			value := args[i+1]
			vaal := value.(string)

			println("key: ", key, ",value: ", string(vaal))
			fields = append(fields, zap.Any(key, value))
		}
	}
	fmt.Printf("msg: %s\n", msg)
	l.logger.Info(msg, fields...)
}

// LogDebug logs a debug message with timestamp and optional context
func (l *Logger) LogDebug(msg string, args ...interface{}) {
	var fields []zap.Field
	for i := 0; i < len(args); i += 2 {
		if i+1 < len(args) {
			key, ok := args[i].(string)
			if !ok {
				l.logger.Error("Invalid key type, key must be a string", zap.Any("key", args[i]))
				continue
			}
			value := args[i+1]
			fields = append(fields, zap.Any(key, value))
		}
	}
	l.logger.Debug(msg, fields...)
}

// LogWarn logs a warning message with timestamp and optional context
func (l *Logger) LogWarn(msg string, args ...interface{}) {
	var fields []zap.Field
	for i := 0; i < len(args); i += 2 {
		if i+1 < len(args) {
			key, ok := args[i].(string)
			if !ok {
				l.logger.Error("Invalid key type, key must be a string", zap.Any("key", args[i]))
				continue
			}
			value := args[i+1]
			fields = append(fields, zap.Any(key, value))
		}
	}
	l.logger.Warn(msg, fields...)
}

// LogError logs an error message with timestamp, details like function, filename, and line number
func (l *Logger) LogError(msg string, err error, args ...interface{}) {
	pc, file, line, _ := runtime.Caller(1)
	funcName := runtime.FuncForPC(pc).Name()
	funcNameParts := strings.Split(funcName, ".")
	funcName = funcNameParts[len(funcNameParts)-1]

	var fields []zap.Field
	for i := 0; i < len(args); i += 2 {
		if i+1 < len(args) {
			key, ok := args[i].(string)
			if !ok {
				l.logger.Error("Invalid key type, key must be a string", zap.Any("key", args[i]))
				continue
			}
			value := args[i+1]
			fields = append(fields, zap.Any(key, value))
		}
	}

	l.logger.Error(msg,
		append(fields,
			zap.String("function", funcName),
			zap.String("file", file),
			zap.Int("line", line),
			zap.Error(err),
		)...,
	)
}

// WithContext adds context-specific fields (like request ID or session ID) to the logger
func (l *Logger) WithContext(fields ...zap.Field) *Logger {
	newLogger := l.logger.With(fields...)
	return &Logger{logger: newLogger}
}

// Close the logger (flush any buffered log entries)
func (l *Logger) Close() {
	_ = l.logger.Sync()
}
