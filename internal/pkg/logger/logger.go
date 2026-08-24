// Package logger is a minimal leveled console logger.
//
// It replaces the previously used github.com/Aleksandr-Kai/logger module,
// whose repository is no longer reachable (deleted/renamed upstream), which
// made the project impossible to build.
package logger

import (
	"log"
	"os"
	"strings"
)

// Level is a logging severity level.
type Level int

// Severity levels, from most to least verbose.
const (
	Debug Level = iota
	Info
	Warn
	Error
)

// ParseLevel converts a level name (as found in the config file) into a Level.
// Unknown values default to Info.
func ParseLevel(s string) Level {
	switch strings.ToLower(s) {
	case "debug":
		return Debug
	case "info", "text":
		return Info
	case "warn", "warning":
		return Warn
	case "error":
		return Error
	default:
		return Info
	}
}

func (l Level) String() string {
	switch l {
	case Debug:
		return "DEBUG"
	case Info:
		return "INFO"
	case Warn:
		return "WARN"
	case Error:
		return "ERROR"
	default:
		return "INFO"
	}
}

// Logger writes leveled messages to the console, dropping anything below
// its GlobalLevel.
type Logger struct {
	level  Level
	logger *log.Logger
}

// NewLogger creates a Logger with Info as the default level.
func NewLogger() Logger {
	return Logger{
		level:  Info,
		logger: log.New(os.Stdout, "", log.LstdFlags),
	}
}

// GlobalLevel sets the minimum level that will be printed.
func (l *Logger) GlobalLevel(level Level) {
	l.level = level
}

// LogToConsole prints msg if level is at or above the logger's GlobalLevel.
func (l *Logger) LogToConsole(level Level, msg string) {
	if level < l.level {
		return
	}

	l.logger.Printf("[%s] %s", level, msg)
}

// LogToConsole is a package-level helper for logging before a Logger has
// been configured (always printed, regardless of level).
func LogToConsole(level Level, msg string) {
	log.Printf("[%s] %s", level, msg)
}
