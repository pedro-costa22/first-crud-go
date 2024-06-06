package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/pedro-costa22/first-crud-go/src/common/middleware"
	"github.com/pedro-costa22/first-crud-go/src/controller"
	"github.com/pedro-costa22/first-crud-go/src/repository"
	"github.com/pedro-costa22/first-crud-go/src/service"
	"gorm.io/gorm"
)

func UserRoutes(r *gin.RouterGroup, db *gorm.DB) {
	repository := repository.NewUserRepository(db)
	service := service.NewUserService(repository)
	controller := controller.NewUserController(service)

	r.GET("/:id", middleware.VerifyTokenMiddleware ,controller.FindByID)
	r.GET("/getUserByEmail/:email", middleware.VerifyTokenMiddleware, controller.FindByEmail)
	r.POST("", controller.Create)
	r.PATCH("/:id", middleware.VerifyTokenMiddleware, controller.Update)
	r.DELETE("/:id", middleware.VerifyTokenMiddleware, controller.Delete)
}