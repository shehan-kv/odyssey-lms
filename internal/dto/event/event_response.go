package dto

type EventResponse struct {
	Id          int64  `json:"id"`
	CreatedAt   string `json:"createdAt"`
	Type        string `json:"type"`
	Description string `json:"description"`
	Severity    string `json:"severity"`
}
