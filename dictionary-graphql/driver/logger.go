package driver

import (
	"fmt"
	"io"
	"log"
)

type AppLogger struct {
	logger *log.Logger
}

func NewLogger(f io.Writer) *AppLogger {
	logger := log.New(f, "", log.LstdFlags|log.Ldate|log.Llongfile)

	return &AppLogger{
		logger: logger,
	}
}

func (l *AppLogger) Print(message string) {
	l.logger.Println(message)
}

func (l *AppLogger) Printf(format string, v ...interface{}) {
	l.logger.Printf(fmt.Sprintf(format, v...))
}

func (l *AppLogger) Error(err error) {
	l.logger.Printf("%+v\n", err)
}
