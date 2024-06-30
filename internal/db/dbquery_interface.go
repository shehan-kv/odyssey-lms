package db

import (
	"context"

	"odyssey.lms/internal/db/models"
	"odyssey.lms/internal/db/params"
	courseDto "odyssey.lms/internal/dto/course"
	evntDto "odyssey.lms/internal/dto/event"
	queryParams "odyssey.lms/internal/dto/params"
	ticketDto "odyssey.lms/internal/dto/ticket"
	usrDto "odyssey.lms/internal/dto/user"
)

type DBQuery interface {
	FindUserWithPasswordByEmail(ctx context.Context, email string) (models.User, error)
	GetUsers(ctx context.Context, params queryParams.UserQueryParams) ([]usrDto.UserResponse, error)
	FindUserById(ctx context.Context, userId int64) (models.User, error)
	CreateUser(ctx context.Context, arg params.CreateUser) (int64, error)
	UpdateUser(ctx context.Context, userId int64, arg params.UpdateUser) error
	SetUserPassword(ctx context.Context, userId int64, password string) error
	SetUserLastSignInTime(ctx context.Context, userId int64) error
	GetSignUpStats(ctx context.Context) ([]usrDto.SignUpStat, error)
	CountUsersByRole(ctx context.Context, role string) (int64, error)
	CountUsers(ctx context.Context, arg queryParams.UserQueryParams) (int64, error)
	DeleteUserById(ctx context.Context, id int64) error
	SetUserIsActive(ctx context.Context, id int64, status bool) error
	CreateRole(ctx context.Context, name string) (models.Role, error)
	CountRoles(ctx context.Context) (int64, error)
	GetRoles(ctx context.Context) ([]models.Role, error)
	FindRoleById(ctx context.Context, roleId int64) (models.Role, error)
	AssignUserRole(ctx context.Context, arg params.AssignUserRole) error
	CreateEvent(ctx context.Context, arg params.CreateEvent) error
	GetEvents(ctx context.Context, arg queryParams.EventQueryParams) ([]evntDto.EventResponse, error)
	CountEvents(ctx context.Context, arg queryParams.EventQueryParams) (int64, error)
	CreateTicket(ctx context.Context, arg params.CreateTicket) error
	GetTickets(ctx context.Context, arg queryParams.TicketQueryParams) ([]ticketDto.TicketResponse, error)
	CountTickets(ctx context.Context, arg queryParams.TicketQueryParams) (int64, error)
	GetTicketsByUserId(ctx context.Context, userId int64, arg queryParams.TicketQueryParams) ([]ticketDto.TicketResponse, error)
	CountTicketsByUserId(ctx context.Context, userId int64, arg queryParams.TicketQueryParams) (int64, error)
	GetTicketByIdWithUser(ctx context.Context, ticketId int64) (ticketDto.TicketResponse, error)
	GetTicketMessagesByTicketId(ctx context.Context, ticketId int64) ([]ticketDto.TicketMessageResponse, error)
	FindTicketById(ctx context.Context, ticketId int64) (models.Ticket, error)
	CreateTicketMessage(ctx context.Context, args params.CreateTicketMessage) error
	SetTicketStatus(ctx context.Context, status string, ticketId int64) error
	CreateCourseCategory(ctx context.Context, name string) error
	GetCourseCategories(ctx context.Context) ([]courseDto.CategoryResponse, error)
	FindCourseCategoryById(ctx context.Context, categoryId int64) (models.CourseCategory, error)
	CreateCourse(ctx context.Context, args params.CreateCourse) (int64, error)
	CreateCourseSection(ctx context.Context, args params.CreateCourseSection) error
	GetSectionsByCourseId(ctx context.Context, courseId int64) ([]courseDto.SectionResponse, error)
	GetCourses(ctx context.Context, args queryParams.CourseQueryParams) ([]courseDto.CourseResponse, error)
	CountCourses(ctx context.Context, args queryParams.CourseQueryParams) (int64, error)
	GetCourseById(ctx context.Context, courseId int64) (courseDto.CourseResponse, error)
	CreateCourseEnroll(ctx context.Context, userId int64, courseId int64) error
	GetCourseEnroll(ctx context.Context, userId int64, courseId int64) (models.CourseEnroll, error)
	GetEnrolledCourses(ctx context.Context, userId int64) ([]courseDto.CourseResponse, error)
	GetEnrolledSectionsByCourseId(ctx context.Context, userId int64, courseId int64) ([]courseDto.EnrollSectionResponse, error)
	GetEnrolledSectionById(ctx context.Context, userId int64, sectionId int64) (courseDto.EnrollSectionResponse, error)
	GetCourseSectionComplete(ctx context.Context, userId int64, sectionId int64) (models.CourseSectionComplete, error)
	CreateCourseSectionComplete(ctx context.Context, userId int64, sectionId int64) error
}
