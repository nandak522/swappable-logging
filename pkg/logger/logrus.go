package logger

import (
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
	return &LogrusLogger{
		log: &logrus.Logger{},
	}
}
