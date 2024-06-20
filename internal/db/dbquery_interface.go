package db

import (
	"context"

	"odyssey.lms/internal/db/models"
	"odyssey.lms/internal/db/params"
)

type DBQuery interface {
	FindUserWithPasswordByEmail(ctx context.Context, email string) (models.User, error)
	GetUsers(ctx context.Context, params params.UserQueryParams) ([]models.User, error)
	CreateUser(ctx context.Context, arg params.CreateUser) (models.User, error)
	CountUsersByRole(ctx context.Context, role string) (int64, error)
	CountUsers(ctx context.Context) (int64, error)
	DeleteUserById(ctx context.Context, id int64) error
	CreateRole(ctx context.Context, name string) (models.Role, error)
	CountRoles(ctx context.Context) (int64, error)
	AssignUserRole(ctx context.Context, arg params.AssignUserRole) error
}
