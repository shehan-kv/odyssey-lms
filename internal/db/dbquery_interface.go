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
}
