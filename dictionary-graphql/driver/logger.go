package driver

import (
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

func (l *AppLogger) Errorf(format string, a ...interface{}) {
	l.error(format, a...)
}

func (l *AppLogger) Error(err error) {
	l.error(err.Error())
}

func (l *AppLogger) Warnf(format string, a ...interface{}) {
	l.warn(format, a...)
}

func (l *AppLogger) Warn(message string) {
	l.warn(message)
}

func (l *AppLogger) Info(message string) {
	l.info(message)
}

func (l *AppLogger) Infof(format string, a ...interface{}) {
	l.info(format, a...)
}

func (l *AppLogger) error(format string, a ...interface{}) {
	l.logger.SetPrefix("[ERROR] ")
	l.logger.Printf(format, a...)
}

func (l *AppLogger) warn(format string, a ...interface{}) {
	l.logger.SetPrefix("[WARN] ")
	l.logger.Printf(format, a...)
}

func (l *AppLogger) info(format string, a ...interface{}) {
	l.logger.SetPrefix("[INFO] ")
	l.logger.Printf(format, a...)
}
