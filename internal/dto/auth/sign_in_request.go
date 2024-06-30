package dto

import (
	"errors"
	"net/mail"
)

type SignInRequest struct {
	Email      string `json:"email"`
	Password   string `json:"password"`
	RememberMe bool   `json:"remember_me"`
}

func (r *SignInRequest) Validate() error {

	if len(r.Email) == 0 {
		return errors.New("email is empty")
	}

	if len(r.Password) == 0 {
		return errors.New("password is empty")
	}

	_, err := mail.ParseAddress(r.Email)
	if err != nil {
		return errors.New("invalid email")
	}

	return nil
}
