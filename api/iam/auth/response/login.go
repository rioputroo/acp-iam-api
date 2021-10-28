package response

import "acp-iam-api/business/users"

//Login response payload
type LoginResponse struct {
	Email string `json:"email"`
	Token string `json:"token"`
}

//NewLoginResponse construct LoginResponse
func NewLoginResponse(users users.Users, token string) *LoginResponse {
	var LoginResponse LoginResponse

	LoginResponse.Email = users.Email
	LoginResponse.Token = token

	return &LoginResponse
}
