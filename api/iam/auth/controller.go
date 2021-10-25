package auth

import (
	"acp-iam-api/api/common"
	"acp-iam-api/api/iam/auth/request"
	"acp-iam-api/api/iam/auth/response"
	"acp-iam-api/business/auth"
	"github.com/labstack/echo/v4"
)

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

	response := response.NewLoginResponse(token)

	return c.JSON(common.NewSuccessResponse(response))
}
