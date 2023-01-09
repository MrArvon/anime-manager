package repositories

import (
	"anime-manager/models"
	"gorm.io/gorm"
)

type FamilyRepository interface {
	CreateFamily(family models.Family) (models.Family, error)
	GetAllFamily() ([]models.Family, error)
	GetFamilyByID(string) (models.Family, error)
	UpdateFamily(string, models.Family) (models.Family, error)
	DeleteFamily(string) error
}

type familyRepository struct {
	db *gorm.DB
}

func NewFamilyRepository(db *gorm.DB) *familyRepository {
	return &familyRepository{db}
}

func (fr *familyRepository) CreateFamily(family models.Family) (models.Family, error) {
	err := fr.db.Create(&family).Error
	return family, err
}

func (fr *familyRepository) GetAllFamily() ([]models.Family, error) {
	var families []models.Family
	err := fr.db.Find(&families).Error
	return families, err
}

func (fr *familyRepository) GetFamilyByID(id string) (models.Family, error) {
	var family models.Family
	err := fr.db.Where("id = ?", id).First(&family).Error
	return family, err
}

func (fr *familyRepository) UpdateFamily(id string, family models.Family) (models.Family, error) {
	err := fr.db.Where("id = ?", id).Save(&family).Error
	return family, err
}

func (fr *familyRepository) DeleteFamily(id string) error {
	var family models.Family
	err := fr.db.Where("id = ?", id).Delete(&family).Error
	return err
}
