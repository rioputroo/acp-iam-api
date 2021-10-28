package users

import (
	"acp-iam-api/api/iam/users/request"
	"acp-iam-api/business/roles"
)

type Service interface {
	GetAllUsers() ([]Users, error)
	GetUsers(id uint) (*Users, *roles.Roles, error)
	AddUsers(request *request.InsertUsersRequest) error
	UpdateUsers(id uint, usersRequest *request.UpdateUsersRequest) error
	DeleteUsers(id uint) error
	Login(email string, password string) (*Users, error)
	Register(email string, password string) (*UsersCreds, error)
}

type Repository interface {
	GetAllUsers() ([]Users, error)
	GetUsers(id uint) (*Users, *roles.Roles, error)
	AddUsers(request *request.InsertUsersRequest) error
	UpdateUsers(id uint, users Users) error
	DeleteUsers(id uint) error
	Login(email string, password string) (*Users, error)
	Register(email string, password string) (*UsersCreds, error)
}
