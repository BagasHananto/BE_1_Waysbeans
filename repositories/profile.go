package repositories

import (
	"waysbeans/models"

	"gorm.io/gorm"
)

type ProfileRepository interface {
	FindProfile() ([]models.Profile, error)
	CreateProfile(profile models.Profile) (models.Profile, error)
	UpdateProfile(profile models.Profile) (models.Profile, error)
}

func RepositoryProfile(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindProfile() ([]models.Profile, error) {
	var profiles []models.Profile
	err := r.db.Preload("User").Find(&profiles).Error
	return profiles, err
}

func (r *repository) CreateProfile(profile models.Profile) (models.Profile, error) {
	err := r.db.Create(&profile).Error

	return profile, err
}

func (r *repository) UpdateProfile(profile models.Profile) (models.Profile, error) {
	err := r.db.Save(&profile).Error

	return profile, err
}
