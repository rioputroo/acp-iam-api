package auth

import (
	"acp-iam-api/business"
	"acp-iam-api/business/users"
	"acp-iam-api/middleware"
)

//=============== The implementation of those interface put below =======================
type service struct {
	userService users.Service
}

//NewService Construct user service object
func NewService(userService users.Service) Service {
	return &service{
		userService,
	}
}

//Login by given user Username and Password, return error if not exist
func (s *service) Login(email string, password string) (*users.Users, string, error) {
	user, err := s.userService.Login(email, password)

	if err != nil {
		return nil, "", nil
	} else if user == nil {
		return nil, "", business.ErrNotFound
	}

	token, err := middleware.CreateToken(user.Email, user.ID, user.RolesId)

	if err != nil {
		return nil, "", nil
	}

	return user, token, nil
}

func (s *service) Register(email string, password string) (users.UsersCreds, error) {
	userRegister, err := s.userService.Register(email, password)

	if err != nil {
		return users.UsersCreds{}, err
	}

	return *userRegister, nil
}

func (s *service) FindUserByEmail(email string) bool {
	return s.userService.FindUserByEmail(email)
}
