package three

import (
	"log"
	"os"
)

// Logger is used to extend log.Logger.
type Logger struct {
	*log.Logger
}

// NewLogger creates a new logger object with the given prefix.
func NewLogger(prefix string) *Logger {
	logger := &Logger{log.New(os.Stdout, prefix, 0)}
	logger.SetFlags(log.LstdFlags)

	return logger
}
