package dto

import (
	"errors"
	"net/mail"
)

type UserCreateRequest struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Role      string `json:"role"`
	IsActive  bool   `json:"isActive"`
	Bio       string `json:"bio"`
}

// This function DOES NOT validate the role
// because it requires a db query.
// Validate roles in the service layer
func (r *UserCreateRequest) Validate() error {

	if r.FirstName == "" {
		return errors.New("first name is required")
	}

	if r.LastName == "" {
		return errors.New("last name is required")
	}

	_, err := mail.ParseAddress(r.Email)
	if err != nil {
		return errors.New("email is invalid")
	}

	if r.Password == "" {
		return errors.New("password is required")
	}

	return nil
}
