package dto

import "database/sql"

type UserResponse struct {
	ID        int64        `json:"id"`
	FirstName string       `json:"firstName"`
	LastName  string       `json:"lastName"`
	Email     string       `json:"email"`
	Password  string       `json:"password"`
	Role      string       `json:"role"`
	IsActive  bool         `json:"isActive"`
	CreatedAt sql.NullTime `json:"createdAt"`
	LastLogin sql.NullTime `json:"lastLogin"`
}
