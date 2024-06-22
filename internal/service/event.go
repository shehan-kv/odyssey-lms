package service

import (
	"context"

	"odyssey.lms/internal/db"
	dto "odyssey.lms/internal/dto/event"
	queryParams "odyssey.lms/internal/dto/params"
)

func GetEvents(ctx context.Context, args queryParams.EventQueryParams) (dto.EventsResponse, error) {

	var eventsResp dto.EventsResponse
	events, err := db.QUERY.GetEvents(ctx, args)
	if err != nil {
		return eventsResp, err
	}

	count, err := db.QUERY.CountEvents(ctx, args)
	if err != nil {
		return eventsResp, err
	}

	eventsResp.Events = events
	eventsResp.TotalCount = count

	return eventsResp, err
}
