package service

import (
	"context"

	"odyssey.lms/internal/db"
	"odyssey.lms/internal/db/params"
	dto "odyssey.lms/internal/dto/user"
)

func GetUsers(ctx context.Context, page int, limit int, search string, role string) (dto.UsersResponse, error) {

	var resp dto.UsersResponse
	params := params.UserQueryParams{
		Page:   page,
		Limit:  limit,
		Search: search,
		Role:   role,
	}
	users, err := db.QUERY.GetUsers(ctx, params)
	if err != nil {
		return resp, err
	}

	userCount, err := db.QUERY.CountUsers(ctx, params)
	if err != nil {
		return resp, err
	}

	resp.TotalCount = userCount
	resp.Users = users
	return resp, err
}
