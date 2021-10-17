package response

//RegisterResponse register response payload
type RegisterResponse struct {
	Token string `json:"token"`
}

//NewRegisterResponse construct RegisterResponse
func NewRegisterResponse(token string) *RegisterResponse {
	var RegisterResponse RegisterResponse

	RegisterResponse.Token = token

	return &RegisterResponse
}
