package dto

type UsersResponse struct {
	TotalCount int64          `json:"totalCount"`
	Users      []UserResponse `json:"users"`
}
