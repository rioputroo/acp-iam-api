package request

//Register Request payload
type RegisterRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
