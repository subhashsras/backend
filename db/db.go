package db

import (
	"fmt"

	"github.com/subhashsras/backend_service/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DBClient *gorm.DB

func InitializeDB() {
	ConnectionURI := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.EnvironmentData.DBUser,
		config.EnvironmentData.DBPass,
		config.EnvironmentData.DBHost,
		config.EnvironmentData.DBPort,
		config.EnvironmentData.DBName,
	)

	dbconn, err := gorm.Open(
		mysql.Open(ConnectionURI),
		&gorm.Config{},
	)

	if err != nil {
		fmt.Println("Cannot connect with DB:", err.Error())
		return
	}

	DBClient = dbconn
}