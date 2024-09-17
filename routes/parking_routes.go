package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"Gin/Hello/controllers"
)

// Setup
func SetupRoutes(router *gin.Engine, db *gorm.DB) {
	controllers.Setup(db)
	router.GET("/", controllers.Index)
	router.GET("/parking", controllers.GetParkings)
	router.GET("/parking/:id", controllers.GetParking)
	router.POST("/parking", controllers.CreateParking)
	router.PUT("/parking/:id", controllers.UpdateParking)
	router.DELETE("/parking/:id", controllers.DeleteParking)
}