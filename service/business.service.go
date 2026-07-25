package service

import (
	"net/http"

	"gorm.io/gorm"
	"smartdeals.rw/dto"
	"smartdeals.rw/model"
	"smartdeals.rw/response"
	"smartdeals.rw/utils"
)

// BusinessService provides methods to manage businesses
type BusinessService struct {
	// You can add dependencies here, such as a database connection
	db *gorm.DB // Placeholder for the database connection
}

// NewBusinessService creates a new instance of BusinessService
func NewBusinessService(db *gorm.DB) *BusinessService {
	return &BusinessService{db: db}
}

// CreateBusiness creates a new business in the database
func (s *BusinessService) CreateBusiness(business *dto.BusinessesDTO) (int, error) {
	// validate the business data before creating it in the database
	err := utils.ValidateBusiness(*business)
	if err != nil {
		return 0, err
	}
	businessModel := model.Businesses{
		Name:           business.Name,
		Description:    business.Description,
		LogoURL:        business.LogoURL,
		Street:         business.Street,
		PopularName:    business.PopularName,
		Email:          business.Email,
		PhoneNumber:    business.PhoneNumber,
		Twitter:        business.Twitter,
		Facebook:       business.Facebook,
		Instagram:      business.Instagram,
		Province:       business.Province,
		District:       business.District,
		UserProfilesID: business.UserProfile,
	}
	result := s.db.Create(&businessModel)
	if result.Error != nil {
		return 0, result.Error
	}
	return (int(businessModel.ID)), nil // Placeholder return value
}

// GetAllBusinesses retrieves all businesses from the database
func (s *BusinessService) GetAllBusinesses(profileID int) response.BusinessResponse {
	// Implement logic to fetch all businesses from the database
	// Return the list of businesses and any error encountered

	status := "success"
	code := 200
	message := "Businesses retrieved successfully"
	errMsg := ""

	var businesses []model.Businesses
	result := s.db.Where("user_profiles_id = ?", profileID).Find(&businesses)
	if result.Error != nil {
		status = "error"
		code = http.StatusInternalServerError
		message = "Failed to retrieve businesses"
		errMsg = result.Error.Error()
	}

	return response.BusinessResponse{
		Status:  status,
		Code:    code,
		Message: message,
		Err:     errMsg,
		Data:    businesses,
	} // Placeholder return value
}

// GetUserBusiness is used to load business of users
func (s *BusinessService) GetUserBusiness(userID int) response.BusinessResponse {
	var businesses []model.Businesses
	result := s.db.Joins("JOIN user_profiles ON user_profiles.id = businesses.user_profiles_id").Joins("JOIN user_auths ON user_auths.id = user_profiles.user_auths_id AND user_auths.id = ?", userID).Find(&businesses)
	if result.Error != nil {
		return response.BusinessResponse{
			Status:  "Error",
			Code:    http.StatusInternalServerError,
			Message: result.Error.Error(),
			Err:     result.Error.Error(),
			Data:    businesses,
		}
	}
	return response.BusinessResponse{
		Status:  "Success",
		Code:    http.StatusAccepted,
		Message: "Successfully returned data",
		Err:     "",
		Data:    businesses,
	}
}
