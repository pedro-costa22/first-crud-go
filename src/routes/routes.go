package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitRoutes(router *gin.Engine, db *gorm.DB) {
	UserRoutes(router.Group("/users"), db)
	AuthRoutes(router, db)
}

