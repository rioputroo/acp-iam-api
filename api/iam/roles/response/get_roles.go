package response

import "acp-iam-api/business/roles"

type GetRolesResponse struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	IsActive bool   `json:"is_active"`
}

func NewGetRolesResponse(roles roles.Roles) GetRolesResponse {
	var getRolesResponse GetRolesResponse

	getRolesResponse.ID = roles.ID
	getRolesResponse.Name = roles.Name
	getRolesResponse.IsActive = roles.IsActive

	return getRolesResponse
}
