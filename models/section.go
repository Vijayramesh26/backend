package models

import (
	"time"
)

type Section struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	PortfolioID uint      `gorm:"not null" json:"portfolio_id"`
	Name        string    `gorm:"not null" json:"name"` // e.g., Profile, Skills, Projects
	Order       int       `json:"order"`                // display order
	Enabled     bool      `gorm:"default:true" json:"enabled"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
