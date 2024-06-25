package dto

type TicketResponse struct {
	Id        int64  `json:"id"`
	Subject   string `json:"subject"`
	User      string `json:"user"`
	Type      string `json:"type"`
	CreatedAt string `json:"createdAt"`
	Status    string `json:"status"`
}
