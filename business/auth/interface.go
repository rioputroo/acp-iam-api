package auth

import "acp-iam-api/business/users"

//Service outgoing port for user
type Service interface {
	//Login If data not found will return nil without error
	Login(email string, password string) (*users.Users, string, error)
	Register(email string, password string) (users.UsersCreds, error)
	FindUserByEmail(email string) bool
}
