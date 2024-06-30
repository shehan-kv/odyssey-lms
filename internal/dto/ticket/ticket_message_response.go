package dto

type TicketMessageResponse struct {
	Id        string `json:"id"`
	User      string `json:"user"`
	CreatedAt string `json:"createdAt"`
	Content   string `json:"content"`
}
