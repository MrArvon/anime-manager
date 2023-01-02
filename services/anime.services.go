package services

import (
	"anime-manager/models"
	"anime-manager/repositories"
	"github.com/google/uuid"
)

type AnimeService interface {
	CreateAnime(anime models.Anime) (models.Anime, error)
	GetAllAnime() ([]models.Anime, error)
}

type animeService struct {
	repository repositories.AnimeRepository
}

func NewAnimeService(repository repositories.AnimeRepository) *animeService {
	return &animeService{repository}
}

func (as *animeService) CreateAnime(anime models.Anime) (models.Anime, error) {
	New := models.Anime{
		ID:          uuid.New(),
		Title:       anime.Title,
		FamilyID:    anime.FamilyID,
		Type:        anime.Type,
		Description: anime.Description,
		Episode:     anime.Episode,
		Duration:    anime.Duration,
		Season:      anime.Season,
		Year:        anime.Year,
		Link:        anime.Link,
		Condition:   anime.Condition,
		Position:    anime.Position,
		Rate:        anime.Rate,
	}

	NewAnime, err := as.repository.CreateAnime(New)
	return NewAnime, err
}

func (as *animeService) GetAllAnime() ([]models.Anime, error) {
	animes, err := as.repository.GetAllAnime()
	return animes, err
}
