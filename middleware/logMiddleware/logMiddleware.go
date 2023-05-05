package logMiddleware

import (
	"net/http"
	"time"

	"github.com/Wave-95/pgserver/pkg/logger"
	"github.com/go-chi/chi/middleware"
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

// Middleware injects a logger and builds an http handler. It finds the requestId and correlationId through the
// request context and add its to the logger. It also starts a timer for the request and logs it out.
func Middleware(logger logger.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()

			// Get requestid from context and append field to logger
			ctx := r.Context()
			logger = logger.WithRequestCtx(ctx)

			// Wrap the ResponseWriter to capture the status code and bytes written
			wrappedWriter := &logResponseWriter{
				ResponseWriter: w,
				statusCode:     http.StatusOK,
				bytesWritten:   0,
			}

			next.ServeHTTP(wrappedWriter, r)

			defer func() {
				logger.
					WithoutCaller().
					With("duration", time.Since(start).Milliseconds(), "bytes written", wrappedWriter.bytesWritten).
					Infof("%s %s %v", r.Method, r.URL.Path, wrappedWriter.statusCode)
			}()
		}
		return http.HandlerFunc(fn)
	}
}

func ReqIdToLogger(r *http.Request) {
	ctx := r.Context()
	reqId, ok := ctx.Value(middleware.RequestIDKey).(string)
	if reqId != "" && ok {
		//append reqId to logger
	}
}
