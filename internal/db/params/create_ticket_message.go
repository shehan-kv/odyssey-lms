package params

type CreateTicketMessage struct {
	TicketId int64
	UserId   int64
	Content  string
}
