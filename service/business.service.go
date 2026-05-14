package service

import (
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
	// Implement logic to create a new business in the database using the provided business details
	// Return the created business and any error encountered
	// validate the business data before creating it in the database
	err := utils.ValidateBusiness(*business)
	if err != nil {
		return 0, err
	}
	businessModel := model.Businesses{
		Name:          business.Name,
		Description:   business.Description,
		Contact:       business.Contact,
		Phone:         business.Phone,
		LogoURL:       business.LogoURL,
		AddressID:     business.Address,
		UserDetailsID: business.UserDetails,
	}
	result := s.db.Create(&businessModel)
	if result.Error != nil {
		return 0, result.Error
	}
	return (int(businessModel.ID)), nil // Placeholder return value
}

// GetAllBusinesses retrieves all businesses from the database
func (s *BusinessService) GetAllBusinesses() response.BusinessResponse {
	// Implement logic to fetch all businesses from the database
	// Return the list of businesses and any error encountered

	status := "success"
	code := 200
	message := "Businesses retrieved successfully"
	errMsg := ""

	var businesses []model.Businesses
	result := s.db.Find(&businesses)
	if result.Error != nil {
		status = "error"
		code = 500
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
