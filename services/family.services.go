package services

import (
	"anime-manager/models"
	"anime-manager/repositories"
	"github.com/google/uuid"
)

type FamilyService interface {
	CreateFamily(family models.Family) (models.Family, error)
	GetAllFamily() ([]models.Family, error)
	GetFamilyByID(string) (models.Family, error)
	UpdateFamily(string, models.FamilyRequest) (models.Family, error)
	DeleteFamily(string) error
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
	family, err := fs.repository.GetFamilyByID(id)
	return family, err
}

func (fs *familyService) UpdateFamily(id string, f models.FamilyRequest) (models.Family, error) {
	var New models.Family
	Old, err := fs.repository.GetFamilyByID(id)
	if err != nil {
		return New, err
	}

	New.ID, _ = uuid.Parse(id)
	New.Name = Old.Name
	New.AltName = Old.AltName
	New.AvgRate = Old.AvgRate
	New.TotalDuration = Old.TotalDuration
	New.CreatedAt = Old.CreatedAt

	if f.Name != "" {
		New.Name = f.Name
	}
	if f.AltName != "" {
		New.AltName = f.AltName
	}
	if f.AvgRate != nil {
		New.AvgRate = f.AvgRate
	}
	if f.TotalDuration != nil {
		New.TotalDuration = f.TotalDuration
	}
	family, err := fs.repository.UpdateFamily(id, New)
	return family, err
}

func (fs *familyService) DeleteFamily(id string) error {
	_, err := fs.repository.GetFamilyByID(id)
	if err != nil {
		return err
	}
	return fs.repository.DeleteFamily(id)
}
