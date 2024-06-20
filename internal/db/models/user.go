package models

import "database/sql"

type User struct {
	ID        int64
	FirstName string
	LastName  string
	Email     string
	Password  string
	Role      int64
	CreatedAt sql.NullTime
	LastLogin sql.NullTime
	Bio       sql.NullString
}
