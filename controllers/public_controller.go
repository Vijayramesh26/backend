package controllers

import (
	"net/http"

	"portfolio-admin-api/config"
	"portfolio-admin-api/models"

	"github.com/gin-gonic/gin"
)

/* GET PORTFOLIO BY USERNAME */
func GetPublicPortfolio(c *gin.Context) {
	username := c.Param("username")

	var portfolio models.Portfolio
	err := config.DB.
		Where("username = ? AND published = ?", username, true).
		First(&portfolio).Error

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Portfolio not found"})
		return
	}

	c.JSON(http.StatusOK, portfolio)
}

/* GET ENABLED SECTIONS */
func GetPublicSections(c *gin.Context) {
	username := c.Param("username")

	var portfolio models.Portfolio
	if err := config.DB.
		Where("username = ? AND published = ?", username, true).
		First(&portfolio).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Portfolio not found"})
		return
	}

	var sections []models.Section
	config.DB.
		Where("portfolio_id = ? AND enabled = ?", portfolio.ID, true).
		Order("`order` ASC").
		Find(&sections)

	c.JSON(http.StatusOK, sections)
}
