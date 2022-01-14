package controllers

import (
	"go-lightning/httputil"
	models "go-lightning/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// FindStrikes godoc
// @Summary      Get all lightning strikes
// @Description  Get lightning strikes
// @Tags         lightning strikes
// @Accept       json
// @Produce      json
// @Success      200      {array}   models.LightningStrike
// @Failure      400      {object}  httputil.HTTPError
// @Router       /strikes [get]
func FindStrikes(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var strikes []models.LightningStrike

	if err := db.Find(&strikes).Error; err != nil {
		httputil.NewError(c, http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, strikes)
}

// CreateStrike godoc
// @Summary      Create new lightning strike
// @Description  Post lightning strike
// @Tags         lightning strike
// @Accept       json
// @Produce      json
// @Param        message  body      models.CreateStrikeInput  true  "Lightning Strike Information"
// @Success      200      {object}  models.LightningStrike
// @Failure      400      {object}  httputil.HTTPError
// @Router       /strikes [post]
func CreateStrike(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	// Validate input
	var input models.CreateStrikeInput
	if err := c.ShouldBindJSON(&input); err != nil {
		httputil.NewError(c, http.StatusBadRequest, err)
		return
	}

	// Create strike
	strike := models.LightningStrike{StrikeTime: input.StrikeTime, XCoord: input.XCoord, YCoord: input.YCoord}
	db.Create(&strike)

	c.JSON(http.StatusOK, strike)

}

// FindStrike godoc
// @Summary      Get a lightning strike by id
// @Description  Get a lightning strike
// @Tags         lightning strike
// @Accept       json
// @Produce      json
// @Param        id path int true "ID"
// @Success      200      {object}  models.LightningStrike
// @Failure      400      {object}  httputil.HTTPError
// @Router       /strikes/{id} [get]
func FindStrike(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	// Get model if it exists
	var strike models.LightningStrike
	if err := db.Where("id = ?", c.Param("id")).First(&strike).Error; err != nil {
		httputil.NewError(c, http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, strike)
}

// UpdateStrike godoc
// @Summary      Update a lightning strike by id
// @Description  Update a lightning strike
// @Tags         lightning strike
// @Accept       json
// @Produce      json
// @Param        id path int true "ID"
// @Param        message  body      models.UpdateStrikeInput  true  "Lightning Strike Information"
// @Success      200      {object}  models.LightningStrike
// @Failure      400      {object}  httputil.HTTPError
// @Router       /strikes/{id} [patch]
func UpdateStrike(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	// Get model if it exists
	var strike models.LightningStrike
	if err := db.Where("id = ?", c.Param("id")).First(&strike).Error; err != nil {
		httputil.NewError(c, http.StatusBadRequest, err)
		return
	}

	// Validate input
	var input models.UpdateStrikeInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.Model(&strike).Updates(input)

	c.JSON(http.StatusOK, strike)
}

// DeleteStrike godoc
// @Summary      Delete a lightning strike by id
// @Description  Delete a lightning strike
// @Tags         lightning strike
// @Accept       json
// @Produce      json
// @Param        id path int true "ID"
// @Success      200      {boolean} true
// @Failure      400      {object}  httputil.HTTPError
// @Router       /strikes/{id} [delete]
func DeleteStrike(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	// Get model if exist
	var strike models.LightningStrike
	if err := db.Where("id = ?", c.Param("id")).First(&strike).Error; err != nil {
		httputil.NewError(c, http.StatusBadRequest, err)
		return
	}

	db.Delete(&strike)

	c.JSON(http.StatusOK, true)
}
