package handler

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	dto "odyssey.lms/internal/dto/auth"
	"odyssey.lms/internal/service"
)

func SignIn(w http.ResponseWriter, r *http.Request) {

	var signInData dto.SignInRequest

	err := json.NewDecoder(r.Body).Decode(&signInData)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	err = signInData.Validate()
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusUnauthorized)
	}

	jwtToken, err := service.SignIn(r.Context(), signInData)
	if err != nil {
		if errors.Is(err, service.ErrInvalidPassword) {
			log.Println(err)
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		if errors.Is(err, service.ErrUserNotFound) {
			log.Println(err)
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var maxAge = 0
	if signInData.RememberMe {
		maxAge = 60 * 60 * 24 * 2 //2 days
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "auth-token",
		Value:    jwtToken,
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteStrictMode,
		MaxAge:   maxAge,
	})
	w.WriteHeader(http.StatusOK)
}
