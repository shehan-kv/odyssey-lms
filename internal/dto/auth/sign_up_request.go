package dto

import (
	"errors"
	"net/mail"
)

type SignUpRequest struct {
	FirstName       string `json:"firstName"`
	LastName        string `json:"lastName"`
	Email           string `json:"email"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirmPassword"`
}

func (s *SignUpRequest) Validate() error {
	if s.FirstName == "" {
		return errors.New("first name is invalid")
	}

	if s.LastName == "" {
		return errors.New("last name is invalid")
	}

	_, err := mail.ParseAddress(s.Email)
	if err != nil {
		return errors.New("email is invalid")
	}

	if s.Password == "" {
		return errors.New("password is invalid")
	}

	if s.ConfirmPassword == "" {
		return errors.New("confirm password is invalid")
	}

	if s.Password != s.ConfirmPassword {
		return errors.New("passwords don't match")
	}

	return nil
}
