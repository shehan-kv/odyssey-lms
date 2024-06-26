package dto

type TicketMessagesResponse struct {
	Ticket   TicketResponse          `json:"ticket"`
	Messages []TicketMessageResponse `json:"messages"`
}
