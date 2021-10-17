package auth

//Service outgoing port for auth auth
type Service interface {
	//Login auth into the system, return jwt token
	Login(email string, password string) (string, error)

	//Register add new user into the sytem
	Register(email string, password string) (string, error)
}
