package repositories

import (
	"anime-manager/models"
	"gorm.io/gorm"
)

type AnimeRepository interface {
	CreateAnime(anime models.Anime) (models.Anime, error)
	GetAllAnime() ([]models.Anime, error)
	GetAnimeByID(string) (models.Anime, error)
	UpdateAnime(string, models.Anime) (models.Anime, error)
	DeleteAnime(string) error
}

type animeRepository struct {
	db *gorm.DB
}

func NewAnimeRepository(db *gorm.DB) *animeRepository {
	return &animeRepository{db}
}

func (ar *animeRepository) CreateAnime(anime models.Anime) (models.Anime, error) {
	err := ar.db.Create(&anime).Error
	return anime, err
}

func (ar *animeRepository) GetAllAnime() ([]models.Anime, error) {
	var animes []models.Anime
	err := ar.db.Find(&animes).Error
	return animes, err
}

func (ar *animeRepository) GetAnimeByID(id string) (models.Anime, error) {
	var anime models.Anime
	err := ar.db.Where("id = ?", id).First(&anime).Error
	return anime, err
}

func (ar *animeRepository) UpdateAnime(id string, anime models.Anime) (models.Anime, error) {
	err := ar.db.Where("id = ?", id).Save(&anime).Error
	return anime, err
}

func (ar *animeRepository) DeleteAnime(id string) error {
	var anime models.Anime
	err := ar.db.Where("id = ?", id).Delete(&anime).Error
	return err
}
