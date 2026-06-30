package service

import (
	"net/http"

	"gorm.io/gorm"
	"smartdeals.rw/model"
	"smartdeals.rw/response"
)

// AddressService provides methods to manage addresses
type AddressService struct {
	db *gorm.DB // Placeholder for the database connection
}

// NewAddressService creates a new instance of AddressService
func NewAddressService(db *gorm.DB) *AddressService {
	return &AddressService{db: db}
}

// CreateAddress creates a new address in the database
func (s *AddressService) CreateAddress(address *model.Addresses) (int, error) {
	// Implement logic to create a new address in the database using the provided address details
	// Return the created address and any error encountered

	addressCreated := s.db.Create(&address)
	if addressCreated.Error != nil {
		return 0, addressCreated.Error
	}
	return int(address.ID), nil
}

// GetAllAddresses retrieves all addresses from the database
func (s *AddressService) GetAllAddresses() (response.AddressResponse, error) {
	// Implement logic to fetch all addresses from the database
	// Return the list of addresses and any error encountered

	status := "success"
	code := http.StatusOK
	message := "Addresses retrieved successfully"
	errMsg := ""

	addresses := []model.Addresses{}
	result := s.db.Find(&addresses)
	if result.Error != nil {
		status = "error"
		code = http.StatusInternalServerError
		message = result.Error.Error()
		errMsg = result.Error.Error()
	}

	return response.AddressResponse{
		Status:  status,
		Code:    code,
		Message: message,
		Err:     errMsg,
		Data:    addresses,
	}, nil
}
