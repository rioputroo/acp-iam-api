package api

import (
	"acp-iam-api/api/v1/auth"
	"github.com/labstack/echo/v4"
)

//RegisterPath Register all API with routing path
func RegisterPath(e *echo.Echo, authController *auth.Controller) {
	if authController == nil {
		panic("Controller parameter cannot be nil")
	}

	//auth with Versioning endpoint
	userV1 := e.Group("api/v1/iam/auth")
	userV1.POST("/login", authController.Login)
	userV1.POST("/register", authController.Register)
}
