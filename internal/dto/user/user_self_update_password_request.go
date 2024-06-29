package dto

import "errors"

type UserSelfUpdatePasswordRequest struct {
	CurrentPassword    string `json:"currentPassword"`
	NewPassword        string `json:"newPassword"`
	ConfirmNewPassword string `json:"confirmNewPassword"`
}

func (u *UserSelfUpdatePasswordRequest) Validate() error {
	if u.CurrentPassword == "" {
		return errors.New("current password is invalid")
	}

	if u.NewPassword == "" {
		return errors.New("new password is invalid")
	}

	if u.ConfirmNewPassword == "" {
		return errors.New("confirm new password is invalid")
	}

	if u.NewPassword != u.ConfirmNewPassword {
		return errors.New("passwords don't match")
	}

	return nil
}
