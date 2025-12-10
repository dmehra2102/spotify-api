package util

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Logger *zap.Logger

func InitLogger(logLevel, logFormat string) error {
	config := zap.NewProductionConfig()

	switch logLevel {
	case "debug":
		config.Level = zap.NewAtomicLevelAt(zapcore.DebugLevel)
	case "warn":
		config.Level = zap.NewAtomicLevelAt(zapcore.WarnLevel)
	case "error":
		config.Level = zap.NewAtomicLevelAt(zapcore.ErrorLevel)
	default:
		config.Level = zap.NewAtomicLevelAt(zapcore.InfoLevel)
	}

	if logFormat == "json" {
		config.Encoding = "json"
	} else {
		config.Encoding = "console"
	}

	var err error
	Logger, err = config.Build()
	if err != nil {
		return err
	}

	return nil
}

func Close() error {
	if Logger != nil {
		return Logger.Sync()
	}
	return nil
}

func Info(msg string, fields ...zap.Field) {
	if Logger != nil {
		Logger.Info(msg, fields...)
	} else {
		os.Stdout.WriteString(msg + "\n")
	}
}

func Error(msg string, fields ...zap.Field) {
	if Logger != nil {
		Logger.Error(msg, fields...)
	} else {
		os.Stderr.WriteString(msg + "\n")
	}
}

func Debug(msg string, fields ...zap.Field) {
	if Logger != nil {
		Logger.Debug(msg, fields...)
	}
}
