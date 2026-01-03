package main

import (
	"os"
	"portfolio-admin-api/config"
	"portfolio-admin-api/controllers"
	"portfolio-admin-api/models"
	"portfolio-admin-api/routes"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load env only for local
	if os.Getenv("APP_ENV") != "production" {
		_ = godotenv.Load(".env.local")
	}
	// Connect DB
	config.ConnectDB()

	// Auto migrate tables
	config.DB.AutoMigrate(
		&models.User{},
		&models.Portfolio{},
	)

	r := gin.Default()

	// CORS
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	// Health check (Render needs this)
	r.GET("/health", controllers.HealthCheck)
	routes.SetupRoutes(r)

	r.Run(":8080")
}
