package logger

import (
	"log/slog"
	"os"
)

/**
* All of ths functions have to be improved to send the logs to an s3 bucket
**/

type Logger struct {
	slog *slog.Logger
}

// New Logger creates a new Logger that can be used easily
func NewLogger() *Logger {

	handler := slog.NewJSONHandler(os.Stdout, nil)

	return &Logger{slog: slog.New(handler)}

}

// Use info in normal application progress / noteworthy events that are not problems.
func (l Logger) Info(msg string, args ...any) {
	l.slog.Info(msg, args...)
}

// Use Debug when you want detailed diagnostic information useful only while developing
// or troubleshooting. Usually disabled in production.

func (l Logger) Debug(msg string, args ...any) {
	l.slog.Info(msg, args...)
}

// Use Error in  failures that require attention. The operation failed, but the
// process may keep running.

func (l Logger) Error(msg string, args ...any) {
	l.slog.Info(msg, args...)
}

// Use warn when unexpected or suboptimal situations that are not errors (yet).
// The system can continue, but someone should look into it.
func (l Logger) Warn(msg string, args ...any) {
	l.slog.Info(msg, args...)
}
