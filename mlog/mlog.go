package mlog

import "github.com/ariocp/mlog/logger"

var mlog = logger.New()

// SetLogLevel sets the minimum log level for messages to be written
// Messages with a log level lower than the set level will not be logged
func SetLogLevel(level logger.LogLevel) { mlog.SetLogLevel(level) }

// SetOutput sets a new log file for output. It flushes and closes the previous file if necessary
// If a file path is specified, it opens the file for appending logs and writes to both the file and stdout
func SetOutput(filepath string) error { return mlog.SetOutput(filepath) }

// Flush forces all buffered logs to be written to the output
// This ensures that all log messages are actually written to the output, including any buffered data
func Flush() error { return mlog.Flush() }

// Close flushes any buffered log entries and closes the log file if one was opened
// It's important to call this before exiting the program to ensure all logs are written
func Close() error { return mlog.Close() }

// Log writes a log message at the specified log level
// It checks if the log level is at or above the current configured log level and writes the message
func Log(level logger.LogLevel, args ...interface{}) { mlog.Log(level, args...) }

// Logf writes a formatted log message at the specified log level
// It checks if the log level is at or above the current configured log level and writes the formatted message
func Logf(format string, level logger.LogLevel, args ...interface{}) { mlog.Logf(format, level, args...) }

// Debug logs a debug message at the DEBUG level
// These messages are typically used for debugging purposes and are only logged when the log level includes DEBUG or lower
func Debug(args ...interface{}) { mlog.Debug(args...) }

// Debugf logs a formatted debug message at the DEBUG level
// These messages are typically used for debugging purposes and are only logged when the log level includes DEBUG or lower
func Debugf(format string, args ...interface{}) { mlog.Debugf(format, args...) }

// Info logs an informational message at the INFO level
// These messages provide general information about the application's state or execution
func Info(args ...interface{}) { mlog.Info(args...) }

// Infof logs a formatted informational message at the INFO level
// These messages provide general information about the application's state or execution
func Infof(format string, args ...interface{}) { mlog.Infof(format, args...) }

// Warn logs a warning message at the WARNING level
// These messages indicate a potential issue or something that may require attention but doesn't stop execution
func Warn(args ...interface{}) { mlog.Warn(args...) }

// Warnf logs a formatted warning message at the WARNING level
// These messages indicate a potential issue or something that may require attention but doesn't stop execution
func Warnf(format string, args ...interface{}) { mlog.Warnf(format, args...) }

// Error logs an error message at the ERROR level
// These messages indicate an error that occurred during execution, typically representing a failure or exception
func Error(args ...interface{}) { mlog.Error(args...) }

// Errorf logs a formatted error message at the ERROR level
// These messages indicate an error that occurred during execution, typically representing a failure or exception
func Errorf(format string, args ...interface{}) { mlog.Errorf(format, args...) }

// Fatal logs a fatal error message at the FATAL level and exits the program
// This level of logging indicates a critical error that causes the program to terminate immediately
func Fatal(args ...interface{}) { mlog.Fatal(args...) }

// Fatalf logs a formatted fatal error message at the FATAL level and exits the program
// This level of logging indicates a critical error that causes the program to terminate immediately
func Fatalf(format string, args ...interface{}) { mlog.Fatalf(format, args...) }
