package response

import (
	"acp-iam-api/business/users"
)

type GetAllUsersResponse struct {
	Users []GetUsersResponse `json:"users"`
}

//NewGetAllUsersResponse construct GetAllUsersResponse
func NewGetAllUsersResponse(users []users.Users) GetAllUsersResponse {
	getAllUsersResponse := GetAllUsersResponse{}

	for _, value := range users {

		var getUsersResponse GetUsersResponse

		getUsersResponse.ID = value.ID
		getUsersResponse.Name = value.Name
		getUsersResponse.Email = value.Email
		getUsersResponse.IsActive = value.IsActive
		getUsersResponse.RolesId = value.RolesId

		getAllUsersResponse.Users = append(getAllUsersResponse.Users, getUsersResponse)
	}

	if getAllUsersResponse.Users == nil {
		getAllUsersResponse.Users = []GetUsersResponse{}
	}

	return getAllUsersResponse
}
