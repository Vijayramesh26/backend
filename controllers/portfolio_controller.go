package controllers

import (
	"encoding/json"
	"net/http"
	"strings"

	"portfolio-admin-api/config"
	"portfolio-admin-api/models"

	"github.com/gin-gonic/gin"
)

// CREATE PORTFOLIO
func CreatePortfolio(c *gin.Context) {
	var portfolio models.Portfolio

	// Bind incoming JSON
	if err := c.ShouldBindJSON(&portfolio); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var profileData map[string]interface{}
	if err := json.Unmarshal(portfolio.Profile, &profileData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid profile JSON"})
		return
	}

	// profile is inside profileData["data"]
	name := ""
	if dataMap, ok := profileData["data"].(map[string]interface{}); ok {
		if n, ok := dataMap["name"].(string); ok {
			name = strings.TrimSpace(n)
		}
	}

	enabled, _ := profileData["enabled"].(bool)
	if enabled && name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Profile name is required"})
		return
	}

	// Generate unique username
	username := strings.ToLower(strings.ReplaceAll(name, " ", ""))
	portfolio.Username = username

	// Save portfolio
	if err := config.DB.Create(&portfolio).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, portfolio)
}

// GET ALL PORTFOLIOS
func GetPortfolios(c *gin.Context) {
	var portfolios []models.Portfolio
	config.DB.Find(&portfolios)
	c.JSON(http.StatusOK, portfolios)
}

// GET BY ID
func GetPortfolio(c *gin.Context) {
	id := c.Param("id")
	var portfolio models.Portfolio
	if err := config.DB.First(&portfolio, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Portfolio not found"})
		return
	}
	c.JSON(http.StatusOK, portfolio)
}

// UPDATE PORTFOLIO
func UpdatePortfolio(c *gin.Context) {
	id := c.Param("id")
	var portfolio models.Portfolio

	// Find existing
	if err := config.DB.First(&portfolio, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Portfolio not found"})
		return
	}

	// Bind new data
	if err := c.ShouldBindJSON(&portfolio); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Update username again if profile.name changed
	var profileData map[string]interface{}
	if err := json.Unmarshal(portfolio.Profile, &profileData); err == nil {
		if n, ok := profileData["name"].(string); ok && n != "" {
			portfolio.Username = strings.ToLower(strings.ReplaceAll(strings.TrimSpace(n), " ", ""))
		}
	}

	// Save
	config.DB.Save(&portfolio)
	c.JSON(http.StatusOK, portfolio)
}
