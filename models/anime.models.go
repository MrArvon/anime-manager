package models

import "github.com/google/uuid"

type Anime struct {
	ID          uuid.UUID `json:"id" gorm:"primaryKey" swaggerignore:"true"`
	Title       string    `json:"title" gorm:"not null"`
	FamilyID    uuid.UUID `json:"family_id" gorm:"not null"`
	Family      Family    `json:"family" gorm:"references:ID; constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" swaggerignore:"true"`
	Type        string    `json:"type"`
	Description string    `json:"description"`
	Episode     int       `json:"episode"`
	Duration    int16     `json:"duration"`
	Season      string    `json:"season"`
	Year        int16     `json:"year"`
	Link        string    `json:"link" gorm:"not null"`
	Condition   string    `json:"condition" gorm:"default:'saved'"`
	Position    string    `json:"position"`
	Rate        *int8     `json:"rate" gorm:"default:NULL"`
	CreatedAt   int64     `json:"created_at" gorm:"autoCreateTime" swaggerignore:"true"`
	UpdatedAt   int64     `json:"updated_at" gorm:"autoUpdateTime" swaggerignore:"true"`
}

type AnimeRequest struct {
	Title       string    `json:"title"`
	FamilyID    uuid.UUID `json:"family_id"`
	Type        string    `json:"type"`
	Description string    `json:"description"`
	Episode     int       `json:"episode"`
	Duration    int16     `json:"duration"`
	Season      string    `json:"season"`
	Year        int16     `json:"year"`
	Link        string    `json:"link"`
	Condition   string    `json:"condition"`
	Position    string    `json:"position"`
	Rate        *int8     `json:"rate"`
}

type AnimeResponse struct {
	ID          uuid.UUID     `json:"id" swaggerignore:"true"`
	Title       string        `json:"title"`
	Family      FamilyRequest `json:"family"`
	Type        string        `json:"type"`
	Description string        `json:"description"`
	Episode     int           `json:"episode"`
	Duration    int16         `json:"duration"`
	Season      string        `json:"season"`
	Year        int16         `json:"year"`
	Link        string        `json:"link"`
	Condition   string        `json:"condition"`
	Position    string        `json:"position"`
	Rate        *int8         `json:"rate"`
}
