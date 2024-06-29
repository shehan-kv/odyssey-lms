package dto

import "errors"

type UserSelfUpdateRequest struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Bio       string `json:"bio"`
}

func (u *UserSelfUpdateRequest) Validate() error {
	if u.FirstName == "" {
		return errors.New("first name is invalid")
	}

	if u.LastName == "" {
		return errors.New("last name is invalid")
	}

	return nil
}
