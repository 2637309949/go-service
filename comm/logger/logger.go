// Package log provides a log interface
package logger

import (
	"go-micro.dev/v5/logger"
)

var (
	// Default logger.
	DefaultLogger Logger = NewLogger()

	// Default logger helper.
	DefaultHelper *Helper = NewHelper(DefaultLogger)
)

// Logger is a generic logging interface.
type Logger interface {
	// Init initializes options
	Init(options ...logger.Option) error
	// The Logger options
	Options() logger.Options
	// Fields set fields to always be logged
	Fields(fields map[string]interface{}) Logger
	// Log writes a log entry
	Log(level logger.Level, v ...interface{})
	// Logf writes a formatted log entry
	Logf(level logger.Level, format string, v ...interface{})
	// String returns the name of logger
	String() string
}

func Init(opts ...logger.Option) error {
	return DefaultLogger.Init(opts...)
}

func Fields(fields map[string]interface{}) Logger {
	return DefaultLogger.Fields(fields)
}

func Log(level logger.Level, v ...interface{}) {
	DefaultLogger.Log(level, v...)
}

func Logf(level logger.Level, format string, v ...interface{}) {
	DefaultLogger.Logf(level, format, v...)
}

func String() string {
	return DefaultLogger.String()
}

func LoggerOrDefault(l Logger) Logger {
	if l == nil {
		return DefaultLogger
	}
	return l
}
