package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"sync"
)

type ConfigIPForwarding struct {
	Enabled bool   `mapstructure:"enabled"`
	IP      string `mapstructure:"ip"`
	Port    string `mapstructure:"port"`
}

//AppConfig Application configuration
type AppConfig struct {
	AppPort        int    `mapstructure:"app_port"`
	AppEnvironment string `mapstructure:"app_environment"`
	DbDriver       string `mapstructure:"db_driver"`
	DbAddress      string `mapstructure:"db_address"`
	DbPort         int    `mapstructure:"db_port"`
	DbUsername     string `mapstructure:"db_username"`
	DbPassword     string `mapstructure:"db_password"`
	DbName         string `mapstructure:"db_name"`
	JwtSecretKey   string `mapstructure:"jwt_secret_key"`
}

var lock = &sync.Mutex{}
var appConfig *AppConfig

//GetConfig Initiatilize config in singleton way
func GetConfig() *AppConfig {
	if appConfig != nil {
		return appConfig
	}

	lock.Lock()
	defer lock.Unlock()

	//re-check after locking
	if appConfig != nil {
		return appConfig
	}

	appConfig = initConfig()

	return appConfig
}

func initConfig() *AppConfig {
	godotenv.Load(".env")

	var defaultConfig AppConfig

	defaultConfig.AppPort = 8000
	defaultConfig.AppEnvironment = ""
	defaultConfig.DbDriver = "mysql"
	defaultConfig.DbAddress = "localhost"
	defaultConfig.DbPort = 3306
	defaultConfig.DbUsername = "root"
	defaultConfig.DbPassword = ""
	defaultConfig.DbName = "acpfinalproject"
	defaultConfig.JwtSecretKey = "secret"

	viper.AutomaticEnv()
	viper.SetEnvPrefix("gohexagonal")
	viper.BindEnv("app_port")
	viper.BindEnv("app_environment")
	viper.BindEnv("db_driver")
	viper.BindEnv("db_address")
	viper.BindEnv("db_port")
	viper.BindEnv("db_username")
	viper.BindEnv("db_password")
	viper.BindEnv("db_name")

	var finalConfig AppConfig
	err := viper.Unmarshal(&finalConfig)
	if err != nil {
		fmt.Print("failed to extract config, will use default value")
		return &defaultConfig
	}

	return &finalConfig

	//var (
	//	err         error
	//	currdir     string
	//	finalConfig AppConfig
	//)
	//
	//currdir, err = os.Getwd()
	//if err != nil {
	//	log.Printf("Failed to get current directory, set to default")
	//	return &defaultConfig
	//}
	//
	//viper.SetConfigFile(currdir + "/config/.env")
	//
	//err = viper.ReadInConfig()
	//if err != nil {
	//	log.Printf("Failed read config, set to default")
	//	return &defaultConfig
	//}
	//
	//err = viper.Unmarshal(&finalConfig)
	//if err != nil {
	//	log.Printf("failed to extract config, will use default value")
	//	return &defaultConfig
	//}
	//
	//return &finalConfig
}
