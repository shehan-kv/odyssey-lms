package middleware

import (
	"context"
	"net/http"
	"strconv"

	"odyssey.lms/internal/auth"
	"odyssey.lms/internal/db"
)

func Admin(next http.Handler) http.Handler {
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

		userId, err := strconv.ParseInt(subject, 10, 64)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		user, err := db.QUERY.FindUserById(r.Context(), userId)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		role, err := db.QUERY.FindRoleById(r.Context(), user.Role)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		if role.Name != "administrator" {
			w.WriteHeader(http.StatusForbidden)
			return
		}

		ctx := context.WithValue(r.Context(), USER_ID, userId)
		req := r.WithContext(ctx)
		next.ServeHTTP(w, req)
	})
}
