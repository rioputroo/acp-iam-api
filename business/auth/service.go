package auth

import (
	"acp-iam-api/business/user"
	"github.com/golang-jwt/jwt"
	"time"
)

type service struct {
	userService user.Service
}

func NewService(userService user.Service) Service {
	return &service{userService: userService}
}

//Register add new user to the system
func (s service) Register(email string, password string) (string, error) {
	_, err := s.userService.Register(email, password)

	if err != nil {
		return "", err
	}

	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["id"] = 1
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix() //Token expires after 1 hour
	claims["name"] = "Rio Trilaksono Putro"
	claims["email"] = "riotrilaksonop@gmail.com"

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte("secret"))
}

func (s service) Login(email string, password string) (string, error) {
	user, err := s.userService.FindUserByEmailAndPassword(email, password)
	if err != nil {
		return "", err
	}

	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["id"] = user.Id
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix() //Token expires after 1 hour
	claims["name"] = user.Name

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte("secret"))
}
