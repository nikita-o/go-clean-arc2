package logger

import (
	"context"
	"go.uber.org/zap"
	"sync"
)

type ctxKeyType string

const CtxKey ctxKeyType = "logger"

var (
	globalLogger *zap.Logger
	once         sync.Once
)

func GetCtxLogger(ctx context.Context) *zap.Logger {
	if logger, ok := ctx.Value(CtxKey).(*zap.Logger); ok {
		return logger
	}
	return GetLogger()
}

func GetLogger() *zap.Logger {
	return NewLogger()
}

func NewLogger() *zap.Logger {
	once.Do(func() {
		loggerConfig := zap.NewProductionConfig()
		loggerConfig.OutputPaths = []string{"stdout"}
		logger, err := loggerConfig.Build()
		if err != nil {
			panic(err)
		}
		globalLogger = logger
	})
	return globalLogger
}
