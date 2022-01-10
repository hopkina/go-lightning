package controllers

import (
	models "go-lightning/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// GET /strikes
// Get all lightning strikes
func FindStrikes(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var strikes []models.LightningStrike
	db.Find(&strikes)

	c.JSON(http.StatusOK, gin.H{"data": strikes})
}

// POST /strikes
// Create new lightning strikes
func CreateStrike(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	// Validate input
	var input models.CreateStrikeInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create strike
	strike := models.LightningStrike{StrikeTime: input.StrikeTime, XCoord: input.XCoord, YCoord: input.YCoord}
	db.Create(&strike)

	c.JSON(http.StatusOK, gin.H{"data": strike})

}

// GET /strikes/:id
// Find a book
func FindStrike(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	// Get model if it exists
	var strike models.LightningStrike
	if err := db.Where("id = ?", c.Param("id")).First(&strike).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": strike})
}

// PATCH /strikes/:id
// Update a lightning strike
func UpdateStrike(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	// Get model if it exists
	var strike models.LightningStrike
	if err := db.Where("id = ?", c.Param("id")).First(&strike).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input models.UpdateStrikeInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.Model(&strike).Updates(input)

	c.JSON(http.StatusOK, gin.H{"data": strike})
}

// DELETE /strikes/:id
// Delete a book
func DeleteStrike(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	// Get model if exist
	var strike models.LightningStrike
	if err := db.Where("id = ?", c.Param("id")).First(&strike).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	db.Delete(&strike)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
