package request

import (
	"context"
	"net/http"

	"github.com/google/uuid"
)

type requestIdKey int
type correlationIdKey int

const (
	RequestIdKey     requestIdKey     = 0
	CorrelationIdKey correlationIdKey = 0
)

// Middleware returns a handler that sets request IDs to the request context
func Middleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// add request ID to the request context so they can be added
			// to the log messages in the logger middleware
			ctx := r.Context()
			ctx = withIDs(ctx, r)
			reqWithIDs := r.WithContext(ctx)

			next.ServeHTTP(w, reqWithIDs)
		})
	}
}

// withIDs looks for an existing request id and generates a new one if none found.
// It returns a new context object containg the request id value
func withIDs(ctx context.Context, r *http.Request) context.Context {
	reqId := getRequestID(r)
	corrId := getCorrelationID(r)
	if reqId == "" {
		reqId = uuid.NewString()
	}
	if corrId == "" {
		corrId = uuid.NewString()
	}
	ctx = context.WithValue(ctx, RequestIdKey, reqId)
	ctx = context.WithValue(ctx, CorrelationIdKey, corrId)
	return ctx
}

// getRequestID grabs the request ID string off the X-Request-ID header
func getRequestID(r *http.Request) string {
	return r.Header.Get("X-Request-ID")
}

// getCorrelationId grabs the correlation ID string off the X-Correlation-ID header
// The correlation id groups together multiple request ids
func getCorrelationID(r *http.Request) string {
	return r.Header.Get("X-Correlation-ID")
}
