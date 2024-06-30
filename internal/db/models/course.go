package models

import "database/sql"

type Course struct {
	Id          int64
	Name        string
	Code        string
	Description string
	Image       string
	CategoryId  int64
	CreatedAt   sql.NullTime
}
