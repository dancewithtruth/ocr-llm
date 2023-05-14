package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Wave-95/pgserver/internal/resource/session"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestSession(t *testing.T) {
	t.Run("sets a session cookie to response and stores it", func(t *testing.T) {
		// Set up server
		mux := http.NewServeMux()

		// Set up middleware with mock session repository
		mockSessionRepository := session.NewMockRepository()
		sessionService := session.NewService(&mockSessionRepository)
		sessionMiddleware := Session(sessionService)
		handler := sessionMiddleware(mux)

		// Set up request/recorder and send mock request
		request, _ := http.NewRequest(http.MethodGet, "/", nil)
		response := httptest.NewRecorder()
		handler.ServeHTTP(response, request)

		// Check if response cookie has session
		cookies := response.Result().Cookies()
		sessionCookie := getSessionCookie(cookies)
		assert.NotNil(t, sessionCookie, "expected session cookie but none found")
		sessionId, err := uuid.Parse(sessionCookie.Value)
		if err != nil {
			t.Fatal("Issue parsing session cookie value into UUID")
		}

		// Check if session was stored into database
		gotSession, err := sessionService.GetSession(sessionId)
		assert.Nil(t, err)
		assert.Equal(t, gotSession.Id, sessionId, "sessionId stored does not match sessionId in cookie: got %q, wanted %q", gotSession.Id, sessionId)
	})
}

func getSessionCookie(cookies []*http.Cookie) *http.Cookie {
	for _, cookie := range cookies {
		if cookie.Name == CookieNameSession && cookie.Value != "" {
			return cookie
		}
	}
	return nil
}
