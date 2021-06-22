package logger

import (
	"os"
	"time"

	"github.com/sirupsen/logrus"
)

type LogrusLogger struct {
	log *logrus.Logger
}

func (l *LogrusLogger) Debug(args ...interface{}) {
	l.log.Debug(args...)
}
func (l *LogrusLogger) Info(args ...interface{}) {
	l.log.Info(args...)
}
func (l *LogrusLogger) Warn(args ...interface{}) {
	l.log.Warn(args...)
}
func (l *LogrusLogger) Error(args ...interface{}) {
	l.log.Error(args...)
}

func CreateLogger() *LogrusLogger {
	originalLogger := logrus.Logger{}
	originalLogger.SetFormatter(&logrus.JSONFormatter{
		DisableHTMLEscape: true,
		TimestampFormat:   time.RFC3339,
	})
	originalLogger.SetLevel(logrus.DebugLevel)
	originalLogger.SetOutput(os.Stdout)
	return &LogrusLogger{
		log: &originalLogger,
	}
}
