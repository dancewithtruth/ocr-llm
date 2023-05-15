package middleware

import (
	"net/http"

	"github.com/Wave-95/pgserver/internal/apiresponse"
	"github.com/Wave-95/pgserver/internal/resource/session"
	"github.com/Wave-95/pgserver/pkg/logger"
	"github.com/Wave-95/pgserver/pkg/validator"
)

type CreateSessionRequest struct {
	IPAddress string `validate:"required,uuid4"`
}

func (r CreateSessionRequest) Validate(v validator.Validate) error {
	return v.Struct(r)
}

const CookieNameSession = "dataextract_sessionID"

func Session(service session.Service) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			l := logger.FromContext(r.Context())
			sessionCookie, err := r.Cookie(CookieNameSession)

			//Generate new cookie if none found
			if err == http.ErrNoCookie {
				// Store new cookie
				ipAddress := r.RemoteAddr
				session, err := service.CreateSession(ipAddress)
				if err != nil {
					l.Errorf("Issue storing new session: %v", err)
					apiresponse.RespondError(w, 500, apiresponse.ErrInternalServer)
				}
				sessionCookie = &http.Cookie{
					Name:  CookieNameSession,
					Value: session.Id.String(),
				}
			}

			// Set existing or new cookie to response
			http.SetCookie(w, sessionCookie)
			next.ServeHTTP(w, r)
		})
	}

}
