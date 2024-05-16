package config

import (
	"github.com/pedro-costa22/first-crud-go/src/config/database/entity"
	"gorm.io/gorm"
)

func GenerateMigrations(db *gorm.DB) {
	db.AutoMigrate(
		&entity.UserEntity{},
	)
}