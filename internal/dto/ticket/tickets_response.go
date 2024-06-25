package dto

type TicketsResponse struct {
	TotalCount int64            `json:"totalCount"`
	Tickets    []TicketResponse `json:"tickets"`
}
