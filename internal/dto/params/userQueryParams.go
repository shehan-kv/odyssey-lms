package dto

type UserQueryParams struct {
	Search string
	Page   int
	Limit  int
	Role   string
}
