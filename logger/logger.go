package logger

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"sync"
	"time"
)

type LogLevel int8

const (
	LogLevelDebug LogLevel = iota
	LogLevelInfo
	LogLevelWarning
	LogLevelError
	LogLevelFatal
)

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

type Logger interface {
	SetLogLevel(level LogLevel)
	SetOutput(filepath string) error
	Flush() error
	Close() error

	Log(level LogLevel, args ...interface{})
	Logf(format string, level LogLevel, args ...interface{})

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
}

type LogExtended struct {
	bufferWriter *bufio.Writer
	logFile      *os.File
	logLevel     LogLevel
	logger       *log.Logger
	mu           sync.RWMutex
}

func New() *LogExtended {
	bufferWriter := bufio.NewWriter(os.Stdout)
	return &LogExtended{
		logLevel:     LogLevelInfo,
		bufferWriter: bufferWriter,
		logger:       log.New(bufferWriter, "", 0),
	}
}

func (l *LogExtended) SetLogLevel(level LogLevel) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.logLevel = level
}

func (l *LogExtended) SetOutput(filepath string) error {
	l.mu.Lock()
	defer l.mu.Unlock()

	if l.bufferWriter != nil {
		l.bufferWriter.Flush()
	}
	if l.logFile != nil {
		l.logFile.Close()
	}

	file, err := os.OpenFile(filepath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return err
	}

	multiWriter := io.MultiWriter(file, os.Stdout)
	bufferWriter := bufio.NewWriter(multiWriter)

	l.logger = log.New(bufferWriter, "", 0)
	l.bufferWriter = bufferWriter
	l.logFile = file

	return nil
}

func (l *LogExtended) Flush() error {
	l.mu.Lock()
	defer l.mu.Unlock()
	if l.bufferWriter != nil {
		err := l.bufferWriter.Flush()
		if err != nil {
			return err
		}
	}
	if l.logFile != nil {
		return l.logFile.Sync()
	}
	return nil
}

func (l *LogExtended) Close() error {
	l.mu.Lock()
	defer l.mu.Unlock()

	var err error
	if l.bufferWriter != nil {
		if e := l.bufferWriter.Flush(); e != nil {
			err = e
		}
	}
	if l.logFile != nil {
		if e := l.logFile.Close(); e != nil {
			err = e
		}
	}
	return err
}

func (l *LogExtended) formatMessage(level LogLevel, args ...interface{}) string {
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	return fmt.Sprintf("[%s] [%s] %v", timestamp, level.String(), fmt.Sprint(args...))
}

func (l *LogExtended) formatMessagef(format string, level LogLevel, args ...interface{}) string {
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	return fmt.Sprintf("[%s] [%s] %s", timestamp, level.String(), fmt.Sprintf(format, args...))
}

func (l *LogExtended) Log(level LogLevel, args ...interface{}) {
	l.mu.RLock()

	if level < l.logLevel {
		l.mu.RUnlock()
		return
	}

	l.mu.RUnlock()

	l.mu.Lock()

	l.logger.Println(l.formatMessage(level, args...))

	if level == LogLevelFatal {
		l.Flush()
		if l.logFile != nil {
			l.logFile.Close()
		}
		os.Exit(1)
	}

	l.mu.Unlock()
}

func (l *LogExtended) Logf(format string, level LogLevel, args ...interface{}) {
	l.mu.RLock()

	if level < l.logLevel {
		l.mu.RUnlock()
		return
	}

	l.mu.RUnlock()

	l.mu.Lock()

	l.logger.Println(l.formatMessagef(format, level, args...))

	if level == LogLevelFatal {
		l.Flush()
		if l.logFile != nil {
			l.logFile.Close()
		}
		os.Exit(1)
	}

	l.mu.Unlock()
}

func (l *LogExtended) Debug(args ...interface{})  { l.Log(LogLevelDebug, args...) }
func (l *LogExtended) Debugf(format string, args ...interface{}) { l.Logf(format, LogLevelDebug, args...) }
func (l *LogExtended) Info(args ...interface{})   { l.Log(LogLevelInfo, args...) }
func (l *LogExtended) Infof(format string, args ...interface{})  { l.Logf(format, LogLevelInfo, args...) }
func (l *LogExtended) Warn(args ...interface{})   { l.Log(LogLevelWarning, args...) }
func (l *LogExtended) Warnf(format string, args ...interface{})  { l.Logf(format, LogLevelWarning, args...) }
func (l *LogExtended) Error(args ...interface{})  { l.Log(LogLevelError, args...) }
func (l *LogExtended) Errorf(format string, args ...interface{}) { l.Logf(format, LogLevelError, args...) }
func (l *LogExtended) Fatal(args ...interface{})  { l.Log(LogLevelFatal, args...) }
func (l *LogExtended) Fatalf(format string, args ...interface{}) { l.Logf(format, LogLevelFatal, args...) }
