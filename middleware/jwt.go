package middleware

import (
	"github.com/golang-jwt/jwt"
	"os"
	"time"
)

func CreateToken(email string, id uint, rolesId uint) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["id"] = id
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix() //Token expires after 1 hour
	claims["email"] = email
	claims["roles"] = rolesId

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(os.Getenv("ACP_IAM_API_JWT_SECRET")))
}
