package logging

import (
	"fmt"
	"log"
	"os"
)

type Level int

const (
	OFF Level = iota
	FATAL
	ERROR
	WARN
	INFO
	DEBUG
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

// NewLogger creates a new logger object with the given prefix.
func GetLogger(module string) *Logger {
	logger := &Logger{
		Logger: log.New(os.Stdout, "", 0),
		Level:  DEBUG,
		Module: module,
	}
	logger.SetFlags(log.LstdFlags)

	return logger
}

func (l *Logger) SetLevel(level Level) {
	l.Level = level
}

func (l *Logger) logLevel(level Level, message string) {
	if l.Level < level {
		return
	}

	formatted := fmt.Sprintf("[%v] - %v - %v", level.String()[0:4], l.Module, message)
	log.Println(formatted)
}

func (l *Logger) Trace(message string) {
	l.logLevel(TRACE, message)
}

func (l *Logger) Debug(message string) {
	l.logLevel(DEBUG, message)
}

func (l *Logger) Info(message string) {
	l.logLevel(INFO, message)
}

func (l *Logger) Warn(message string) {
	l.logLevel(WARN, message)
}

func (l *Logger) Error(message string) {
	l.logLevel(ERROR, message)
}

func (l *Logger) Fatal(message string) {
	l.logLevel(FATAL, message)
}
