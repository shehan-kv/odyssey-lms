package params

type CreateTicket struct {
	Subject     string
	Description string
	UserId      int64
	Type        string
	Status      string
}
