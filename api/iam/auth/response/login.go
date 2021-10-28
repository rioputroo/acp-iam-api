package response

//Login response payload
type LoginResponse struct {
	Email string `json:"email"`
	Token string `json:"token"`
}

//NewLoginResponse construct LoginResponse
func NewLoginResponse(token string) *LoginResponse {
	var LoginResponse LoginResponse

	LoginResponse.Token = token

	return &LoginResponse
}
