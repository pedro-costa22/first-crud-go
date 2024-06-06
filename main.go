package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/pedro-costa22/first-crud-go/docs"
	config "github.com/pedro-costa22/first-crud-go/src/config/database"
	"github.com/pedro-costa22/first-crud-go/src/config/logger"
	"github.com/pedro-costa22/first-crud-go/src/routes"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var (
	PORT = "WEB_SERVER_PORT"
)

// @title Meu Primeiro CRUD em Go 
// @version 1.0
// @description API for crud operations on users

// @host localhost:8000
// @SecurityDefinitions.apiKey Bearer
// @In header
// @Name Authorization

// @schemes http
// @license MIT

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
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
	// Configuração do Swagger

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	if err := router.Run(":" + port); err != nil {
		log.Fatal(err)
	}
}