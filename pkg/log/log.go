package log

import (
	"fmt"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var logger *zap.Logger

func Setup() {
	cfg := zap.NewProductionConfig()
	logger = zap.Must(cfg.Build())
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
