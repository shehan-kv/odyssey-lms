package postgresql

import (
	"context"
	"log"
	"strconv"
	"strings"

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

func (q *Queries) GetUsers(ctx context.Context, params params.UserQueryParams) ([]models.User, error) {
	var sb strings.Builder
	sb.WriteString("SELECT u.id, u.first_name, u.last_name, u.email, u.created_at, u.last_login, u.is_active, r.name FROM users u")
	sb.WriteString("JOIN roles r ON u.role = r.id")
	if params.Search != "" || params.Role != "" {
		sb.WriteString(" WHERE")
		if params.Search != "" {
			sb.WriteString(" first_name LIKE '")
			sb.WriteString(params.Search)
			sb.WriteString("'")
			sb.WriteString(" OR last_name LIKE '")
			sb.WriteString(params.Search)
			sb.WriteString("'")
		}

		if params.Search != "" && params.Role != "" {
			sb.WriteString(" AND")
		}

		if params.Role != "" {
			sb.WriteString(" role = '")
			sb.WriteString(params.Role)
			sb.WriteString("'")
		}
	}

	if params.Page > 0 {
		offset := (params.Page - 1) * params.Limit
		sb.WriteString(" LIMIT ")
		sb.WriteString(strconv.Itoa(params.Limit))
		sb.WriteString(" ")
		sb.WriteString(" OFFSET ")
		sb.WriteString(strconv.Itoa(offset))

	}

	log.Println(sb.String())
	rows, err := q.db.QueryContext(ctx, sb.String())
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		err := rows.Scan(
			&user.ID,
			&user.FirstName,
			&user.LastName,
			&user.Email,
			&user.CreatedAt,
			&user.LastLogin,
			&user.IsActive,
			&user.Role,
		)

		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
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

func (q *Queries) CountUsers(ctx context.Context) (int64, error) {

	const query = "SELECT count(*) FROM users"
	row := q.db.QueryRowContext(ctx, query)

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

func (q *Queries) AssignUserRole(ctx context.Context, arg params.AssignUserRole) error {

	const query = "UPDATE user SET role = (SELECT id FROM roles WHERE name = $1) WHERE id = $2"
	_, err := q.db.ExecContext(ctx, query, arg.RoleName, arg.UserID)

	return err
}
