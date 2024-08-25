package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)


type environmentDataType struct {
	DBUser string `validate:"required"` 
	DBPass string
	DBHost string
	DBPort string
	DBName string
}


var EnvironmentData *environmentDataType

func InitializeEnv() {
	_ = godotenv.Load("Config.env")
	
	EnvironmentData = &environmentDataType{
		DBUser: os.Getenv("DB_USER"),
		DBPass: os.Getenv("DB_PASS"),
		DBHost: os.Getenv("DB_HOST"),
		DBPort: os.Getenv("DB_PORT"),
		DBName: os.Getenv("DB_NAME"),
	}
	fmt.Println(EnvironmentData.DBHost)
	
	fmt.Println("Environment initialized.")

}

