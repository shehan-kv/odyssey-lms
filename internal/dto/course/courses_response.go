package dto

type CoursesResponse struct {
	TotalCount int64            `json:"totalCount"`
	Courses    []CourseResponse `json:"courses"`
}
