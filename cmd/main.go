package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	config "github.com/pedro-costa22/first-crud-go/src/config/database"
	"github.com/pedro-costa22/first-crud-go/src/routes"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db := config.DatabaseConnection()
	config.GenerateMigrations(db)
	
	router := gin.Default()
	routes.InitRoutes(router, db)
	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}