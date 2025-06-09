package logger

import (
	"context"

	"go.uber.org/zap"
)

type Logger interface {
	Error(context.Context, string, ...zap.Field)
	Info(context.Context, string, ...zap.Field)
	Warn(context.Context, string, ...zap.Field)
}
