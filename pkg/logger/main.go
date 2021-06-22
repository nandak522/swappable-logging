package logger

import "fmt"

type Logger interface {
	Debug(args ...interface{})
	Info(args ...interface{})
	Warn(args ...interface{})
	Error(args ...interface{})
}

func CreateLogger(loggerName string) Logger {
	var log Logger
	switch loggerName {
	case "logrus":
		log = CreateLogrusLogger()
	case "zap":
		log = CreateZapLogger()
	default:
		fmt.Println("Invalid logger detected. Defaulting to logrus")
		log = CreateLogrusLogger()
	}
	return log
}
