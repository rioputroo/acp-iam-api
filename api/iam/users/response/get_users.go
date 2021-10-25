package response

import (
	"acp-iam-api/business/roles"
	"acp-iam-api/business/users"
)

type GetUsersRoleResponse struct {
	id   uint
	name string
}

type GetUsersResponse struct {
	ID       uint   `json:"id"`
	Email    string `json:"email"`
	Name     string `json:"name"`
	RolesId  uint   `json:"roles_id"`
	IsActive bool   `json:"is_active"`
}

func NewGetUsersResponse(users users.Users, roles roles.Roles) GetUsersResponse {
	var getUsersResponse GetUsersResponse

	getUsersResponse.ID = users.ID
	getUsersResponse.Name = users.Name
	getUsersResponse.Email = users.Email
	getUsersResponse.IsActive = users.IsActive
	getUsersResponse.RolesId = users.RolesId

	return getUsersResponse
}
