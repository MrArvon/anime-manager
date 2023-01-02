package models

import "github.com/google/uuid"

type Anime struct {
	ID          uuid.UUID `gorm:"primaryKey" swaggerignore:"true"`
	Title       string    `gorm:"not null"`
	FamilyID    uuid.UUID `gorm:"not null" json:"family_id"`
	Family      Family    `gorm:"foreignKey:FamilyID" swaggerignore:"true"`
	Type        string
	Description string
	Episode     int
	Duration    int16
	Season      string
	Year        int16
	Link        string  `gorm:"not null"`
	Condition   *string `gorm:"default:'recorded'"`
	Position    string
	Rate        *int8 `gorm:"default:NULL"`
	CreatedAt   int64 `gorm:"autoCreateTime" swaggerignore:"true"`
	UpdatedAt   int64 `gorm:"autoUpdateTime" swaggerignore:"true"`
}

type Family struct {
	ID            uuid.UUID `gorm:"primaryKey" swaggerignore:"true"`
	Name          string    `gorm:"not null"`
	AltName       string
	AvgRate       *int8
	TotalDuration *int
	CreatedAt     int64 `gorm:"autoCreateTime" swaggerignore:"true"`
	UpdatedAt     int64 `gorm:"autoUpdateTime" swaggerignore:"true"`
}
