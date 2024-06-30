package dto

type TicketQueryParams struct {
	Search string
	Page   int
	Limit  int
	Type   string
	Status string
}
