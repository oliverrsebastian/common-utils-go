package logger

import (
	"context"
	"encoding/json"
	"fmt"
	"go.uber.org/zap"
)

func Info(ctx context.Context, msg string, args ...interface{}) {
	parsedMsg := formatMsg(ctx, InfoLevel, msg, args...)
	zap.L().Info(parsedMsg)
}

func Debug(ctx context.Context, msg string, args ...interface{}) {
	parsedMsg := formatMsg(ctx, DebugLevel, msg, args...)
	zap.L().Debug(parsedMsg)
}

func Warn(ctx context.Context, msg string, args ...interface{}) {
	parsedMsg := formatMsg(ctx, WarnLevel, msg, args...)
	zap.L().Warn(parsedMsg)
}

func Error(ctx context.Context, msg string, args ...interface{}) {
	parsedMsg := formatMsg(ctx, ErrorLevel, msg, args...)
	zap.L().Error(parsedMsg)
}

func formatMsg(ctx context.Context, level Level, msg string, args ...interface{}) string {
	traceID := ctx.Value(ContextKey)
	messageMap := map[string]interface{}{
		ContextKey: traceID,
		"Level":    level,
		"Message":  fmt.Sprintf(msg, args...),
	}
	msgBytes, err := json.Marshal(messageMap)
	if err != nil {
		fmt.Printf("got err when formatting json log message")
	}
	return string(msgBytes)
}
