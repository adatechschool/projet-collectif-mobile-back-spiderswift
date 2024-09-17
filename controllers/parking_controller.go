package controllers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"Gin/Hello/models"
)

var db *gorm.DB

func Setup(database *gorm.DB) {
	db = database
}

// DEFAULT
func Index(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello World",
	})
}

// POST
func CreateParking(c *gin.Context) {
	var parking models.Parking
	if err := c.ShouldBindJSON(&parking); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := db.Create(&parking).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, parking)
}

// GET one
func GetParking(c *gin.Context) {
	id := c.Param("id")
	var parking models.Parking
	if err := db.First(&parking, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Parking non trouvé"})
		return
	}
	c.JSON(http.StatusOK, parking)
}

// GET all
func GetParkings(c *gin.Context) {
	var parkings []models.Parking
	if err := db.Find(&parkings).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, parkings)
}

// PUT
func UpdateParking(c *gin.Context) {
	id := c.Param("id")
	var parking models.Parking
	if err := db.First(&parking, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Parking non trouvé"})
		return
	}
	if err := c.ShouldBindJSON(&parking); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := db.Save(&parking).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, parking)
}

// DELETE
func DeleteParking(c *gin.Context) {
	id := c.Param("id")
	if err := db.Delete(&models.Parking{}, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Parking non trouvé"})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}