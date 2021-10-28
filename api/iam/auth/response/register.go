package response

//RegisterResponse payload
type RegisterResponse struct {
	Email string `json:"email"`
	Msg   string `json:"msg"`
}

//NewRegisterResponse construct LoginResponse
func NewRegisterResponse(name string, email string) *RegisterResponse {
	var registerResponse RegisterResponse

	registerResponse.Email = email
	registerResponse.Msg = "Successfully register the user"

	return &registerResponse
}
