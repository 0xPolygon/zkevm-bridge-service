package utils

const (
	CtxTraceID string = "traceID"
	TraceID           = "traceID"
	traceIDLen        = 16
)

// GenerateTraceID generates a random trace ID.
func GenerateTraceID() string {
	return generateRandomString(traceIDLen)
}
