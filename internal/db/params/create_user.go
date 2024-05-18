package params

import "database/sql"

type CreateUser struct {
	FirstName  string
	LastName   string
	Email      string
	Password   string
	AvatarName sql.NullString
	Bio        sql.NullString
}
