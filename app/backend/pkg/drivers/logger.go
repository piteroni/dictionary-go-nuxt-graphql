package drivers

import (
	"io"
	"log"
)

type appLogger struct {
	logger *log.Logger
}

func NewLogger(f io.Writer) *appLogger {
	logger := log.New(f, "", log.LstdFlags|log.Ldate)

	return &appLogger{
		logger: logger,
	}
}

func (l *appLogger) Errorf(format string, a ...interface{}) {
	l.error(format, a...)
}

func (l *appLogger) Error(err error) {
	l.error(err.Error())
}

func (l *appLogger) Warnf(format string, a ...interface{}) {
	l.warn(format, a...)
}

func (l *appLogger) Info(message string) {
	l.info(message)
}

func (l *appLogger) Infof(format string, a ...interface{}) {
	l.info(format, a...)
}

func (l *appLogger) error(format string, a ...interface{}) {
	l.logger.SetPrefix("[ERROR] ")
	l.logger.Printf(format, a...)
}

func (l *appLogger) warn(format string, a ...interface{}) {
	l.logger.SetPrefix("[WARN] ")
	l.logger.Printf(format, a...)
}

func (l *appLogger) info(format string, a ...interface{}) {
	l.logger.SetPrefix("[INFO] ")
	l.logger.Printf(format, a...)
}
