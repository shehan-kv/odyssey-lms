package service

import (
	"context"
	"errors"

	"odyssey.lms/internal/db"
	"odyssey.lms/internal/db/params"
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
