package params

import "database/sql"

type CreateUser struct {
	FirstName string
	LastName  string
	Email     string
	Password  string
	Bio       sql.NullString
}
