package dto

type TicketResponse struct {
	Id          int64  `json:"id"`
	Subject     string `json:"subject"`
	Description string `json:"description"`
	User        string `json:"user"`
	Type        string `json:"type"`
	CreatedAt   string `json:"createdAt"`
	ClosedAt    string `json:"closedAt"`
	Status      string `json:"status"`
}
