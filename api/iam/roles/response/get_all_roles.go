package response

import (
	"acp-iam-api/business/roles"
)

type GetAllRolesResponse struct {
	Roles []GetRolesResponse `json:"roles"`
}

//NewGetAllRolesResponse construct GetAllRolesResponse
func NewGetAllRolesResponse(roles []roles.Roles) GetAllRolesResponse {
	getAllRolesResponse := GetAllRolesResponse{}

	for _, value := range roles {

		var getRoleResponse GetRolesResponse

		getRoleResponse.ID = value.ID
		getRoleResponse.Name = value.Name
		getRoleResponse.IsActive = value.IsActive

		getAllRolesResponse.Roles = append(getAllRolesResponse.Roles, getRoleResponse)
	}

	if getAllRolesResponse.Roles == nil {
		getAllRolesResponse.Roles = []GetRolesResponse{}
	}

	return getAllRolesResponse
}
