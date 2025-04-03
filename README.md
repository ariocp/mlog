```
go get github.com/ariocp/mlog
```

```go
// SetLogLevel updates the log level for the logger
func SetLogLevel(lvl logger.LogLevel) { mlog.SetLogLevel(lvl) }

// SetOutput sets the output destination for the logger
func SetOutput(filePath string) error { return mlog.SetOutput(filePath) }

// Flush flushes any buffered log entries
func Flush() error { return mlog.Flush() }

// Debug logs a message at debug level
func Debug(args ...interface{}) { mlog.Debug(args...) }

// Debugf logs a formatted message at debug level
func Debugf(format string, args ...interface{}) { mlog.Debugf(format, args...) }

// Info logs a message at info level
func Info(args ...interface{}) { mlog.Info(args...) }

// Infof logs a formatted message at info level
func Infof(format string, args ...interface{}) { mlog.Infof(format, args...) }

// Warn logs a message at warning level
func Warn(args ...interface{}) { mlog.Warn(args...) }

// Warnf logs a formatted message at warning level
func Warnf(format string, args ...interface{}) { mlog.Warnf(format, args...) }

// Error logs a message at error level
func Error(args ...interface{}) { mlog.Error(args...) }

// Errorf logs a formatted message at error level
func Errorf(format string, args ...interface{}) { mlog.Errorf(format, args...) }

// Fatal logs a message at fatal level and terminates the program
func Fatal(args ...interface{}) { mlog.Fatal(args...) }

// Fatalf logs a formatted message at fatal level and terminates the program
func Fatalf(format string, args ...interface{}) { mlog.Fatalf(format, args...) }

// Log logs a message with the given log level if it meets the current logging threshold
func Log(lvl logger.LogLevel, args ...interface{}) { mlog.Log(lvl, args...) }

// Logf logs a formatted message with the given log level if it meets the current logging threshold
func Logf(lvl logger.LogLevel, format string, args ...interface{}) { mlog.Logf(lvl, format, args...) }
```
