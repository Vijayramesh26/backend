package models

import (
	"time"

	"gorm.io/datatypes"
)

type Portfolio struct {
	ID         uint           `gorm:"primaryKey" json:"id"`
	UserID     uint           `gorm:"not null" json:"user_id"`
	Username   string         `gorm:"unique;not null" json:"username"`
	Profile    datatypes.JSON `gorm:"type:json" json:"profile"`
	About      datatypes.JSON `gorm:"type:json" json:"about"`
	Skills     datatypes.JSON `gorm:"type:json" json:"skills"`
	Projects   datatypes.JSON `gorm:"type:json" json:"projects"`
	Experience datatypes.JSON `gorm:"type:json" json:"experience"`
	Contact    datatypes.JSON `gorm:"type:json" json:"contact"`
	Published  bool           `gorm:"default:false" json:"published"` // admin can publish
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
}

type Skill struct {
	ID          uint `gorm:"primaryKey"`
	PortfolioID uint
	Name        string
	ImageURL    string
	Experience  string
}
