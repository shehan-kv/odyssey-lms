package service

import (
	"context"
	"database/sql"
	"errors"

	"odyssey.lms/internal/auth"
	"odyssey.lms/internal/db"
	"odyssey.lms/internal/db/params"
	queryParams "odyssey.lms/internal/dto/params"
	dto "odyssey.lms/internal/dto/user"
	"odyssey.lms/internal/middleware"
)

func GetUsers(ctx context.Context, args queryParams.UserQueryParams) (dto.UsersResponse, error) {

	var resp dto.UsersResponse
	users, err := db.QUERY.GetUsers(ctx, args)
	if err != nil {
		return resp, err
	}

	userCount, err := db.QUERY.CountUsers(ctx, args)
	if err != nil {
		return resp, err
	}

	resp.TotalCount = userCount
	resp.Users = users
	return resp, err
}

func GetUserSelf(ctx context.Context) (dto.UserResponse, error) {
	var userRsp dto.UserResponse

	userId, ok := ctx.Value(middleware.USER_ID).(int64)
	if !ok {
		return userRsp, errors.New("could not get user-id from context")
	}

	user, err := db.QUERY.FindUserById(ctx, userId)
	if err != nil {
		return userRsp, err
	}

	role, err := db.QUERY.FindRoleById(ctx, user.Role)

	userRsp.FirstName = user.FirstName
	userRsp.LastName = user.LastName
	if user.Bio.Valid {
		userRsp.Bio = user.Bio.String
	}
	userRsp.Email = user.Email
	userRsp.IsActive = user.IsActive
	userRsp.Role = role.Name

	return userRsp, nil
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

var ErrLastAdminDeletion = errors.New("last admin account cannot be deleted")

func DeleteUser(ctx context.Context, userId int64) error {

	existingUser, err := db.QUERY.FindUserById(ctx, userId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return ErrUserNotFound
		}
		return err
	}

	role, err := db.QUERY.FindRoleById(ctx, existingUser.Role)
	if err != nil {
		return err
	}

	if role.Name == "administrator" {
		adminCount, err := db.QUERY.CountUsersByRole(ctx, "administrator")
		if err != nil {
			return err
		}

		if adminCount == 1 {
			return ErrLastAdminDeletion
		}
	}

	err = db.QUERY.DeleteUserById(ctx, int64(userId))

	_ = db.QUERY.CreateEvent(ctx, params.CreateEvent{
		Type:        "user",
		Severity:    "info",
		Description: "User account Deleted: " + existingUser.FirstName + " " + existingUser.LastName,
	})
	return err
}

func ActivateUser(ctx context.Context, userId int64) error {
	user, err := db.QUERY.FindUserById(ctx, userId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return ErrUserNotFound
		}
		return err
	}

	err = db.QUERY.SetUserIsActive(ctx, int64(userId), true)

	_ = db.QUERY.CreateEvent(ctx, params.CreateEvent{
		Type:        "user",
		Severity:    "info",
		Description: "User account activated: " + user.FirstName + " " + user.LastName,
	})

	return err
}

func DeactivateUser(ctx context.Context, userId int64) error {
	existingUser, err := db.QUERY.FindUserById(ctx, userId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return ErrUserNotFound
		}
		return err
	}

	role, err := db.QUERY.FindRoleById(ctx, existingUser.Role)
	if err != nil {
		return err
	}

	if role.Name == "administrator" {
		adminCount, err := db.QUERY.CountUsersByRole(ctx, "administrator")
		if err != nil {
			return err
		}

		if adminCount == 1 {
			return ErrLastAdminDeletion
		}
	}

	err = db.QUERY.SetUserIsActive(ctx, int64(userId), false)

	_ = db.QUERY.CreateEvent(ctx, params.CreateEvent{
		Type:        "user",
		Severity:    "info",
		Description: "User account deactivated: " + existingUser.FirstName + " " + existingUser.LastName,
	})

	return err
}
