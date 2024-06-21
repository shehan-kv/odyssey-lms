package service

import (
	"context"
	"database/sql"
	"errors"

	"odyssey.lms/internal/auth"
	"odyssey.lms/internal/db"
	"odyssey.lms/internal/db/params"
	dto "odyssey.lms/internal/dto/user"
)

func GetUsers(ctx context.Context, page int, limit int, search string, role string) (dto.UsersResponse, error) {

	var resp dto.UsersResponse
	params := params.UserQueryParams{
		Page:   page,
		Limit:  limit,
		Search: search,
		Role:   role,
	}
	users, err := db.QUERY.GetUsers(ctx, params)
	if err != nil {
		return resp, err
	}

	userCount, err := db.QUERY.CountUsers(ctx, params)
	if err != nil {
		return resp, err
	}

	resp.TotalCount = userCount
	resp.Users = users
	return resp, err
}

var ErrInvalidRole = errors.New("role is invalid")

func CreateUser(ctx context.Context, createReq dto.UserCreateRequest) error {
	userParams := params.CreateUser{
		FirstName: createReq.FirstName,
		LastName:  createReq.LastName,
		Email:     createReq.Email,
		IsActive:  createReq.IsActive,
	}
	hash, err := auth.HashPassword(createReq.Password)
	if err != nil {
		return err
	}

	userParams.Password = hash

	if createReq.Bio != "" {
		userParams.Bio = sql.NullString{String: createReq.Bio, Valid: true}
	}

	if createReq.Role == "" {
		return ErrInvalidRole
	}

	roles, err := db.QUERY.GetRoles(ctx)
	if err != nil {
		return err
	}

	isRoleValid := false
	for _, role := range roles {
		if role.Name == createReq.Role {
			isRoleValid = true
		}
	}

	if !isRoleValid {
		return ErrInvalidRole
	}

	userId, err := db.QUERY.CreateUser(ctx, userParams)
	if err != nil {
		return err
	}

	assignRoleParams := params.AssignUserRole{UserID: userId, RoleName: createReq.Role}
	err = db.QUERY.AssignUserRole(ctx, assignRoleParams)
	if err != nil {
		return err
	}

	return nil
}
