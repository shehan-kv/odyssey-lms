package middleware

import (
	"context"
	"net/http"
	"strconv"

	"odyssey.lms/internal/auth"
)

type key string

const USER_ID key = "user-id"

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

		userId, err := strconv.Atoi(subject)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		ctx := context.WithValue(r.Context(), USER_ID, int64(userId))
		req := r.WithContext(ctx)
		next.ServeHTTP(w, req)
	})
}
