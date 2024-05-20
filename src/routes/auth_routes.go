package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/pedro-costa22/first-crud-go/src/controller"
	"github.com/pedro-costa22/first-crud-go/src/repository"
	"github.com/pedro-costa22/first-crud-go/src/service"
	"gorm.io/gorm"
)

func AuthRoutes(r *gin.Engine, db *gorm.DB) {
	repository := repository.NewUserRepository(db)
	service := service.NewUserService(repository)
	controller := controller.NewAuthController(service)

	r.POST("/login", controller.Login)
}