package dto

type EnrollSectionResponse struct {
	Id         int64  `json:"id"`
	Title      string `json:"title"`
	IsComplete bool   `json:"isComplete"`
}
