package controllers

import (
	"net/http"

	"portfolio-admin-api/config"
	"portfolio-admin-api/models"

	"github.com/gin-gonic/gin"
)

func CreateSection(c *gin.Context) {
	var section models.Section
	c.ShouldBindJSON(&section)
	config.DB.Create(&section)
	c.JSON(http.StatusCreated, section)
}

func GetSections(c *gin.Context) {
	var sections []models.Section
	portfolioID := c.Param("portfolioId")

	config.DB.Where("portfolio_id = ?", portfolioID).
		Order("`order` asc").
		Find(&sections)

	c.JSON(http.StatusOK, sections)
}
