package service

import (
	"context"
	"errors"

	"odyssey.lms/internal/db"
	"odyssey.lms/internal/db/params"
	queryParams "odyssey.lms/internal/dto/params"
	dto "odyssey.lms/internal/dto/ticket"
	"odyssey.lms/internal/middleware"
)

func CreateSupportTicket(ctx context.Context, args dto.TicketCreateRequest) error {
	userId, ok := ctx.Value(middleware.USER_ID).(int64)
	if !ok {
		return errors.New("could not get user-id from context")
	}
	_, err := db.QUERY.FindUserById(ctx, userId)
	if err != nil {
		return ErrUserNotFound
	}

	err = db.QUERY.CreateTicket(ctx, params.CreateTicket{
		Subject:     args.Subject,
		Description: args.Description,
		Type:        args.Type,
		UserId:      int64(userId),
		Status:      "unresolved",
	})

	_ = db.QUERY.CreateEvent(ctx, params.CreateEvent{
		Type:        "user",
		Severity:    "info",
		Description: "Support ticket created by",
	})

	return err
}

func GetSupportTickets(ctx context.Context, args queryParams.TicketQueryParams) (dto.TicketsResponse, error) {

	var ticketsRsp dto.TicketsResponse
	tickets, err := db.QUERY.GetTickets(ctx, args)
	if err != nil {
		return ticketsRsp, err
	}

	ticketsCount, err := db.QUERY.CountTickets(ctx, args)
	if err != nil {
		return ticketsRsp, err
	}

	ticketsRsp.TotalCount = ticketsCount
	ticketsRsp.Tickets = tickets

	return ticketsRsp, nil
}

func GetSupportTicketsSelf(ctx context.Context, args queryParams.TicketQueryParams) (dto.TicketsResponse, error) {
	var ticketsRsp dto.TicketsResponse

	userId, ok := ctx.Value(middleware.USER_ID).(int64)
	if !ok {
		return ticketsRsp, errors.New("could not get user-id from context")
	}
	_, err := db.QUERY.FindUserById(ctx, userId)
	if err != nil {
		return ticketsRsp, ErrUserNotFound
	}

	tickets, err := db.QUERY.GetTicketsByUserId(ctx, userId, args)
	if err != nil {
		return ticketsRsp, err
	}

	ticketsCount, err := db.QUERY.CountTicketsByUserId(ctx, userId, args)
	if err != nil {
		return ticketsRsp, err
	}

	ticketsRsp.TotalCount = ticketsCount
	ticketsRsp.Tickets = tickets

	return ticketsRsp, nil
}

func GetSupportTicketSelf(ctx context.Context, ticketId int64) (dto.TicketMessagesResponse, error) {
	var ticketRsp dto.TicketMessagesResponse

	userId, ok := ctx.Value(middleware.USER_ID).(int64)
	if !ok {
		return ticketRsp, errors.New("could not get user-id from context")
	}

	existingTicket, err := db.QUERY.FindTicketById(ctx, ticketId)
	if err != nil {
		return ticketRsp, err
	}

	if existingTicket.UserId != userId {
		return ticketRsp, ErrNotAllowed
	}

	tickets, err := db.QUERY.GetTicketByIdWithUser(ctx, ticketId)
	if err != nil {
		return ticketRsp, err
	}

	messages, err := db.QUERY.GetTicketMessagesByTicketId(ctx, ticketId)
	if err != nil {
		return ticketRsp, err
	}

	ticketRsp.Ticket = tickets
	ticketRsp.Messages = messages

	return ticketRsp, nil
}

func GetSupportTicketById(ctx context.Context, ticketId int64) (dto.TicketMessagesResponse, error) {
	var ticketRsp dto.TicketMessagesResponse

	tickets, err := db.QUERY.GetTicketByIdWithUser(ctx, ticketId)
	if err != nil {
		return ticketRsp, err
	}

	messages, err := db.QUERY.GetTicketMessagesByTicketId(ctx, ticketId)
	if err != nil {
		return ticketRsp, err
	}

	ticketRsp.Ticket = tickets
	ticketRsp.Messages = messages

	return ticketRsp, nil
}

var ErrRoleNotFound = errors.New("role not found")
var ErrTicketNotFound = errors.New("ticket not found")
var ErrNotAllowed = errors.New("not allowed")
var ErrTicketAlreadyClosed = errors.New("ticket already closed")

func CreateSupportTicketMessage(ctx context.Context, ticketId int64, args dto.TicketMessageCreateRequest) error {
	userId, ok := ctx.Value(middleware.USER_ID).(int64)
	if !ok {
		return errors.New("could not get user-id from context")
	}
	user, err := db.QUERY.FindUserById(ctx, userId)
	if err != nil {
		return ErrUserNotFound
	}

	role, err := db.QUERY.FindRoleById(ctx, user.Role)
	if err != nil {
		return ErrRoleNotFound
	}

	ticket, err := db.QUERY.FindTicketById(ctx, ticketId)
	if err != nil {
		return ErrTicketNotFound
	}

	if ticket.Status == "resolved" {
		return ErrTicketAlreadyClosed
	}

	if role.Name != "administrator" && ticket.UserId != userId {
		return ErrNotAllowed
	}

	err = db.QUERY.CreateTicketMessage(ctx, params.CreateTicketMessage{
		TicketId: ticketId,
		UserId:   userId,
		Content:  args.Message,
	})

	return err
}

func ResolveTicket(ctx context.Context, ticketId int64) error {

	_, err := db.QUERY.FindTicketById(ctx, ticketId)
	if err != nil {
		return ErrTicketNotFound
	}

	err = db.QUERY.SetTicketStatus(ctx, "resolved", ticketId)

	_ = db.QUERY.CreateEvent(ctx, params.CreateEvent{
		Type:        "user",
		Severity:    "info",
		Description: "Support ticket closed",
	})

	return err
}
