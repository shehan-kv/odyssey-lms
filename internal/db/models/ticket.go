package models

import "database/sql"

type Ticket struct {
	Id          int64
	Subject     string
	Description string
	UserId      int64
	Type        string
	CreatedAt   sql.NullTime
	ClosedAt    sql.NullTime
	Status      string
}
