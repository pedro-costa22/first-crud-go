package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	config "github.com/pedro-costa22/first-crud-go/src/config/database"
	"github.com/pedro-costa22/first-crud-go/src/config/logger"
	"github.com/pedro-costa22/first-crud-go/src/routes"
)

var (
	PORT = "WEB_SERVER_PORT"
)

func main() {
	err := godotenv.Load()
	port := os.Getenv(PORT)
	if err != nil {
		logger.Error("Error loading .env file", err)
	}

	db := config.DatabaseConnection()
	config.GenerateMigrations(db)
	
	router := gin.Default()
	routes.InitRoutes(router, db)
	if err := router.Run(":" + port); err != nil {
		log.Fatal(err)
	}
}