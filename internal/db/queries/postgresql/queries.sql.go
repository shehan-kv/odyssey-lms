package postgresql

import (
	"context"

	"odyssey.lms/internal/db/models"
	"odyssey.lms/internal/db/params"
)

func (q *Queries) FindUserWithPasswordByEmail(ctx context.Context, email string) (models.User, error) {

	const query = `SELECT * FROM users WHERE email = $1`
	row := q.db.QueryRowContext(ctx, query, email)

	var user models.User

	err := row.Scan(
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
		&user.Bio,
	)

	return user, err

}

func (q *Queries) CreateUser(ctx context.Context, arg params.CreateUser) (models.User, error) {

	const query = `INSERT INTO users (first_name, last_name, email, password, avatar_name, bio)
	VALUES ($1,$2,$3,$4,$5,$6)
	RETURNING id, first_name, last_name, email, avatar_name, created_at, bio
	`

	row := q.db.QueryRowContext(ctx, query,
		arg.FirstName,
		arg.LastName,
		arg.Email,
		arg.Password,
		arg.Bio,
	)

	var user models.User

	err := row.Scan(
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.CreatedAt,
		&user.Bio,
	)

	return user, err
}

func (q *Queries) CountUsersByRole(ctx context.Context, role string) (int64, error) {

	const query = `SELECT count(*) FROM user_roles
	JOIN roles ON roles.id = user_roles.role_id
	WHERE roles.name = $1
	`
	row := q.db.QueryRowContext(ctx, query, role)

	var userCount int64
	err := row.Scan(&userCount)

	return userCount, err
}

func (q *Queries) DeleteUserById(ctx context.Context, id int64) error {
	const query = "DELETE FROM users WHERE id = $1"
	_, err := q.db.ExecContext(ctx, query, id)

	return err
}

func (q *Queries) CreateRole(ctx context.Context, name string) (models.Role, error) {

	const query = "INSERT INTO roles (name) VALUES ($1) RETURNING id, name"
	row := q.db.QueryRowContext(ctx, query, name)

	var role models.Role
	err := row.Scan(&role.ID, &role.Name)

	return role, err
}

func (q *Queries) CountRoles(ctx context.Context) (int64, error) {

	const query = "SELECT count(*) FROM roles"
	row := q.db.QueryRowContext(ctx, query)

	var count int64
	err := row.Scan(&count)

	return count, err
}

func (q *Queries) AssignUserRole(ctx context.Context, arg params.AssignUserRole) (models.UserRole, error) {

	const query = `INSERT INTO user_roles (user_id, role_id)
	VALUES ($1, (SELECT id FROM roles WHERE name = $2))
	RETURNING user_id, role_id`

	row := q.db.QueryRowContext(ctx, query, arg.UserID, arg.RoleName)

	var userRole models.UserRole
	err := row.Scan(
		&userRole.UserID,
		&userRole.RoleID,
	)

	return userRole, err
}
