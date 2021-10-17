package user

//Service outgoing port for auth
type Service interface {
	//FindUserByEmailAndPassword find auth by form email and password
	FindUserByEmailAndPassword(email string, password string) (*User, error)
	Login(email string, password string) (string, error)
	Register(email string, password string) (string, error)
}

//Repository in going port for auth
type Repository interface {
	//FindUserByEmailAndPassword find auth by form email and password
	FindUserByEmailAndPassword(email string, password string) (*User, error)
	Login(email string, password string) (string, error)
	Register(email string, password string) (string, error)
}
