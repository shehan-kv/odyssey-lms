package params

import "database/sql"

type UpdateUser struct {
	FirstName string
	LastName  string
	Bio       sql.NullString
}
