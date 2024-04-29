package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const defaultLevel = zapcore.InfoLevel

func NewLogger(level ...zapcore.Level) *zap.Logger {
	logLevel := defaultLevel
	if len(level) > 0 {
		logLevel = level[0]
	}

	config := &zap.Config{
		Level:            zap.NewAtomicLevelAt(logLevel),
		Encoding:         "json",
		EncoderConfig:    getEncoderConfig(),
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
	}

	plain, err := config.Build(zap.AddStacktrace(zap.DPanicLevel))
	if err != nil {
		plain = zap.NewNop()
	}
	return plain
}

func getEncoderConfig() zapcore.EncoderConfig {
	return zapcore.EncoderConfig{
		TimeKey:       "time",
		LevelKey:      "severity",
		NameKey:       "logger",
		CallerKey:     "caller",
		MessageKey:    "message",
		StacktraceKey: "stacktrace",
		LineEnding:    zapcore.DefaultLineEnding,
		EncodeLevel:   encodeLevel(),
		EncodeTime:    zapcore.ISO8601TimeEncoder,
		EncodeCaller:  zapcore.ShortCallerEncoder,
	}
}

func encodeLevel() zapcore.LevelEncoder {
	return func(l zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
		switch l {
		case zapcore.DebugLevel:
			enc.AppendString("DEBUG")
		case zapcore.InfoLevel:
			enc.AppendString("INFO")
		case zapcore.WarnLevel:
			enc.AppendString("WARNING")
		case zapcore.ErrorLevel:
			enc.AppendString("ERROR")
		case zapcore.DPanicLevel:
			enc.AppendString("DEV_PANIC")
		case zapcore.PanicLevel:
			enc.AppendString("PANIC")
		case zapcore.FatalLevel:
			enc.AppendString("FATAL")
		}
	}
}
