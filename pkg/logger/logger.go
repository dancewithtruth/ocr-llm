package logger

import (
	"context"

	"github.com/Wave-95/pgserver/middleware/requestid"
	"go.uber.org/zap"
)

type Logger interface {
	With(args ...interface{}) Logger

	WithRequestCtx(ctx context.Context) Logger

	WithoutCaller() Logger
	Debug(args ...interface{})
	// Info uses fmt.Sprint to construct and log a message at INFO level
	Info(args ...interface{})
	// Error uses fmt.Sprint to construct and log a message at ERROR level
	Error(args ...interface{})

	// Debugf uses fmt.Sprintf to construct and log a message at DEBUG level
	Debugf(format string, args ...interface{})
	// Infof uses fmt.Sprintf to construct and log a message at INFO level
	Infof(format string, args ...interface{})
	// Errorf uses fmt.Sprintf to construct and log a message at ERROR level
	Errorf(format string, args ...interface{})
}

type logger struct {
	*zap.SugaredLogger
}

// New creates a new logger using the default configuration
func New() Logger {
	l, _ := zap.NewProduction()
	return NewSugar(l)
}

// NewSugar returns a SugaredLogger and implements the Logger interface
func NewSugar(l *zap.Logger) Logger {
	return &logger{l.Sugar()}
}

// With returns a logger based off the root logger and decorates it with the arguments
func (l *logger) With(args ...interface{}) Logger {
	return &logger{l.SugaredLogger.With(args...)}
}

// WithoutCaller returns a logger that does not output the caller field and location of the calling code.
func (l *logger) WithoutCaller() Logger {
	return &logger{l.SugaredLogger.WithOptions(zap.WithCaller(false))}
}

// WithContext takes in a request context and appends request ID or any other header fields to the logger
func (l *logger) WithRequestCtx(ctx context.Context) Logger {
	l = withRequestID(ctx, l)
	return l
}

// withRequestID takes a context to find a request ID and append it to a logger
func withRequestID(ctx context.Context, l *logger) *logger {
	if reqId, ok := ctx.Value(requestid.RequestIdKey).(string); ok {
		return &logger{l.SugaredLogger.With("requestID", reqId)}
	}
	return l
}
