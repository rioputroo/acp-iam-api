package users

import (
	"acp-iam-api/api/iam/users/request"
	"gorm.io/gorm"
)

type Users struct {
	gorm.Model
	Email    string
	Password string
	Name     string
	IsActive bool
	RolesId  uint
	Roles    string
}

//NewUsers create new users
func NewUsers(request request.InsertUsersRequest) Users {
	return Users{
		Name:     request.Name,
		Email:    request.Email,
		Password: request.Password,
		IsActive: request.IsActive,
		RolesId:  request.RolesId,
	}
}

//ModifyUsers update existing users data
func (oldData *Users) ModifyUsers(usersRequest *request.UpdateUsersRequest) Users {
	return Users{
		Name:     usersRequest.Name,
		Email:    usersRequest.Email,
		RolesId:  usersRequest.RolesId,
		IsActive: usersRequest.IsActive,
	}
}
