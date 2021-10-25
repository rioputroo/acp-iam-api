package roles

import (
	"acp-iam-api/api/common"
	"acp-iam-api/api/iam/roles/request"
	responses "acp-iam-api/api/iam/roles/response"
	"acp-iam-api/business/roles"
	"github.com/labstack/echo/v4"
	"strconv"
)

//Controller Get roles API controller
type Controller struct {
	service roles.Service
}

//NewController Construct roles API controller
func NewController(service roles.Service) *Controller {
	return &Controller{
		service,
	}
}

func (controller *Controller) GetAllRoles(c echo.Context) error {
	roles, err := controller.service.GetAllRoles()

	if err != nil {
		return c.JSON(common.NewErrorBusinessResponse(err))
	}

	response := responses.NewGetAllRolesResponse(roles)

	return c.JSON(common.NewSuccessResponse(response))
}

func (controller *Controller) AddRoles(c echo.Context) error {
	insertRolesRequest := new(request.InsertRolesRequest)

	if err := c.Bind(insertRolesRequest); err != nil {
		return c.JSON(common.NewBadRequestResponse())
	}

	err := controller.service.AddRoles(insertRolesRequest.Name, insertRolesRequest.IsActive)

	if err != nil {
		return c.JSON(common.NewErrorBusinessResponse(err))
	}

	return c.JSON(common.NewSuccessResponseWithoutData())
}

func (controller *Controller) GetRoles(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	getRoles, err := controller.service.GetRoles(uint(id))

	if err != nil {
		return c.JSON(common.NewErrorBusinessResponse(err))
	}

	response := responses.NewGetRolesResponse(*getRoles)

	return c.JSON(common.NewSuccessResponse(response))
}

func (controller *Controller) UpdateRoles(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	updateRolesRequest := new(request.UpdateRolesRequest)

	if err := c.Bind(updateRolesRequest); err != nil {
		return c.JSON(common.NewBadRequestResponse())
	}

	err := controller.service.UpdateRoles(uint(id), updateRolesRequest.Name, updateRolesRequest.IsActive)

	if err != nil {
		return c.JSON(common.NewErrorBusinessResponse(err))
	}

	return c.JSON(common.NewSuccessResponseWithoutData())
}

func (controller *Controller) DeleteRoles(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	updateRolesRequest := new(request.UpdateRolesRequest)

	if err := c.Bind(updateRolesRequest); err != nil {
		return c.JSON(common.NewBadRequestResponse())
	}

	err := controller.service.DeleteRoles(uint(id))

	if err != nil {
		return c.JSON(common.NewErrorBusinessResponse(err))
	}

	return c.JSON(common.NewSuccessResponseWithoutData())
}
