/*
Package logging provides basic logging features (modules, log levels).
*/
package logging

import (
	"fmt"
	"log"
	"os"
)

// Level defines the log level. Can be one of the following:
// * OFF
// * FATAL
// * ERROR
// * WARN
// * INFO
// * DEBUG
// * TRACE
type Level int

const (
	// OFF disables logging
	OFF Level = iota
	// FATAL logs fatals
	FATAL
	// ERROR logs at least errors
	ERROR
	// WARN logs at least warnings
	WARN
	// INFO logs at least infos
	INFO
	// DEBUG logs at least debug output
	DEBUG
	// TRACE logs everything (verbose!)
	TRACE
)

var levelNames = []string{
	"OFF",
	"FATAL",
	"ERROR",
	"WARNING",
	"INFO",
	"DEBUG",
	"TRACE",
}

func (ll Level) String() string {
	return levelNames[ll]
}

// Logger is used to extend log.Logger.
type Logger struct {
	*log.Logger
	Level  Level
	Module string
}

// GetLogger creates a new logger object with the given prefix.
func GetLogger(module string) *Logger {
	logger := &Logger{
		Logger: log.New(os.Stdout, "", 0),
		Level:  DEBUG,
		Module: module,
	}
	logger.SetFlags(log.LstdFlags)

	return logger
}

// SetLevel sets the current log leve.
// Messages with a lower level than the given level
// will be omitted.
func (l *Logger) SetLevel(level Level) {
	l.Level = level
}

// Trace logs trace level messages.
func (l *Logger) Trace(message string) {
	l.logLevel(TRACE, message)
}

// Debug logs debug level messages.
func (l *Logger) Debug(message string) {
	l.logLevel(DEBUG, message)
}

// Info logs info level messages.
func (l *Logger) Info(message string) {
	l.logLevel(INFO, message)
}

// Warn logs warn level messages.
func (l *Logger) Warn(message string) {
	l.logLevel(WARN, message)
}

// Error logs error level messages.
func (l *Logger) Error(message string) {
	l.logLevel(ERROR, message)
}

// Fatal logs fatal level messages.
func (l *Logger) Fatal(message string) {
	l.logLevel(FATAL, message)
}

func (l *Logger) logLevel(level Level, message string) {
	if l.Level < level {
		return
	}

	formatted := fmt.Sprintf("[%v] - %v - %v", level.String()[0:4], l.Module, message)
	log.Println(formatted)
}
