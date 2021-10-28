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

	user, token, err := controller.service.Login(loginRequest.Email, loginRequest.Password)

	if err != nil {
		return c.JSON(common.NewErrorBusinessResponse(err))
	}

	response := response.NewLoginResponse(*user, token)

	return c.JSON(common.NewSuccessResponse(response))
}

func (controller *Controller) Register(c echo.Context) error {
	registerRequest := new(request.RegisterRequest)

	if err := c.Bind(registerRequest); err != nil {
		return c.JSON(common.NewBadRequestResponse())
	}

	err := controller.service.FindUserByEmail(registerRequest.Email)

	if err != true {
		return c.JSON(common.NewConflictResponse())
	}

	registerResponse, err2 := controller.service.Register(registerRequest.Email, registerRequest.Password)
	if err2 != nil {
		return c.JSON(common.NewErrorBusinessResponse(err2))
	}

	response := response.NewRegisterResponse(registerResponse.Name, registerResponse.Email)

	return c.JSON(common.NewSuccessResponse(response))
}
