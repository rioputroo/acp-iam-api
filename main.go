package main

import (
	"acp-iam-api/api"
	authController "acp-iam-api/api/iam/auth"
	rolesController "acp-iam-api/api/iam/roles"
	usersController "acp-iam-api/api/iam/users"
	authService "acp-iam-api/business/auth"
	rolesService "acp-iam-api/business/roles"
	"acp-iam-api/business/users"
	rolesRepository "acp-iam-api/repository/roles"
	usersRepository "acp-iam-api/repository/users"
	"github.com/labstack/echo/v4"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	//initial DB connection
	dbConnection := initDB()

	//initial roles repository
	rolesRepo := rolesRepository.NewGormDBRepository(dbConnection)

	//initiate roles service
	rolesService := rolesService.NewService(rolesRepo)

	//initiate roles controller
	rolesController := rolesController.NewController(rolesService)

	//initial users repository
	usersRepo := usersRepository.NewGormDBRepository(dbConnection)

	//initiate users service
	usersService := users.NewService(usersRepo)

	//initiate users controller
	usersController := usersController.NewController(usersService)

	//initiate auth service
	authService := authService.NewService(usersService)

	//initiate users controller
	authController := authController.NewController(authService)

	e := echo.New()

	api.RegisterPath(e, rolesController, usersController, authController)

	e.Start(":8000")
}

func initDB() *gorm.DB {

	// 	dsn := "host=172.31.3.115 user=postgres password=Qhanau8oJsP7 dbname=acp_final_project port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	dsn := "host=localhost user=postgres password=gdn123 dbname=acp_final_project port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&rolesRepository.RolesTable{}, &usersRepository.UserTable{})

	return db
}
