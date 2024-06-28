package service

import (
	"context"
	"database/sql"
	"errors"

	"odyssey.lms/internal/db"
	"odyssey.lms/internal/db/params"
	dto "odyssey.lms/internal/dto/course"
	queryParams "odyssey.lms/internal/dto/params"
	"odyssey.lms/internal/middleware"
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

func GetCourses(ctx context.Context, args queryParams.CourseQueryParams) (dto.CoursesResponse, error) {
	var coursesRsp dto.CoursesResponse
	courses, err := db.QUERY.GetCourses(ctx, args)
	if err != nil {
		return coursesRsp, err
	}

	count, err := db.QUERY.CountCourses(ctx, args)
	if err != nil {
		return coursesRsp, err
	}

	coursesRsp.TotalCount = count
	coursesRsp.Courses = courses

	return coursesRsp, err
}

func GetCourseById(ctx context.Context, courseId int64) (dto.CourseResponse, error) {

	course, err := db.QUERY.GetCourseById(ctx, courseId)
	if err != nil {
		return course, err
	}

	userId, ok := ctx.Value(middleware.USER_ID).(int64)
	if !ok {
		return course, errors.New("could not get user-id from context")
	}

	sections, err := db.QUERY.GetSectionsByCourseId(ctx, courseId)
	if err != nil {
		return course, err
	}

	_, err = db.QUERY.GetCourseEnroll(ctx, userId, courseId)
	if err != nil {
		course.IsEnrolled = false
	}

	course.Sections = sections
	course.IsEnrolled = true

	return course, nil
}

func EnrollInCourse(ctx context.Context, courseId int64) error {
	userId, ok := ctx.Value(middleware.USER_ID).(int64)
	if !ok {
		return errors.New("could not get user-id from context")
	}

	_, err := db.QUERY.GetCourseEnroll(ctx, userId, courseId)
	if err != nil {
		// Unexpected errors
		// sql.ErrNoRows is expected when the user isn't enrolled yet
		if !errors.Is(err, sql.ErrNoRows) {
			return err
		}
	}
	if err == nil {
		// User already assigned
		return nil
	}

	err = db.QUERY.CreateCourseEnroll(ctx, userId, courseId)
	return err
}
