package sqlite

import (
	"context"
	"database/sql"
	"strconv"
	"strings"
	"time"

	"odyssey.lms/internal/db/models"
	"odyssey.lms/internal/db/params"
	dto "odyssey.lms/internal/dto/user"
)

func (q *Queries) FindUserWithPasswordByEmail(ctx context.Context, email string) (models.User, error) {

	const query = `SELECT id, first_name, last_name, email, password, created_at, bio FROM users WHERE email = ?`
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

func (q *Queries) GetUsers(ctx context.Context, arg params.UserQueryParams) ([]dto.UserResponse, error) {
	var sb strings.Builder
	sb.WriteString("SELECT u.id, u.first_name, u.last_name, u.email, u.created_at, u.last_login, u.is_active, r.name FROM users u")
	sb.WriteString(" JOIN roles r ON u.role = r.id")
	if arg.Search != "" || arg.Role != "" {
		sb.WriteString(" WHERE")
		if arg.Search != "" {
			sb.WriteString(" first_name LIKE '%")
			sb.WriteString(arg.Search)
			sb.WriteString("%'")
			sb.WriteString(" OR last_name LIKE '%")
			sb.WriteString(arg.Search)
			sb.WriteString("%'")
		}

		if arg.Search != "" && arg.Role != "" {
			sb.WriteString(" AND")
		}

		if arg.Role != "" {
			sb.WriteString(" r.name = '")
			sb.WriteString(arg.Role)
			sb.WriteString("'")
		}
	}

	if arg.Page > 0 {
		offset := (arg.Page - 1) * arg.Limit
		sb.WriteString(" LIMIT ")
		sb.WriteString(strconv.Itoa(arg.Limit))
		sb.WriteString(" ")
		sb.WriteString(" OFFSET ")
		sb.WriteString(strconv.Itoa(offset))
	}

	rows, err := q.db.QueryContext(ctx, sb.String())
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var users []dto.UserResponse
	for rows.Next() {
		var user dto.UserResponse
		var createdAt sql.NullTime
		var lastLogin sql.NullTime
		err := rows.Scan(
			&user.ID,
			&user.FirstName,
			&user.LastName,
			&user.Email,
			&createdAt,
			&lastLogin,
			&user.IsActive,
			&user.Role,
		)

		if err != nil {
			return nil, err
		}

		if createdAt.Valid {
			user.CreatedAt = createdAt.Time.Format(time.RFC3339)
		}
		if lastLogin.Valid {
			user.LastLogin = lastLogin.Time.Format(time.RFC3339)
		}
		users = append(users, user)
	}

	return users, nil
}

func (q *Queries) FindUserById(ctx context.Context, userId int) (models.User, error) {
	const query = `SELECT id, first_name, last_name, email, created_at, last_login, is_active, role FROM users
	WHERE id = ?
	`
	row := q.db.QueryRowContext(ctx, query, userId)

	var user models.User
	err := row.Scan(
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.CreatedAt,
		&user.LastLogin,
		&user.IsActive,
		&user.Role,
	)

	return user, err
}

func (q *Queries) CreateUser(ctx context.Context, arg params.CreateUser) (int64, error) {

	const query = `INSERT INTO users (first_name, last_name, email, password, is_active, bio)
	VALUES (?,?,?,?,?,?)`

	result, err := q.db.ExecContext(ctx, query,
		arg.FirstName,
		arg.LastName,
		arg.Email,
		arg.Password,
		arg.IsActive,
		arg.Bio,
	)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	return id, err
}

func (q *Queries) CountUsersByRole(ctx context.Context, role string) (int64, error) {

	const query = `SELECT count(*) FROM users
	JOIN roles ON roles.id = users.role
	WHERE roles.name = ?
	`
	row := q.db.QueryRowContext(ctx, query, role)

	var userCount int64
	err := row.Scan(&userCount)

	return userCount, err
}

func (q *Queries) CountUsers(ctx context.Context, arg params.UserQueryParams) (int64, error) {

	var sb strings.Builder
	sb.WriteString("SELECT count(*) FROM users u")
	sb.WriteString(" JOIN roles r ON u.role = r.id")
	if arg.Search != "" || arg.Role != "" {
		sb.WriteString(" WHERE")
		if arg.Search != "" {
			sb.WriteString(" first_name LIKE '%")
			sb.WriteString(arg.Search)
			sb.WriteString("%'")
			sb.WriteString(" OR last_name LIKE '%")
			sb.WriteString(arg.Search)
			sb.WriteString("%'")
		}

		if arg.Search != "" && arg.Role != "" {
			sb.WriteString(" AND")
		}

		if arg.Role != "" {
			sb.WriteString(" r.name = '")
			sb.WriteString(arg.Role)
			sb.WriteString("'")
		}
	}

	row := q.db.QueryRowContext(ctx, sb.String())

	var userCount int64
	err := row.Scan(&userCount)

	return userCount, err
}

func (q *Queries) DeleteUserById(ctx context.Context, id int64) error {
	const query = "DELETE FROM users WHERE id = ?"
	_, err := q.db.ExecContext(ctx, query, id)

	return err
}

func (q *Queries) SetUserIsActive(ctx context.Context, id int64, status bool) error {
	const query = "UPDATE users SET is_active = ? WHERE id = ?"
	_, err := q.db.ExecContext(ctx, query, status, id)

	return err
}

func (q *Queries) CreateRole(ctx context.Context, name string) (models.Role, error) {

	const query = "INSERT INTO roles (name) VALUES (?) RETURNING id, name"
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

func (q *Queries) GetRoles(ctx context.Context) ([]models.Role, error) {

	var roles []models.Role

	const query = "SELECT * FROM roles"
	rows, err := q.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var role models.Role
		err := rows.Scan(&role.ID, &role.Name, &role.IsDefault)
		if err != nil {
			return nil, err
		}

		roles = append(roles, role)
	}

	return roles, nil
}

func (q *Queries) FindRoleById(ctx context.Context, roleId int64) (models.Role, error) {
	const query = "SELECT * FROM roles WHERE id = ?"
	row := q.db.QueryRowContext(ctx, query, roleId)

	var role models.Role
	err := row.Scan(&role.ID, &role.Name, &role.IsDefault)

	return role, err
}

func (q *Queries) AssignUserRole(ctx context.Context, arg params.AssignUserRole) error {

	const query = "UPDATE users SET role = (SELECT id FROM roles WHERE name = ?) WHERE id = ?"
	_, err := q.db.ExecContext(ctx, query, arg.RoleName, arg.UserID)

	return err
}
