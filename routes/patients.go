package routes

import (
	"github.com/Himneesh-Kalra/makerble-coding-assessment/controllers"
	"github.com/Himneesh-Kalra/makerble-coding-assessment/middleware"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterPatientRoutes(r *gin.RouterGroup, db *gorm.DB) {
	r.Use(middleware.AuthMiddleware("doctor", "receptionist"))

	r.GET("/", controllers.GetAllPatients(db))
	r.GET("/:id", controllers.GetPatientByID(db))
	r.PUT("/:id", controllers.UpdatePatient(db))

	// Receptionist only
	r.POST("/", middleware.AuthMiddleware("receptionist"), controllers.CreatePatient(db))
	r.DELETE("/:id", middleware.AuthMiddleware("receptionist"), controllers.DeletePatient(db))
}
