package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func DatabaseConnection() *gorm.DB {
	err := godotenv.Load()
		if err != nil {
			fmt.Println("Error loading .env file")
		}

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	sqlInfo := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", user, password, host, port, dbName)

	db, err := gorm.Open(mysql.Open(sqlInfo), &gorm.Config{})
	if err != nil {
        panic(err)
    }
	
	fmt.Print("Conexão com o banco bem sucessida!")
	return db
}