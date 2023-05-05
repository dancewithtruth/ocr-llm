package logMiddleware

import (
	"net/http"
	"time"

	"github.com/Wave-95/pgserver/middleware/requestid"
	"github.com/Wave-95/pgserver/pkg/logger"
	"github.com/go-chi/chi/middleware"
)

type responseWriterWithStatusCode struct {
	http.ResponseWriter
	statusCode   int
	bytesWritten int
}

func (rw *responseWriterWithStatusCode) WriteHeader(statusCode int) {
	rw.statusCode = statusCode
	rw.ResponseWriter.WriteHeader(statusCode)
}

func (rw *responseWriterWithStatusCode) Write(buf []byte) (int, error) {
	if n, err := rw.ResponseWriter.Write(buf); err != nil {
		return 0, err
	} else {
		rw.bytesWritten = n
		return n, err
	}
}

// Middleware injects a logger and builds an http handler. It finds the requestId and correlationId through the
// request context and add its to the logger. It also starts a timer for the request and logs it out.
func Middleware(logger logger.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			//Log request ids

			//Log request time using defer
			// generate an access log message
			ctx := r.Context()
			requestID, ok := ctx.Value(requestid.RequestIdKey).(string)

			if ok {
				logger.Info(requestID)
			}

			// Wrap the ResponseWriter to capture the status code
			wrappedWriter := &responseWriterWithStatusCode{
				ResponseWriter: w,
				statusCode:     http.StatusOK,
			}
			next.ServeHTTP(wrappedWriter, r)
			defer func() {
				logger.
					WithoutCaller().
					With("duration", time.Since(start).Milliseconds(), "status", wrappedWriter.statusCode, "bytes written", wrappedWriter.bytesWritten).
					Infof("%s %s", r.Method, r.URL.Path)
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
