package services

import (
	"anime-manager/models"
	"anime-manager/repositories"
	"fmt"
	"github.com/google/uuid"
)

type AnimeService interface {
	CreateAnime(anime models.Anime) (models.Anime, error)
	GetAllAnime() ([]models.Anime, error)
	GetAnimeByID(string) (models.Anime, error)
	UpdateAnime(string, models.AnimeRequest) (models.Anime, error)
	DeleteAnime(string) error
}

type animeService struct {
	repository repositories.AnimeRepository
}

func NewAnimeService(repository repositories.AnimeRepository) *animeService {
	return &animeService{repository}
}

func (as *animeService) CreateAnime(anime models.Anime) (models.Anime, error) {
	//var family models.Family
	strid := anime.FamilyID.String()
	fmt.Println(strid)
	//family, _ := repositories.FamilyRepository.GetFamilyByID(strid)
	New := models.Anime{
		ID:       uuid.New(),
		Title:    anime.Title,
		FamilyID: anime.FamilyID,
		//Family:      family,
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

func (as *animeService) GetAnimeByID(id string) (models.Anime, error) {
	anime, err := as.repository.GetAnimeByID(id)
	return anime, err
}

func (as *animeService) UpdateAnime(id string, a models.AnimeRequest) (models.Anime, error) {
	var New models.Anime
	Old, err := as.repository.GetAnimeByID(id)
	if err != nil {
		return New, err
	}

	New.ID, _ = uuid.Parse(id)

	New.FamilyID = Old.FamilyID
	New.Type = Old.Type
	New.Description = Old.Description
	New.Episode = Old.Episode
	New.Duration = Old.Duration
	New.Season = Old.Season
	New.Year = Old.Year
	New.Link = Old.Link
	New.Condition = Old.Condition
	New.Position = Old.Position
	New.Rate = Old.Rate
	New.CreatedAt = Old.CreatedAt

	if a.Title != "" {
		New.Title = a.Title
	} else {
		New.Title = Old.Title
	}
	// unfinished, clear CreateAnime first
	anime, err := as.repository.UpdateAnime(id, New)
	return anime, err
}

func (as *animeService) DeleteAnime(id string) error {
	_, err := as.repository.GetAnimeByID(id)
	if err != nil {
		return err
	}
	return as.repository.DeleteAnime(id)
}
