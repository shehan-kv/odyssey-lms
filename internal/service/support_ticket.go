package service

import (
	"context"
	"errors"
	"log"

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

	return err
}

func GetSupportTickets(ctx context.Context, args queryParams.TicketQueryParams) (dto.TicketsResponse, error) {

	var ticketsRsp dto.TicketsResponse
	tickets, err := db.QUERY.GetTickets(ctx, args)
	if err != nil {
		log.Println(err)
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
		log.Println(err)
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
