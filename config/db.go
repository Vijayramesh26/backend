package config

import (
	"fmt"
	"os"
	"strings"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	user := strings.TrimSpace(os.Getenv("DB_USER"))
	pass := strings.TrimSpace(os.Getenv("DB_PASSWORD"))
	host := strings.TrimSpace(os.Getenv("DB_HOST"))
	port := strings.TrimSpace(os.Getenv("DB_PORT"))
	name := strings.TrimSpace(os.Getenv("DB_NAME"))
	env := strings.TrimSpace(os.Getenv("APP_ENV"))

	baseDSN := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		user, pass, host, port, name,
	)

	// Enable TLS only in production
	if env == "production" {
		baseDSN += "&tls=true"
	}

	db, err := gorm.Open(mysql.Open(baseDSN), &gorm.Config{})
	if err != nil {
		panic(fmt.Errorf("database connection failed: %w", err))
	}

	DB = db
	fmt.Println("✅ Database connected:", env)
}
