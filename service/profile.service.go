package service

import (
	"gorm.io/gorm"
	"smartdeals.rw/model"
)

// ProfileService provides methods to manage addresses
type ProfileService struct {
	db *gorm.DB // Placeholder for the database connection
}

// NewProfileService creates a new instance of AddressService
func NewProfileService(db *gorm.DB) *ProfileService {
	return &ProfileService{db: db}
}

// CreateAddress creates a new address in the database
func (s *ProfileService) CreateProfile(profile *model.UserProfiles) (int, error) {
	// Implement logic to create a new profile in the database using the provided address details
	// Return the created address and any error encountered

	profileCreated := s.db.Create(&profile)
	if profileCreated.Error != nil {
		return 0, profileCreated.Error
	}
	return int(profile.ID), nil
}
