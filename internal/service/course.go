package service

import (
	"context"

	"odyssey.lms/internal/db"
	"odyssey.lms/internal/db/params"
	dto "odyssey.lms/internal/dto/course"
)

func CreateCategory(ctx context.Context, args dto.CategoryCreateRequest) error {
	err := db.QUERY.CreateCourseCategory(ctx, args.Name)

	return err
}

func GetCategories(ctx context.Context) ([]dto.CategoryResponse, error) {
	categories, err := db.QUERY.GetCourseCategories(ctx)

	return categories, err
}

func CreateCourse(ctx context.Context, args dto.CourseCreateRequest) error {

	_, err := db.QUERY.FindCourseCategoryById(ctx, args.CategoryId)
	if err != nil {
		return err
	}

	courseId, err := db.QUERY.CreateCourse(ctx, params.CreateCourse{
		Name:        args.Name,
		Code:        args.Code,
		Description: args.Description,
		Image:       args.Image,
		CategoryId:  args.CategoryId,
	})
	if err != nil {
		return err
	}

	for _, s := range args.Sections {
		err := db.QUERY.CreateCourseSection(ctx, params.CreateCourseSection{
			Title:    s.Title,
			Content:  s.Content,
			CourseId: courseId,
		})
		if err != nil {
			return err
		}
	}

	return nil
}
