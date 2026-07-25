package service

import (
	"net/http"

	"gorm.io/gorm"
	"smartdeals.rw/model"
	"smartdeals.rw/response"
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
	profileCreated := s.db.Create(&profile)
	if profileCreated.Error != nil {
		return 0, profileCreated.Error
	}
	return int(profile.ID), nil
}

// GetProfileID is used used to get the profile details based on the user id
func (s *ProfileService) GetProfileID(userID int) (*response.ProfileResponse, error) {
	var userProfile *model.UserProfiles
	if err := s.db.Where("user_auths_id = ?", userID).First(&userProfile).Error; err != nil {
		res := &response.ProfileResponse{
			Status:  "Error",
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
			Err:     err.Error(),
		}
		return res, err
	}

	res := &response.ProfileResponse{
		Status:  "Success",
		Code:    http.StatusOK,
		Message: "Record found",
		Err:     "",
		Data: response.UserProfiles{
			ID: int(userProfile.ID),
		},
	}
	return res, nil
}

// GetProfile gets the profile of the user
func (s *ProfileService) GetProfile(userID int) (*response.ProfileResponse, error) {

	var uProfile *model.UserProfiles
	if err := s.db.Where("user_auths_id = ?", userID).First(&uProfile).Error; err != nil {
		res := &response.ProfileResponse{
			Status:  "Failed",
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
			Err:     err.Error(),
		}
		return res, err
	}
	res := &response.ProfileResponse{
		Status:  "Success",
		Code:    http.StatusAccepted,
		Message: "Record found",
		Err:     "",
		Data: response.UserProfiles{
			ID:          0,
			Firstname:   uProfile.FirstName,
			Lastname:    uProfile.LastName,
			UserAuthsId: uProfile.UserAuthsId,
		},
	}
	return res, nil
}
