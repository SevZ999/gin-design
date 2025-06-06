package logger

import "go.uber.org/zap"

type Logger interface {
	Error(string, ...zap.Field)
	Info(string, ...zap.Field)
}
