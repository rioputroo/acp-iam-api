package request

//RegisterRequest payload
type RegisterRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
