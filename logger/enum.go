package logger

const ContextKey = "TraceID"

type Level string

const (
	InfoLevel  Level = "INFO"
	DebugLevel Level = "DEBUG"
	WarnLevel  Level = "WARN"
	ErrorLevel Level = "ERROR"
)
