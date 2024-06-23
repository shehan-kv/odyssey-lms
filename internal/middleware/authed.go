package middleware

import (
	"net/http"

	"odyssey.lms/internal/auth"
)

func Authed(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		cookie, err := r.Cookie("auth-token")
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		err = cookie.Valid()
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		subject, err := auth.VerifyJWTToken(cookie.Value)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		r.Header.Set("subject", subject)
		next.ServeHTTP(w, r)
	})
}
