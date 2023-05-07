package logMiddleware

import (
	"context"
	"net/http"
	"time"

	"github.com/Wave-95/pgserver/middleware/request"
	"github.com/Wave-95/pgserver/pkg/logger"
)

// logResponseWriter wraps http.ResponseWriter in order to capture the status code and bytes written
type logResponseWriter struct {
	http.ResponseWriter
	statusCode   int
	bytesWritten int
}

// Sets the status code to logResponseWriter when WriteHeader is called in http package
func (rw *logResponseWriter) WriteHeader(statusCode int) {
	rw.statusCode = statusCode
	rw.ResponseWriter.WriteHeader(statusCode)
}

// Sets the bytes written to logResponseWriter when Write is called in http package
func (rw *logResponseWriter) Write(p []byte) (int, error) {
	n, err := rw.ResponseWriter.Write(p)
	rw.bytesWritten = n
	return n, err
}

type RequestLogger int

const RequestLoggerKey RequestLogger = 0

// Middleware injects a logger and builds an http handler. It finds the requestId and correlationId through the
// request context and creates a new logger with those fields. The logger is set to the request context so that
// it is available to any downstream handler. This middleware starts a timer for the request and logs it out at the end.
func Middleware(logger logger.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()

			// Get request ID and correlation ID from context and append fields to logger
			ctx := r.Context()
			requestLogger := logger
			if reqID, ok := ctx.Value(request.RequestIdKey).(string); ok {
				requestLogger = requestLogger.With("requestID", reqID)
			}
			if corrID, ok := ctx.Value(request.CorrelationIdKey).(string); ok {
				requestLogger = requestLogger.With("correlationID", corrID)
			}

			// Set logger to request context so that the logger always contains the
			// request ID and correlation ID
			ctx = context.WithValue(ctx, RequestLoggerKey, requestLogger)
			r = r.WithContext(ctx)

			// Wrap the ResponseWriter to capture the status code and bytes written
			wrappedWriter := &logResponseWriter{
				ResponseWriter: w,
				statusCode:     http.StatusOK,
				bytesWritten:   0,
			}

			next.ServeHTTP(wrappedWriter, r)

			defer func() {
				requestLogger.
					WithoutCaller().
					With("duration", time.Since(start).Milliseconds(), "bytes written", wrappedWriter.bytesWritten).
					Infof("%s %s %v", r.Method, r.URL.Path, wrappedWriter.statusCode)
			}()
		}
		return http.HandlerFunc(fn)
	}
}
