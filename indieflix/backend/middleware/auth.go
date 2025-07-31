package middleware

import (
	"log"
	"net/http"

	"github.com/alexedwards/scs/v2"
)

func getSession(session *scs.SessionManager, r *http.Request, token string) bool {
	if session.Exists(r.Context(), token) {
		return true
	} else {
		return false
	}
}
func Authenticate(sessionManager *scs.SessionManager) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			cookie, err := r.Cookie("session_token")
			log.Printf("cookie: %v", cookie)
			if err != nil {
				log.Printf("Error getting cookie: %v", err)
				http.Error(w, http.StatusText(500), 500)
				return
			}
			ok := getSession(sessionManager, r, cookie.Value)
			if !ok {
				log.Printf("Invalid session: %v", cookie.Value)
				http.Error(w, http.StatusText(500), 500)
				return
			}
			next.ServeHTTP(w, r)
		})
	}

}
