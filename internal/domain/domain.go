package domain

type ctxKey string

const (
	LogLevelError   = "error"
	LogLevelWarning = "warn"
	LogLevelInfo    = "info"
	LogLevelDebug   = "debug"

	TraceIDCtxKey ctxKey = "trace-id"
)
