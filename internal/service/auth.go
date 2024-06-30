package service

import (
	"context"
	"database/sql"
	"errors"

	"odyssey.lms/internal/auth"
	"odyssey.lms/internal/db"
	"odyssey.lms/internal/db/params"
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

func SignUp(ctx context.Context, request dto.SignUpRequest) error {

	_, err := db.QUERY.FindUserWithPasswordByEmail(ctx, request.Email)
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			return err
		}
	}

	if err == nil {
		return errors.New("user already exists")
	}

	hashedPassword, err := auth.HashPassword(request.Password)
	if err != nil {
		return err
	}

	userId, err := db.QUERY.CreateUser(ctx, params.CreateUser{
		FirstName: request.FirstName,
		LastName:  request.LastName,
		Email:     request.Email,
		IsActive:  true,
		Password:  hashedPassword,
	})
	if err != nil {
		return err
	}

	err = db.QUERY.AssignUserRole(ctx, params.AssignUserRole{UserID: userId, RoleName: "student"})
	if err != nil {
		return err
	}

	_ = db.QUERY.CreateEvent(ctx, params.CreateEvent{
		Type:        "user",
		Severity:    "info",
		Description: "User signed up: " + request.FirstName + " " + request.LastName,
	})

	return nil
}
