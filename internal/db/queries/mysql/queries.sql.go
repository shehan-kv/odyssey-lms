package mysql

import (
	"context"
	"database/sql"
	"strconv"
	"strings"
	"time"

	"odyssey.lms/internal/db/models"
	"odyssey.lms/internal/db/params"
	courseDto "odyssey.lms/internal/dto/course"
	evntDto "odyssey.lms/internal/dto/event"
	queryParams "odyssey.lms/internal/dto/params"
	ticketDto "odyssey.lms/internal/dto/ticket"
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

func (q *Queries) GetTickets(ctx context.Context, arg queryParams.TicketQueryParams) ([]ticketDto.TicketResponse, error) {
	var sb strings.Builder
	sb.WriteString("SELECT t.id, t.subject, CONCAT(u.first_name, ' ', u.last_name) AS user,  t.created_at, t.type, t.status FROM tickets t")
	sb.WriteString(" JOIN users u ON t.user_id = u.id")
	if arg.Search != "" || arg.Type != "" || arg.Status != "" {
		sb.WriteString(" WHERE")
		if arg.Search != "" {
			sb.WriteString(" t.subject LIKE '%")
			sb.WriteString(arg.Search)
			sb.WriteString("%'")
		}

		var stmts []string
		if arg.Type != "" {
			stmts = append(stmts, " t.type = '"+arg.Type+"'")
		}
		if arg.Status != "" {
			stmts = append(stmts, " t.status = '"+arg.Status+"'")
		}
		if arg.Search != "" && len(stmts) > 0 {
			sb.WriteString(" AND ")
		}
		sb.WriteString(strings.Join(stmts, " AND "))
	}

	sb.WriteString(" ORDER BY t.created_at DESC")
	if arg.Page > 0 {
		offset := (arg.Page - 1) * arg.Limit
		sb.WriteString(" LIMIT ")
		sb.WriteString(strconv.Itoa(arg.Limit))
		sb.WriteString(" ")
		sb.WriteString(" OFFSET ")
		sb.WriteString(strconv.Itoa(offset))
	}

	var tickets = make([]ticketDto.TicketResponse, 0)

	rows, err := q.db.QueryContext(ctx, sb.String())
	if err != nil {
		return tickets, err
	}

	defer rows.Close()

	for rows.Next() {
		var ticket ticketDto.TicketResponse
		var createdAt sql.NullTime
		err := rows.Scan(
			&ticket.Id,
			&ticket.Subject,
			&ticket.User,
			&createdAt,
			&ticket.Type,
			&ticket.Status,
		)

		if err != nil {
			return tickets, err
		}
		ticket.CreatedAt = createdAt.Time.Format(time.RFC3339)
		tickets = append(tickets, ticket)
	}

	return tickets, nil
}

func (q *Queries) CountTickets(ctx context.Context, arg queryParams.TicketQueryParams) (int64, error) {
	var sb strings.Builder
	sb.WriteString("SELECT count(*) FROM tickets")
	if arg.Search != "" || arg.Type != "" || arg.Status != "" {
		sb.WriteString(" WHERE")
		if arg.Search != "" {
			sb.WriteString(" subject LIKE '%")
			sb.WriteString(arg.Search)
			sb.WriteString("%'")
		}

		var stmts []string
		if arg.Type != "" {
			stmts = append(stmts, " type = '"+arg.Type+"'")
		}
		if arg.Status != "" {
			stmts = append(stmts, " status = '"+arg.Status+"'")
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

func (q *Queries) GetTicketsByUserId(ctx context.Context, userId int64, arg queryParams.TicketQueryParams) ([]ticketDto.TicketResponse, error) {
	var sb strings.Builder
	sb.WriteString("SELECT t.id, t.subject, CONCAT(u.first_name, ' ', u.last_name) AS user,  t.created_at, t.type, t.status FROM tickets t")
	sb.WriteString(" JOIN users u ON t.user_id = u.id")
	sb.WriteString(" WHERE u.id = ?")
	if arg.Search != "" || arg.Type != "" || arg.Status != "" {
		sb.WriteString(" AND")
		if arg.Search != "" {
			sb.WriteString(" t.subject LIKE '%")
			sb.WriteString(arg.Search)
			sb.WriteString("%'")
		}

		var stmts []string
		if arg.Type != "" {
			stmts = append(stmts, " t.type = '"+arg.Type+"'")
		}
		if arg.Status != "" {
			stmts = append(stmts, " t.status = '"+arg.Status+"'")
		}
		if arg.Search != "" && len(stmts) > 0 {
			sb.WriteString(" AND ")
		}
		sb.WriteString(strings.Join(stmts, " AND "))
	}

	sb.WriteString(" ORDER BY t.created_at DESC")
	if arg.Page > 0 {
		offset := (arg.Page - 1) * arg.Limit
		sb.WriteString(" LIMIT ")
		sb.WriteString(strconv.Itoa(arg.Limit))
		sb.WriteString(" ")
		sb.WriteString(" OFFSET ")
		sb.WriteString(strconv.Itoa(offset))
	}

	var tickets = make([]ticketDto.TicketResponse, 0)

	rows, err := q.db.QueryContext(ctx, sb.String(), userId)
	if err != nil {
		return tickets, err
	}

	defer rows.Close()

	for rows.Next() {
		var ticket ticketDto.TicketResponse
		var createdAt sql.NullTime
		err := rows.Scan(
			&ticket.Id,
			&ticket.Subject,
			&ticket.User,
			&createdAt,
			&ticket.Type,
			&ticket.Status,
		)

		if err != nil {
			return tickets, err
		}
		ticket.CreatedAt = createdAt.Time.Format(time.RFC3339)
		tickets = append(tickets, ticket)
	}

	return tickets, nil
}

func (q *Queries) CountTicketsByUserId(ctx context.Context, userId int64, arg queryParams.TicketQueryParams) (int64, error) {
	var sb strings.Builder
	sb.WriteString("SELECT count(*) FROM tickets")
	sb.WriteString(" WHERE user_id = ?")
	if arg.Search != "" || arg.Type != "" || arg.Status != "" {
		sb.WriteString(" AND")
		if arg.Search != "" {
			sb.WriteString(" subject LIKE '%")
			sb.WriteString(arg.Search)
			sb.WriteString("%'")
		}

		var stmts []string
		if arg.Type != "" {
			stmts = append(stmts, " type = '"+arg.Type+"'")
		}
		if arg.Status != "" {
			stmts = append(stmts, " status = '"+arg.Status+"'")
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

func (q *Queries) GetTicketByIdWithUser(ctx context.Context, ticketId int64) (ticketDto.TicketResponse, error) {
	const query = `SELECT t.id, t.subject, CONCAT(u.first_name, ' ', u.last_name) AS user, t.type, t.created_at, t.closed_at, t.status FROM tickets t
	JOIN users u ON u.id = t.user_id
	WHERE t.id = ?
	`

	row := q.db.QueryRowContext(ctx, query, ticketId)

	var ticketRsp ticketDto.TicketResponse
	var createdAt sql.NullTime
	var closedAt sql.NullTime
	err := row.Scan(
		&ticketRsp.Id,
		&ticketRsp.Subject,
		&ticketRsp.Description,
		&ticketRsp.User,
		&ticketRsp.Type,
		&createdAt,
		&closedAt,
		&ticketRsp.Status,
	)
	if err != nil {
		return ticketRsp, err
	}

	ticketRsp.CreatedAt = createdAt.Time.Format(time.RFC3339)

	if closedAt.Valid {
		ticketRsp.ClosedAt = closedAt.Time.Format(time.RFC3339)
	}
	return ticketRsp, nil
}

func (q *Queries) GetTicketMessagesByTicketId(ctx context.Context, ticketId int64) ([]ticketDto.TicketMessageResponse, error) {
	const query = `SELECT m.id, CONCAT(u.first_name, ' ', u.last_name) user, m.created_at, m.content FROM ticket_messages m
	JOIN users u ON u.id = m.user_id
	WHERE m.ticket_id = ?
	ORDER BY m.created_at ASC
	`

	var messages = make([]ticketDto.TicketMessageResponse, 0)

	rows, err := q.db.QueryContext(ctx, query, ticketId)
	if err != nil {
		return messages, err
	}

	defer rows.Close()

	for rows.Next() {
		var message ticketDto.TicketMessageResponse
		var createdAt sql.NullTime
		err := rows.Scan(
			&message.Id,
			&message.User,
			&createdAt,
			&message.Content,
		)

		if err != nil {
			return messages, err
		}
		message.CreatedAt = createdAt.Time.Format(time.RFC3339)
		messages = append(messages, message)
	}

	return messages, nil
}

func (q *Queries) FindTicketById(ctx context.Context, ticketId int64) (models.Ticket, error) {
	const query = "SELECT * FROM tickets WHERE id = ?"

	row := q.db.QueryRowContext(ctx, query, ticketId)

	var ticketRsp models.Ticket
	err := row.Scan(
		&ticketRsp.Id,
		&ticketRsp.Subject,
		&ticketRsp.Description,
		&ticketRsp.UserId,
		&ticketRsp.Type,
		&ticketRsp.CreatedAt,
		&ticketRsp.ClosedAt,
		&ticketRsp.Status,
	)
	if err != nil {
		return ticketRsp, err
	}

	return ticketRsp, nil
}

func (q *Queries) CreateTicketMessage(ctx context.Context, args params.CreateTicketMessage) error {
	const query = "INSERT INTO ticket_messages(ticket_id, user_id, content) VALUES (?, ?, ?)"
	_, err := q.db.ExecContext(ctx, query, args.TicketId, args.UserId, args.Content)

	return err
}

func (q *Queries) SetTicketStatus(ctx context.Context, status string, ticketId int64) error {
	const query = "UPDATE tickets SET status = ?, closed_at = ? WHERE id = ?"
	_, err := q.db.ExecContext(ctx, query, status, time.Now(), ticketId)

	return err
}

func (q *Queries) CreateCourseCategory(ctx context.Context, name string) error {
	const query = "INSERT INTO course_categories(name) VALUES (?)"
	_, err := q.db.ExecContext(ctx, query, name)

	return err
}

func (q *Queries) GetCourseCategories(ctx context.Context) ([]courseDto.CategoryResponse, error) {
	const query = "SELECT id, name FROM course_categories"
	categoryRsp := make([]courseDto.CategoryResponse, 0)

	rows, err := q.db.QueryContext(ctx, query)
	if err != nil {
		return categoryRsp, err
	}

	defer rows.Close()

	for rows.Next() {
		var category courseDto.CategoryResponse
		err := rows.Scan(&category.Value, &category.Label)
		if err != nil {
			return categoryRsp, err
		}

		categoryRsp = append(categoryRsp, category)
	}

	return categoryRsp, err
}

func (q *Queries) FindCourseCategoryById(ctx context.Context, categoryId int64) (models.CourseCategory, error) {
	const query = "SELECT * FROM course_categories WHERE id = ?"

	row := q.db.QueryRowContext(ctx, query, categoryId)

	var categoryRsp models.CourseCategory
	err := row.Scan(
		&categoryRsp.Id,
		&categoryRsp.Name,
	)
	if err != nil {
		return categoryRsp, err
	}

	return categoryRsp, nil
}

func (q *Queries) CreateCourse(ctx context.Context, args params.CreateCourse) (int64, error) {
	const query = "INSERT INTO courses(name, code, description, image, category_id) VALUES (?, ?, ? , ?, ?)"
	result, err := q.db.ExecContext(ctx, query, args.Name, args.Code, args.Description, args.Image, args.CategoryId)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	return id, err
}

func (q *Queries) CreateCourseSection(ctx context.Context, args params.CreateCourseSection) error {
	const query = "INSERT INTO course_sections(title, content, course_id) VALUES (?, ?, ?)"
	_, err := q.db.ExecContext(ctx, query, args.Title, args.Content, args.CourseId)

	return err
}

func (q *Queries) GetSectionsByCourseId(ctx context.Context, courseId int64) ([]courseDto.SectionResponse, error) {
	const query = "SELECT id, title FROM course_sections WHERE course_id = ?"

	sectionsRsp := make([]courseDto.SectionResponse, 0)

	rows, err := q.db.QueryContext(ctx, query, courseId)
	if err != nil {
		return sectionsRsp, err
	}

	defer rows.Close()

	for rows.Next() {
		var section courseDto.SectionResponse
		err := rows.Scan(
			&section.Id,
			&section.Title,
		)
		if err != nil {
			return sectionsRsp, err
		}
		sectionsRsp = append(sectionsRsp, section)
	}

	return sectionsRsp, nil
}

func (q *Queries) GetCourses(ctx context.Context, args queryParams.CourseQueryParams) ([]courseDto.CourseResponse, error) {
	var sb strings.Builder
	sb.WriteString("SELECT c.id, c.name, c.code, c.description, c.image, cc.name, c.created_at FROM courses c")
	sb.WriteString(" JOIN course_categories cc ON cc.id = c.category_id")
	if args.Search != "" || args.Category != "" {
		sb.WriteString(" WHERE")
		if args.Search != "" {
			sb.WriteString(" description LIKE '%")
			sb.WriteString(args.Search)
			sb.WriteString("%'")
		}

		var stmts []string
		if args.Category != "" {
			stmts = append(stmts, " category_id = '"+args.Category+"'")
		}
		if args.Search != "" && len(stmts) > 0 {
			sb.WriteString(" AND ")
		}
		sb.WriteString(strings.Join(stmts, " AND "))
	}

	sb.WriteString(" ORDER BY created_at DESC")
	if args.Page > 0 {
		offset := (args.Page - 1) * args.Limit
		sb.WriteString(" LIMIT ")
		sb.WriteString(strconv.Itoa(args.Limit))
		sb.WriteString(" ")
		sb.WriteString(" OFFSET ")
		sb.WriteString(strconv.Itoa(offset))
	}

	var courses = make([]courseDto.CourseResponse, 0)

	rows, err := q.db.QueryContext(ctx, sb.String())
	if err != nil {
		return courses, err
	}

	defer rows.Close()

	for rows.Next() {
		var course courseDto.CourseResponse
		var createdAt sql.NullTime
		err := rows.Scan(
			&course.Id,
			&course.Name,
			&course.Code,
			&course.Description,
			&course.Image,
			&course.Category,
			&createdAt,
		)

		if err != nil {
			return courses, err
		}
		course.CreatedAt = createdAt.Time.Format(time.RFC3339)
		courses = append(courses, course)
	}

	return courses, nil
}

func (q *Queries) CountCourses(ctx context.Context, args queryParams.CourseQueryParams) (int64, error) {
	var sb strings.Builder
	sb.WriteString("SELECT count(*) FROM courses c")
	sb.WriteString(" JOIN course_categories cc ON cc.id = c.category_id")
	if args.Search != "" || args.Category != "" {
		sb.WriteString(" WHERE")
		if args.Search != "" {
			sb.WriteString(" description LIKE '%")
			sb.WriteString(args.Search)
			sb.WriteString("%'")
		}

		var stmts []string
		if args.Category != "" {
			stmts = append(stmts, " category_id = '"+args.Category+"'")
		}
		if args.Search != "" && len(stmts) > 0 {
			sb.WriteString(" AND ")
		}
		sb.WriteString(strings.Join(stmts, " AND "))
	}

	row := q.db.QueryRowContext(ctx, sb.String())

	var count int64
	err := row.Scan(&count)

	return count, err
}

func (q *Queries) GetCourseById(ctx context.Context, courseId int64) (courseDto.CourseResponse, error) {
	const query = `SELECT c.id, c.name, c.code, c.description, c.image, cc.name, c.created_at FROM courses c
	JOIN course_categories cc ON cc.id = c.category_id
	WHERE c.id = ?
	`

	row := q.db.QueryRowContext(ctx, query, courseId)

	var course courseDto.CourseResponse
	var createdAt sql.NullTime
	err := row.Scan(
		&course.Id,
		&course.Name,
		&course.Code,
		&course.Description,
		&course.Image,
		&course.Category,
		&createdAt,
	)
	if err != nil {
		return course, err
	}

	course.CreatedAt = createdAt.Time.Format(time.RFC3339)
	return course, nil
}

func (q *Queries) CreateCourseEnroll(ctx context.Context, userId int64, courseId int64) error {
	const query = "INSERT INTO user_enrolled(user_id, course_id) VALUES(?, ?)"
	_, err := q.db.ExecContext(ctx, query, userId, courseId)

	return err
}

func (q *Queries) GetCourseEnroll(ctx context.Context, userId int64, courseId int64) (models.CourseEnroll, error) {
	const query = "SELECT user_id, course_id from course_enroll WHERE user_id = ? AND course_id = ?"
	row := q.db.QueryRowContext(ctx, query, userId, courseId)

	var enroll models.CourseEnroll
	err := row.Scan(&enroll.UserId, &enroll.CourseId)
	return enroll, err
}
