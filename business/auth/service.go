package auth

import (
	"acp-iam-api/business/users"
	"github.com/golang-jwt/jwt"
	"os"
	"time"
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
func (s *service) Login(email string, password string) (string, error) {
	user, err := s.userService.Login(email, password)
	if err != nil {
		return "", nil
	}

	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["id"] = user.ID
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix() //Token expires after 1 hour
	claims["name"] = user.Name

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
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
