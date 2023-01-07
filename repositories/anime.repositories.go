package repositories

import (
	"anime-manager/models"
	"encoding/json"
	"fmt"
	"gorm.io/gorm"
)

type AnimeRepository interface {
	CreateAnime(anime models.Anime) (models.Anime, error)
	GetAllAnime() ([]models.Anime, error)
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
	err := ar.db.Debug().Find(&animes).Error
	//_, b := ar.db.Find(&animes).Rows()
	b, _ := json.Marshal(animes[0])
	fmt.Println(string(b))
	return animes, err
}
