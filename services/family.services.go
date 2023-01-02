package services

import (
	"anime-manager/models"
	"anime-manager/repositories"
	"github.com/google/uuid"
)

type FamilyService interface {
	CreateFamily(family models.Family) (models.Family, error)
	GetAllFamily() ([]models.Family, error)
	GetFamilyByID(id string) (models.Family, error)
	//UpdateFamily(id string) (models.Family, error)
	DeleteFamily(id string) error
}

type familyService struct {
	repository repositories.FamilyRepository
}

func NewFamilyService(repository repositories.FamilyRepository) *familyService {
	return &familyService{repository}
}

func (fs *familyService) CreateFamily(family models.Family) (models.Family, error) {
	New := models.Family{
		ID:            uuid.New(),
		Name:          family.Name,
		AltName:       family.AltName,
		AvgRate:       family.AvgRate,
		TotalDuration: family.TotalDuration,
	}

	NewFamily, err := fs.repository.CreateFamily(New)
	return NewFamily, err
}

func (fs *familyService) GetAllFamily() ([]models.Family, error) {
	families, err := fs.repository.GetAllFamily()
	return families, err
}

func (fs *familyService) GetFamilyByID(id string) (models.Family, error) {
	families, err := fs.repository.GetFamilyByID(id)
	return families, err
}

func (fs *familyService) DeleteFamily(id string) error {
	_, err := fs.repository.GetFamilyByID(id)
	if err != nil {
		return err
	}
	return fs.repository.DeleteFamily(id)
}
