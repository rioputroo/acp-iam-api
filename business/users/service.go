package users

import (
	"acp-iam-api/api/iam/users/request"
	"acp-iam-api/business"
	"acp-iam-api/business/roles"
	"crypto/md5"
	"encoding/hex"
)

type service struct {
	repository Repository
}

//NewService Construct users service object
func NewService(repository Repository) Service {
	return &service{
		repository,
	}
}

func GetMD5Hash(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}

func (s service) GetAllUsers() ([]Users, error) {
	users, err := s.repository.GetAllUsers()

	if err != nil {
		return []Users{}, err
	}

	return users, err
}

func (s service) GetUsers(id uint) (*Users, *roles.Roles, error) {
	users, rolesData, err := s.repository.GetUsers(id)

	if err != nil {
		return nil, nil, business.ErrNotFound
	}

	return users, rolesData, nil
}

func (s service) AddUsers(request *request.InsertUsersRequest) error {
	request.Password = GetMD5Hash(request.Password)

	return s.repository.AddUsers(request)
}

func (s service) UpdateUsers(id uint, usersRequest *request.UpdateUsersRequest) error {
	usersData, _, err := s.repository.GetUsers(id)

	if err != nil {
		return nil
	} else if usersData == nil {
		return business.ErrNotFound
	}

	updatedUsers := usersData.ModifyUsers(usersRequest)

	return s.repository.UpdateUsers(id, &updatedUsers)
}

func (s service) DeleteUsers(id uint) error {
	return s.repository.DeleteUsers(id)
}

func (s service) Login(email string, password string) (*Users, error) {
	password = GetMD5Hash(password)

	login, err := s.repository.Login(email, password)

	if err != nil {
		return nil, nil
	}

	return login, nil
}

func (s service) Register(email string, password string) (*UsersCreds, error) {
	password = GetMD5Hash(password)

	return s.repository.Register(email, password)
}

func (s service) FindUserByEmail(email string) bool {
	return s.repository.FindUserByEmail(email)
}
