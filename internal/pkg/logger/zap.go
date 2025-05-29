// internal/pkg/logger/zap.go
package logger

import (
	"loan-admin/internal/config"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

type Logger struct {
	*zap.Logger
}

func NewZapLogger(cfg *config.Config) (*Logger, error) {
	var cores []zapcore.Core

	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.TimeKey = "timestamp"
	encoderConfig.LevelKey = "level"
	encoderConfig.NameKey = "logger"
	encoderConfig.CallerKey = "caller"
	encoderConfig.MessageKey = "message"
	encoderConfig.StacktraceKey = "stacktrace"
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder

	if cfg.Env == "dev" {
		consoleEncoder := zapcore.NewConsoleEncoder(encoderConfig)
		consoleCore := zapcore.NewCore(
			consoleEncoder,
			zapcore.AddSync(os.Stdout),
			getLogLevel(cfg.Log.Level),
		)
		cores = append(cores, consoleCore)
	}

	if cfg.Env == "prod" {
		jsonEncoder := zapcore.NewJSONEncoder(encoderConfig)

		hook := lumberjack.Logger{
			Filename:   "./log/app/app.log",
			MaxSize:    cfg.Log.MaxSize,
			MaxBackups: cfg.Log.MaxBackups,
			MaxAge:     cfg.Log.MaxAge,
			Compress:   cfg.Log.Compress,
		}

		fileCore := zapcore.NewCore(
			jsonEncoder,
			zapcore.AddSync(&hook),
			getLogLevel(cfg.Log.Level),
		)

		errorHook := lumberjack.Logger{
			Filename:   "./log/app/error.log",
			MaxSize:    cfg.Log.MaxSize,
			MaxBackups: cfg.Log.MaxBackups,
			MaxAge:     cfg.Log.MaxAge,
			Compress:   cfg.Log.Compress,
		}

		errorFilter := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
			return lvl >= zapcore.ErrorLevel
		})

		errorCore := zapcore.NewCore(
			jsonEncoder,
			zapcore.AddSync(&errorHook),
			errorFilter,
		)

		cores = append(cores, fileCore, errorCore)
	}

	combinedCore := zapcore.NewTee(cores...)

	logger := zap.New(
		combinedCore,
		zap.AddCaller(),
		zap.AddStacktrace(zapcore.PanicLevel),
		zap.Development(),
	)

	zap.ReplaceGlobals(logger)

	return &Logger{logger}, nil
}

func getLogLevel(level string) zapcore.Level {
	switch level {
	case "debug":
		return zap.DebugLevel
	case "info":
		return zap.InfoLevel
	case "warn":
		return zap.WarnLevel
	case "error":
		return zap.ErrorLevel
	default:
		return zap.InfoLevel
	}
}

func Log() *zap.Logger {
	return zap.L()
}
