package requestid

import (
	"context"
	"net/http"

	"github.com/google/uuid"
)

type requestIdKey int

const RequestIdKey requestIdKey = 0

// Middleware returns a handler that sets request IDs to the request context
func Middleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			// add request ID to the request context so they can be added
			// to the log messages in the logger middleware
			ctx := r.Context()
			ctx = withRequestId(ctx, r)
			reqWithIDs := r.WithContext(ctx)

			next.ServeHTTP(w, reqWithIDs)
		}
		return http.HandlerFunc(fn)
	}
}

// withRequestId looks for an existing request id and generates a new one if none found.
// It returns a new context object containg the request id value
func withRequestId(ctx context.Context, r *http.Request) context.Context {
	reqId := getRequestID(r)
	if reqId == "" {
		reqId = uuid.New().String()
	}
	return context.WithValue(ctx, RequestIdKey, reqId)
}

// getRequestID grabs the request ID string off the X-Request-ID header
func getRequestID(r *http.Request) string {
	return r.Header.Get("X-Request-ID")
}
