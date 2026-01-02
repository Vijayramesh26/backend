package routes

import (
	"portfolio-admin-api/controllers"
	"portfolio-admin-api/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	api := r.Group("/api")

	api.POST("/admin/login", controllers.Login)

	admin := api.Group("/admin")
	admin.Use(middleware.AuthRequired())

	admin.POST("/portfolio", controllers.CreatePortfolio)
	admin.GET("/portfolios", controllers.GetPortfolios)
	admin.GET("/portfolio/:id", controllers.GetPortfolio)
	admin.PUT("/portfolio/:id", controllers.UpdatePortfolio)

	api.GET("/portfolio/:username", controllers.GetPublicPortfolio)
}
