package auth

import (
	"acp-iam-api/api/common"
	"acp-iam-api/api/v1/auth/request"
	response2 "acp-iam-api/api/v1/auth/response"
	"acp-iam-api/business/auth"
	"github.com/labstack/echo/v4"
)

//Controller Get item API controller
type Controller struct {
	service auth.Service
}

//NewController Construct item API controller
func NewController(service auth.Service) *Controller {
	return &Controller{
		service,
	}
}

//Login by given username and password will return JWT token
func (controller *Controller) Login(c echo.Context) error {
	loginRequest := new(request.LoginRequest)

	if err := c.Bind(loginRequest); err != nil {
		return c.JSON(common.NewBadRequestResponse())
	}

	token, err := controller.service.Login(loginRequest.Email, loginRequest.Password)
	if err != nil {
		return c.JSON(common.NewErrorBusinessResponse(err))
	}

	response := response2.NewLoginResponse(token)

	return c.JSON(common.NewSuccessResponse(response))
}

//Register registering user into the system
func (controller *Controller) Register(c echo.Context) error {
	registerRequest := new(request.RegisterRequest)

	if err := c.Bind(registerRequest); err != nil {
		return c.JSON(common.NewBadRequestResponse())
	}

	token, err := controller.service.Register(registerRequest.Email, registerRequest.Password)

	if err != nil {
		return c.JSON(common.NewErrorBusinessResponse(err))
	}

	response := response2.NewRegisterResponse(token)

	return c.JSON(common.NewSuccessResponse(response))
}
