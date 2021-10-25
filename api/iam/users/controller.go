package users

import (
	"acp-iam-api/api/common"
	"acp-iam-api/api/iam/users/request"
	responses "acp-iam-api/api/iam/users/response"
	"acp-iam-api/business/users"
	"fmt"
	"github.com/labstack/echo/v4"
	"strconv"
)

type Controller struct {
	service users.Service
}

//NewController Construct roles API controller
func NewController(service users.Service) *Controller {
	return &Controller{
		service,
	}
}

func (controller *Controller) GetAllUsers(c echo.Context) error {
	users, err := controller.service.GetAllUsers()

	if err != nil {
		return c.JSON(common.NewErrorBusinessResponse(err))
	}

	response := responses.NewGetAllUsersResponse(users)

	return c.JSON(common.NewSuccessResponse(response))
}

func (controller *Controller) GetUsers(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	getUsers, roles, err := controller.service.GetUsers(uint(id))
	fmt.Println(roles)

	if err != nil {
		return c.JSON(common.NewErrorBusinessResponse(err))
	}

	response := responses.NewGetUsersResponse(*getUsers, *roles)

	return c.JSON(common.NewSuccessResponse(response))
}

func (controller *Controller) AddUsers(c echo.Context) error {
	insertUsersRequest := new(request.InsertUsersRequest)

	if err := c.Bind(insertUsersRequest); err != nil {
		return c.JSON(common.NewBadRequestResponse())
	}

	err := controller.service.AddUsers(insertUsersRequest)

	if err != nil {
		return c.JSON(common.NewErrorBusinessResponse(err))
	}

	return c.JSON(common.NewSuccessResponseWithoutData())
}

func (controller *Controller) UpdateUsers(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	updateUserRequest := new(request.UpdateUsersRequest)

	if err := c.Bind(updateUserRequest); err != nil {
		return c.JSON(common.NewBadRequestResponse())
	}

	err := controller.service.UpdateUsers(uint(id), updateUserRequest)
	if err != nil {
		return c.JSON(common.NewErrorBusinessResponse(err))
	}

	return c.JSON(common.NewSuccessResponseWithoutData())
}

func (controller *Controller) DeleteUsers(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	updateUserRequest := new(request.UpdateUsersRequest)

	if err := c.Bind(updateUserRequest); err != nil {
		return c.JSON(common.NewBadRequestResponse())
	}

	err := controller.service.DeleteUsers(uint(id))
	if err != nil {
		return c.JSON(common.NewErrorBusinessResponse(err))
	}

	return c.JSON(common.NewSuccessResponseWithoutData())
}
