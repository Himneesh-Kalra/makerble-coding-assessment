package routes

import (
	"github.com/Himneesh-Kalra/makerble-coding-assessment/controllers"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterAuthRoutes(r *gin.RouterGroup, db *gorm.DB) {
	r.POST("/login", controllers.Login(db))
	r.POST("/register", controllers.Register(db))
}
