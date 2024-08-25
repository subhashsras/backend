package main

import (
	"fmt"

	"github.com/subhashsras/backend_service/config"
	"github.com/subhashsras/backend_service/db"
	"github.com/subhashsras/backend_service/routes"
)


func main() {
	config.InitializeEnv()
	db.InitializeDB()
	fmt.Println("Server started")
	routes.Start()

	
}