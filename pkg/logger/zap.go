package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type ZapLogger struct {
	log *zap.Logger
}

func (l *ZapLogger) Debug(args ...interface{}) {
	l.log.Debug(args[0].(string))
}
func (l *ZapLogger) Info(args ...interface{}) {
	l.log.Info(args[0].(string))
}
func (l *ZapLogger) Warn(args ...interface{}) {
	l.log.Warn(args[0].(string))
}
func (l *ZapLogger) Error(args ...interface{}) {
	l.log.Error(args[0].(string))
}

func CreateZapLogger() *ZapLogger {
	cfg := zap.Config{
		Level:       zap.NewAtomicLevelAt(zap.DebugLevel),
		Encoding:    "json",
		OutputPaths: []string{"stdout"},
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey:  "msg",
			LevelKey:    "level",
			EncodeLevel: zapcore.LowercaseLevelEncoder,
			TimeKey:     "time",
			EncodeTime:  zapcore.ISO8601TimeEncoder,
		},
	}
	originalLogger, _ := cfg.Build()
	defer originalLogger.Sync()
	return &ZapLogger{
		log: originalLogger,
	}
}
