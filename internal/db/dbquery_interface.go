package db

import (
	"context"

	"odyssey.lms/internal/db/models"
	"odyssey.lms/internal/db/params"
	dto "odyssey.lms/internal/dto/user"
)

type DBQuery interface {
	FindUserWithPasswordByEmail(ctx context.Context, email string) (models.User, error)
	GetUsers(ctx context.Context, params params.UserQueryParams) ([]dto.UserResponse, error)
	FindUserById(ctx context.Context, userId int) (models.User, error)
	CreateUser(ctx context.Context, arg params.CreateUser) (int64, error)
	CountUsersByRole(ctx context.Context, role string) (int64, error)
	CountUsers(ctx context.Context, arg params.UserQueryParams) (int64, error)
	DeleteUserById(ctx context.Context, id int64) error
	SetUserIsActive(ctx context.Context, id int64, status bool) error
	CreateRole(ctx context.Context, name string) (models.Role, error)
	CountRoles(ctx context.Context) (int64, error)
	GetRoles(ctx context.Context) ([]models.Role, error)
	FindRoleById(ctx context.Context, roleId int64) (models.Role, error)
	AssignUserRole(ctx context.Context, arg params.AssignUserRole) error
}
