package dto

type EnrollSectionResponse struct {
	Id         int64  `json:"id"`
	Title      string `json:"title"`
	Content    string `json:"content"`
	IsComplete bool   `json:"isComplete"`
}
