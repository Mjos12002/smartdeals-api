package service

import (
	"gorm.io/gorm"
	"smartdeals.rw/model"
	"smartdeals.rw/response"
	"smartdeals.rw/utils"
)

// RoleService provides methods for role-related operations
type RoleService struct {
	db *gorm.DB // Placeholder for the database connection
}

// NewRoleService creates a new instance of RoleService
func NewRoleService(db *gorm.DB) *RoleService {
	return &RoleService{db: db}
}

// CreateRole creates a new role in the database & return the created record id and any error encountered
func (s *RoleService) CreateRole(role *model.Roles) (int, error) {
	// Return the created role and any error encountered
	roleValidation := utils.ValidateRole(*role)
	if roleValidation != nil {
		return 0, roleValidation
	}
	roleCreated := s.db.Create(&role)
	if roleCreated.Error != nil {
		return 0, roleCreated.Error
	}
	return int(role.ID), nil

}

// GetRole retrieves a role by its ID from the database
func (s *RoleService) GetRole(id string) (*model.Roles, error) {
	// Return the retrieved role and any error encountered
	var role model.Roles
	if err := s.db.First(&role, id).Error; err != nil {
		return nil, err
	}
	return &role, nil
}

// Get all roles
func (s *RoleService) GetAllRoles() (response.RoleResponse, error) {
	var roles []model.Roles
	if err := s.db.Find(&roles).Error; err != nil {
		return response.RoleResponse{
			Status:  "error",
			Code:    500,
			Message: "Failed to retrieve roles",
			Err:     err.Error(),
			Data:    nil,
		}, err
	}
	return response.RoleResponse{
		Status:  "success",
		Code:    200,
		Message: "Roles retrieved successfully",
		Data:    roles,
	}, nil
}
