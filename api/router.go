package api

import (
	"acp-iam-api/api/iam/auth"
	"acp-iam-api/api/iam/roles"
	"acp-iam-api/api/iam/users"
	"github.com/labstack/echo/v4"
)

func RegisterPath(e *echo.Echo, rolesController *roles.Controller, usersController *users.Controller, authController *auth.Controller) {
	if rolesController == nil || usersController == nil {
		panic("Controller parameter cannot be nil")
	}

	iamRolesRoutes := e.Group("api/iam/roles")
	//route to get all roles
	iamRolesRoutes.GET("", rolesController.GetAllRoles)
	//route to get single role
	iamRolesRoutes.GET("/:id", rolesController.GetRoles)
	//route to add new role
	iamRolesRoutes.POST("", rolesController.AddRoles)
	//route to update the role
	iamRolesRoutes.PUT("/:id", rolesController.UpdateRoles)
	//route to delete the role
	iamRolesRoutes.DELETE("/:id", rolesController.DeleteRoles)

	iamUsersRoutes := e.Group("api/iam/users")
	//route to get all users
	iamUsersRoutes.GET("", usersController.GetAllUsers)
	//route to get single user
	iamUsersRoutes.GET("/:id", usersController.GetUsers)
	//route to add new user
	iamUsersRoutes.POST("", usersController.AddUsers)
	//route to update the user
	iamUsersRoutes.PUT("/:id", usersController.UpdateUsers)
	//route to delete the user
	iamUsersRoutes.DELETE("/:id", usersController.DeleteUsers)

	iamAuthRoutes := e.Group("api/iam/auth")
	//route to login user
	iamAuthRoutes.POST("/login", authController.Login)
	//route to register user
	iamAuthRoutes.POST("/register", authController.Register)

}
