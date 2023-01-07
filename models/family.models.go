package models

import "github.com/google/uuid"

type Family struct {
	ID            uuid.UUID `json:"id" gorm:"primaryKey" swaggerignore:"true"`
	Name          string    `json:"name" gorm:"not null"`
	AltName       string    `json:"alt_name"`
	AvgRate       *int8     `json:"avg_rate"`
	TotalDuration *int      `json:"total_duration"`
	CreatedAt     int64     `json:"created_at" gorm:"autoCreateTime" swaggerignore:"true"`
	UpdatedAt     int64     `json:"updated_at" gorm:"autoUpdateTime" swaggerignore:"true"`
}

type FamilyRequest struct {
	Name    string `json:"name" gorm:"not null"`
	AltName string `json:"alt_name"`
}

type FamilyResponse struct {
	Name          string `json:"name" gorm:"not null"`
	AltName       string `json:"alt_name"`
	AvgRate       *int8  `json:"avg_rate" gorm:"default:0"`
	TotalDuration *int   `json:"total_duration" gorm:"default:0"`
}
