package logger

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sync"
	"time"
)

type LogLevel int8

const (
	LogLevelDebug   LogLevel = iota // Debug level logs
	LogLevelInfo                    // Info level logs
	LogLevelWarning                 // Warning level logs
	LogLevelError                   // Error level logs
	LogLevelFatal                   // Fatal level logs
)

// String converts LogLevel to its string representation
func (l LogLevel) String() string {
	switch l {
	case LogLevelDebug:
		return "DEBUG"
	case LogLevelInfo:
		return "INFO"
	case LogLevelWarning:
		return "WARNING"
	case LogLevelError:
		return "ERROR"
	case LogLevelFatal:
		return "FATAL"
	default:
		return "UNKNOWN"
	}
}

// Logger defines the logging interface
type Logger interface {
	SetLogLevel(lvl LogLevel)
	SetOutput(filePath string) error
	Flush() error
	Debug(args ...interface{})
	Debugf(format string, args ...interface{})
	Info(args ...interface{})
	Infof(format string, args ...interface{})
	Warn(args ...interface{})
	Warnf(format string, args ...interface{})
	Error(args ...interface{})
	Errorf(format string, args ...interface{})
	Fatal(args ...interface{})
	Fatalf(format string, args ...interface{})
	Log(lvl LogLevel, args ...interface{})
	Logf(lvl LogLevel, format string, args ...interface{})
}

// LogExtended is the default implementation of the Logger interface
type LogExtended struct {
	mu        sync.RWMutex
	logger    *log.Logger
	bufWriter *bufio.Writer
	logLevel  LogLevel
	logFile   *os.File
}

// New creates a new instance of LogExtended with default settings
func New() *LogExtended {
	bufWriter := bufio.NewWriter(os.Stdout)
	return &LogExtended{
		logger:    log.New(bufWriter, "", log.Lmsgprefix),
		bufWriter: bufWriter,
		logLevel:  LogLevelInfo,
	}
}

// SetLogLevel updates the log level for the logger
func (l *LogExtended) SetLogLevel(lvl LogLevel) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.logLevel = lvl
}

// SetOutput sets the output destination for the logger
func (l *LogExtended) SetOutput(filePath string) error {
	l.mu.Lock()
	defer l.mu.Unlock()

	if l.bufWriter != nil {
		l.bufWriter.Flush()
	}

	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return err
	}

	if l.logFile != nil {
		l.logFile.Close()
	}

	bufWriter := bufio.NewWriter(file)
	l.logger.SetOutput(bufWriter)
	l.bufWriter = bufWriter
	l.logFile = file

	return nil
}

// Flush flushes any buffered log entries
func (l *LogExtended) Flush() error {
	l.mu.Lock()
	defer l.mu.Unlock()
	if l.bufWriter != nil {
		return l.bufWriter.Flush()
	}
	return nil
}

// formatMessage formats the log message with timestamp and level
func (l *LogExtended) formatMessage(lvl LogLevel, args ...interface{}) string {
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	return fmt.Sprintf("[%s] [%s] %v", timestamp, lvl.String(), fmt.Sprint(args...))
}

// formatMessagef formats the log message with timestamp and level using format string
func (l *LogExtended) formatMessagef(lvl LogLevel, format string, args ...interface{}) string {
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	msg := fmt.Sprintf(format, args...)
	return fmt.Sprintf("[%s] [%s] %s", timestamp, lvl.String(), msg)
}

// Log logs a message with the given log level if it meets the current logging threshold
func (l *LogExtended) Log(lvl LogLevel, args ...interface{}) {
	l.mu.Lock()
	defer l.mu.Unlock()

	if lvl < l.logLevel {
		return
	}

	l.logger.Println(l.formatMessage(lvl, args...))

	if lvl == LogLevelFatal {
		l.Flush()
		if l.logFile != nil {
			l.logFile.Close()
		}
		os.Exit(1)
	}
}

// Logf logs a formatted message with the given log level if it meets the current logging threshold
func (l *LogExtended) Logf(lvl LogLevel, format string, args ...interface{}) {
	l.mu.Lock()
	defer l.mu.Unlock()

	if lvl < l.logLevel {
		return
	}

	l.logger.Println(l.formatMessagef(lvl, format, args...))

	if lvl == LogLevelFatal {
		l.Flush()
		if l.logFile != nil {
			l.logFile.Close()
		}
		os.Exit(1)
	}
}

// Debug logs a message at debug level
func (l *LogExtended) Debug(args ...interface{}) { l.Log(LogLevelDebug, args...) }

// Debugf logs a formatted message at debug level
func (l *LogExtended) Debugf(format string, args ...interface{}) { l.Logf(LogLevelDebug, format, args...) }

// Info logs a message at info level
func (l *LogExtended) Info(args ...interface{}) { l.Log(LogLevelInfo, args...) }

// Infof logs a formatted message at info level
func (l *LogExtended) Infof(format string, args ...interface{}) { l.Logf(LogLevelInfo, format, args...) }

// Warn logs a message at warning level
func (l *LogExtended) Warn(args ...interface{}) { l.Log(LogLevelWarning, args...) }

// Warnf logs a formatted message at warning level
func (l *LogExtended) Warnf(format string, args ...interface{}) { l.Logf(LogLevelWarning, format, args...) }

// Error logs a message at error level
func (l *LogExtended) Error(args ...interface{}) { l.Log(LogLevelError, args...) }

// Errorf logs a formatted message at error level
func (l *LogExtended) Errorf(format string, args ...interface{}) { l.Logf(LogLevelError, format, args...) }

// Fatal logs a message at fatal level and terminates the program
func (l *LogExtended) Fatal(args ...interface{}) { l.Log(LogLevelFatal, args...) }

// Fatalf logs a formatted message at fatal level and terminates the program
func (l *LogExtended) Fatalf(format string, args ...interface{}) { l.Logf(LogLevelFatal, format, args...) }