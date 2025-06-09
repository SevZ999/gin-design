// internal/pkg/logger/zap.go
package logger

import (
	"context"
	"gin-design/internal/config"
	"os"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

type ctxKey string

const requestIDKey ctxKey = "request-id"

type ZapLogger struct {
	logger *zap.Logger
}

func NewZapLogger(cfg *config.Config) (*ZapLogger, error) {
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
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder // 显示调用者信息

	// 关键配置：SkipFrameCount 跳过封装层代码帧
	encoderConfig.EncodeCaller = func(caller zapcore.EntryCaller, enc zapcore.PrimitiveArrayEncoder) {
		// 跳过 2 层：1. 封装工具层 2. Zap 内部层
		enc.AppendString(caller.TrimmedPath())
	}

	if cfg.Env == "debug" {
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

	return &ZapLogger{logger: logger}, nil
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

func (l *ZapLogger) GetLogger() *zap.Logger {
	return l.logger
}

func (l *ZapLogger) Error(ctx context.Context, msg string, fields ...zap.Field) {
	// 提取 RequestID（如果有）
	if requestID, ok := ctx.Value(gin.ContextRequestKey).(string); ok {
		fields = append(fields, zap.String("requestID", requestID))
	}

	l.logger.Error(msg, fields...)
}

func (l *ZapLogger) Info(ctx context.Context, msg string, fields ...zap.Field) {
	// 提取 RequestID（如果有）
	if requestID, ok := ctx.Value(gin.ContextRequestKey).(string); ok {
		fields = append(fields, zap.String("requestID", requestID))
	}

	// 调用底层 Logger，Caller 信息自动追溯到业务代码
	l.logger.Info(msg, fields...)
}

func (l *ZapLogger) Warn(ctx context.Context, msg string, fields ...zap.Field) {
	// 提取 RequestID（如果有）
	if requestID, ok := ctx.Value(gin.ContextRequestKey).(string); ok {
		fields = append(fields, zap.String("requestID", requestID))
	}

	l.logger.Warn(msg, fields...)
}
