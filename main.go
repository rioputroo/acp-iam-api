package main

import (
	"acp-iam-api/api"
	_ "acp-iam-api/api/v1/auth"
	authController "acp-iam-api/api/v1/auth"
	authService "acp-iam-api/business/auth"
	userService "acp-iam-api/business/user"
	userRepository "acp-iam-api/modules/user"
	"fmt"
	"github.com/labstack/echo/v4"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

var (
	user = []User{
		{Name: "Rio", Email: "rio@gmail.com", Password: "test"},
		{Name: "Evan", Email: "evan@gmail.com", Password: "test"},
	}
	DB *gorm.DB
)

func main() {
	//initial DB connection
	dbConnection := initDB()

	//initial user repository
	userRepo := userRepository.NewGormDBRepository(dbConnection)

	//initiate user service
	userService := userService.NewService(userRepo)

	//initiate auth service
	authService := authService.NewService(userService)

	//initiate user controller
	authController := authController.NewController(authService)

	//create echo http
	e := echo.New()

	//register API path and handler
	api.RegisterPath(e, authController)

	e.Start(":8000")

}

func initDB() *gorm.DB {
	configDB := map[string]string{
		"DB_Username": os.Getenv("DB_USER"),
		"DB_Password": os.Getenv("DB_PASSWORD"),
	}

// 	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
// 		configDB["DB_Username"],
// 		configDB["DB_Password"],
// 		configDB["DB_Host"],
// 		configDB["DB_Port"],
// 		configDB["DB_Name"])
	
	dsn := fmt.Sprintf("host=172.31.3.115 user=%s password=%s dbname=acp_final_project port=5432 sslmode=disable TimeZone=Asia/Shanghai",
		configDB["DB_Username"],
		configDB["DB_Password"])
	
// 	dsn := "host=172.31.3.115 user=postgres password=Qhanau8oJsP7 dbname=acp_final_project port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	} else {
		fmt.Println("db connected")
	}

	db.AutoMigrate(&userRepository.UserTable{})

	return db
}
