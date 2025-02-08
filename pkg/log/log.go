package log

import (
	"fmt"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	DebugLevel = zapcore.DebugLevel
	InfoLevel  = zapcore.InfoLevel
)

var (
	logger      *zap.Logger
	atomicLevel zap.AtomicLevel
)

func Setup() {
	cfg := zap.NewProductionConfig()
	cfg.Level = zap.NewAtomicLevelAt(InfoLevel)
	logger = zap.Must(cfg.Build())
}

func SetLevel(level zapcore.Level) {
	atomicLevel.SetLevel(level)
}

func GetLogger() *zap.Logger {
	if logger == nil {
		Setup()
	}
	return logger
}

func GetComponentLogger(component string) *zap.Logger {
	logger := GetLogger().With(zap.String("component", component))
	return logger
}

func Debugf(format string, value ...any) {
	msg := fmt.Sprintf(format, value)
	logger.Debug(msg)
}

func Debug(msg string) {
	logger.Debug(msg)
}

func Errorf(format string, value ...any) {
	msg := fmt.Sprintf(format, value)
	logger.Error(msg)
}

func Error(msg string) {
	logger.Error(msg)
}

func Infof(format string, value ...any) {
	msg := fmt.Sprintf(format, value)
	logger.Info(msg)
}

func Info(msg string) {
	logger.Info(msg)
}

func Warnf(format string, value ...any) {
	msg := fmt.Sprintf(format, value)
	logger.Warn(msg)
}

func Warn(msg string) {
	logger.Warn(msg)
}

func Sync() {
	logger.Sync()
}
