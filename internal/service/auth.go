package service

import (
	"context"
	"database/sql"
	"errors"

	"odyssey.lms/internal/auth"
	"odyssey.lms/internal/db"
	dto "odyssey.lms/internal/dto/auth"
)

var ErrInvalidPassword = errors.New("invalid password")
var ErrUserNotFound = errors.New("user not found")

func SignIn(ctx context.Context, request dto.SignInRequest) (string, error) {

	existingUser, err := db.QUERY.FindUserWithPasswordByEmail(ctx, request.Email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", ErrUserNotFound
		}
		return "", err
	}

	isPasswordCorrect := auth.CompareHashAndPassword(existingUser.Password, request.Password)
	if !isPasswordCorrect {
		return "", ErrInvalidPassword
	}

	jwtToken, err := auth.NewJWTToken(existingUser.ID)
	if err != nil {
		return "", err
	}

	return jwtToken, nil
}
