package service

import (
	"context"

	"odyssey.lms/internal/db"
	dto "odyssey.lms/internal/dto/course"
)

func CreateCategory(ctx context.Context, args dto.CategoryCreateRequest) error {
	err := db.QUERY.CreateCourseCategory(ctx, args.Name)

	return err
}
