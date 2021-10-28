package main

import (
	"acp-iam-api/api"
	authController "acp-iam-api/api/iam/auth"
	rolesController "acp-iam-api/api/iam/roles"
	usersController "acp-iam-api/api/iam/users"
	authService "acp-iam-api/business/auth"
	rolesService "acp-iam-api/business/roles"
	"acp-iam-api/business/users"
	"acp-iam-api/config"
	rolesRepository "acp-iam-api/repository/roles"
	usersRepository "acp-iam-api/repository/users"
	"fmt"
	"github.com/labstack/echo/v4"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

func newDatabaseConnection(config *config.AppConfig) *gorm.DB {

	configDB := map[string]string{
		"DB_Username": os.Getenv("ACP_IAM_API_DB_USERNAME"),
		"DB_Password": os.Getenv("ACP_IAM_API_DB_PASSWORD"),
		"DB_Port":     os.Getenv("ACP_IAM_API_DB_PORT"),
		"DB_Host":     os.Getenv("ACP_IAM_API_DB_ADDRESS"),
		"DB_Name":     os.Getenv("ACP_IAM_API_DB_NAME"),
	}

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		configDB["DB_Username"],
		configDB["DB_Password"],
		configDB["DB_Host"],
		configDB["DB_Port"],
		configDB["DB_Name"])

	db, e := gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if e != nil {
		panic(e)
	}

	db.AutoMigrate(&rolesRepository.RolesTable{}, &usersRepository.UserTable{})

	return db
}

func main() {
	//load config if available or set to default
	config := config.GetConfig()

	////initial DB connection
	//dbConnection := initDB()

	//initialize database connection based on given config
	dbConnection := newDatabaseConnection(config)

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
	dsn := "root:rioputro93@tcp(database-acp-final-project.cph9s9nf5t0g.ap-southeast-1.rds.amazonaws.com)/acpfinalproject?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&rolesRepository.RolesTable{}, &usersRepository.UserTable{})

	return db
}
