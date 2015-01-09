package three

import (
	"log"
	"os"
)

type Logger struct {
	*log.Logger
}

func NewLogger(prefix string) *Logger {
	logger := &Logger{log.New(os.Stdout, prefix, 0)}
	logger.SetFlags(log.LstdFlags)

	return logger
}
