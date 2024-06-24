package mysql

import (
	"context"
	"database/sql"
	"strconv"
	"strings"
	"time"

	"odyssey.lms/internal/db/models"
	"odyssey.lms/internal/db/params"
	evntDto "odyssey.lms/internal/dto/event"
	queryParams "odyssey.lms/internal/dto/params"
	usrDto "odyssey.lms/internal/dto/user"
)

func (q *Queries) FindUserWithPasswordByEmail(ctx context.Context, email string) (models.User, error) {

	const query = `SELECT * FROM users WHERE email = ?`
	row := q.db.QueryRowContext(ctx, query, email)

	var user models.User

	err := row.Scan(
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.Password,
		&user.Role,
		&user.CreatedAt,
		&user.LastLogin,
		&user.IsActive,
		&user.Bio,
	)

	return user, err

}

func (q *Queries) GetUsers(ctx context.Context, arg queryParams.UserQueryParams) ([]usrDto.UserResponse, error) {
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

	var users = make([]usrDto.UserResponse, 0)

	rows, err := q.db.QueryContext(ctx, sb.String())
	if err != nil {
		return users, err
	}

	defer rows.Close()

	for rows.Next() {
		var user usrDto.UserResponse
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
			return users, err
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

func (q *Queries) FindUserById(ctx context.Context, userId int64) (models.User, error) {
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

	query := `INSERT INTO users (first_name, last_name, email, password, is_active, bio)
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

	query := `SELECT count(*) FROM users
	JOIN roles ON roles.id = users.role
	WHERE roles.name = ?
	`
	row := q.db.QueryRowContext(ctx, query, role)

	var userCount int64
	err := row.Scan(&userCount)

	return userCount, err
}

func (q *Queries) CountUsers(ctx context.Context, arg queryParams.UserQueryParams) (int64, error) {

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
	query := "DELETE FROM users WHERE id = ?"
	_, err := q.db.ExecContext(ctx, query, id)

	return err
}

func (q *Queries) SetUserIsActive(ctx context.Context, id int64, status bool) error {
	const query = "UPDATE users SET is_active = ? WHERE id = ?"
	_, err := q.db.ExecContext(ctx, query, status, id)

	return err
}

func (q *Queries) CreateRole(ctx context.Context, name string) (models.Role, error) {

	query := "INSERT INTO roles (name) VALUES (?)"
	result, err := q.db.ExecContext(ctx, query, name)

	var role models.Role
	if err != nil {
		return role, err
	}

	insertedId, err := result.LastInsertId()
	if err != nil {
		return role, err
	}

	newRoleQuery := "SELECT id, name FROM roles WHERE id = ?"
	row := q.db.QueryRowContext(ctx, newRoleQuery, insertedId)

	err = row.Scan(&role.ID, &role.Name)

	return role, err
}

func (q *Queries) CountRoles(ctx context.Context) (int64, error) {

	query := "SELECT count(*) FROM roles"
	row := q.db.QueryRowContext(ctx, query)

	var count int64
	err := row.Scan(&count)

	return count, err
}

func (q *Queries) GetRoles(ctx context.Context) ([]models.Role, error) {

	var roles = make([]models.Role, 0)

	const query = "SELECT * FROM roles"
	rows, err := q.db.QueryContext(ctx, query)
	if err != nil {
		return roles, err
	}

	defer rows.Close()

	for rows.Next() {
		var role models.Role
		err := rows.Scan(&role.ID, &role.Name, &role.IsDefault)
		if err != nil {
			return roles, err
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

func (q *Queries) CreateEvent(ctx context.Context, arg params.CreateEvent) error {
	const query = "INSERT INTO events (type, description, severity) VALUES (?,?,?)"
	_, err := q.db.ExecContext(ctx, query, arg.Type, arg.Description, arg.Severity)

	return err
}

func (q *Queries) GetEvents(ctx context.Context, arg queryParams.EventQueryParams) ([]evntDto.EventResponse, error) {
	var sb strings.Builder
	sb.WriteString("SELECT id, created_at, type, description, severity FROM events")
	if arg.Search != "" || arg.Type != "" || arg.Severity != "" {
		sb.WriteString(" WHERE")
		if arg.Search != "" {
			sb.WriteString(" description LIKE '%")
			sb.WriteString(arg.Search)
			sb.WriteString("%'")
		}

		var stmts []string
		if arg.Type != "" {
			stmts = append(stmts, " type = '"+arg.Type+"'")
		}
		if arg.Severity != "" {
			stmts = append(stmts, " severity = '"+arg.Severity+"'")
		}
		if arg.Search != "" && len(stmts) > 0 {
			sb.WriteString(" AND ")
		}
		sb.WriteString(strings.Join(stmts, " AND "))
	}

	sb.WriteString(" ORDER BY created_at DESC")
	if arg.Page > 0 {
		offset := (arg.Page - 1) * arg.Limit
		sb.WriteString(" LIMIT ")
		sb.WriteString(strconv.Itoa(arg.Limit))
		sb.WriteString(" ")
		sb.WriteString(" OFFSET ")
		sb.WriteString(strconv.Itoa(offset))
	}

	var events = make([]evntDto.EventResponse, 0)

	rows, err := q.db.QueryContext(ctx, sb.String())
	if err != nil {
		return events, err
	}

	defer rows.Close()

	for rows.Next() {
		var event evntDto.EventResponse
		var createdAt sql.NullTime
		err := rows.Scan(
			&event.Id,
			&createdAt,
			&event.Type,
			&event.Description,
			&event.Severity,
		)

		if err != nil {
			return events, err
		}
		event.CreatedAt = createdAt.Time.Format(time.RFC3339)
		events = append(events, event)
	}

	return events, nil
}

func (q *Queries) CountEvents(ctx context.Context, arg queryParams.EventQueryParams) (int64, error) {
	var sb strings.Builder
	sb.WriteString("SELECT count(*) FROM events")
	if arg.Search != "" || arg.Type != "" || arg.Severity != "" {
		sb.WriteString(" WHERE")
		if arg.Search != "" {
			sb.WriteString(" description LIKE '%")
			sb.WriteString(arg.Search)
			sb.WriteString("%'")
		}

		var stmts []string
		if arg.Type != "" {
			stmts = append(stmts, " type = '"+arg.Type+"'")
		}
		if arg.Severity != "" {
			stmts = append(stmts, " severity = '"+arg.Severity+"'")
		}
		if arg.Search != "" && len(stmts) > 0 {
			sb.WriteString(" AND ")
		}
		sb.WriteString(strings.Join(stmts, " AND "))
	}

	row := q.db.QueryRowContext(ctx, sb.String())

	var count int64
	err := row.Scan(&count)

	return count, err
}

func (q *Queries) CreateTicket(ctx context.Context, arg params.CreateTicket) error {
	const query = "INSERT INTO tickets(subject, description, user_id, type, status) VALUES(?, ?, ?, ?, ?)"
	_, err := q.db.ExecContext(ctx, query, arg.Subject, arg.Description, arg.UserId, arg.Type, arg.Status)

	return err
}
