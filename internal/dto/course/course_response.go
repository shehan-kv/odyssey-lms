package dto

type CourseResponse struct {
	Id          int64  `json:"id"`
	Name        string `json:"name"`
	Code        string `json:"code"`
	Description string `json:"description"`
	Image       string `json:"image"`
	Category    string `json:"category"`
	CreatedAt   string `json:"createdAt"`
}
