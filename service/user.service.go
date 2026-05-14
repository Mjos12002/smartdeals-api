package service

import (
	"gorm.io/gorm"
	"smartdeals.rw/model"
	"smartdeals.rw/response"
	"smartdeals.rw/utils"
)

// UserService provides methods for user-related operations
type UserService struct {
	// You can add dependencies here, such as a database connection
	db *gorm.DB // Placeholder for the database connection
}

// NewUserService creates a new instance of UserService
func NewUserService(db *gorm.DB) *UserService {
	return &UserService{db: db}
}

// GetUserByID fetches a user by their ID
func (s *UserService) GetUserByID(id uint) (*model.UserDetails, error) {
	// Implement logic to fetch user from the database using the provided ID
	// Return the user and any error encountered
	return nil, nil // Placeholder return statement
}

// CreateUser creates a new user in the database
func (s *UserService) CreateUser(user *model.UserDetails) (int, error) {
	// Implement logic to create a new user in the database using the provided user details
	// Return the created user and any error encountered
	userValidation := utils.ValidateUserDetails(*user)
	if userValidation != nil {
		return 0, userValidation
	}
	userCreated := s.db.Create(&user)
	if userCreated.Error != nil {
		return 0, userCreated.Error
	}
	return int(user.ID), nil
}

func (s *UserService) GetAllUsers() response.UserResponse {
	// Implement logic to fetch all users from the database
	// Return the list of users and any error encountered

	status := "success"
	code := 200
	message := "Users retrieved successfully"
	errMsg := ""

	// Fetch all users from the database
	users := []model.UserDetails{}
	result := s.db.Find(&users)

	if result.Error != nil {
		status = "error"
		code = 500
		message = "Failed to retrieve users"
		errMsg = result.Error.Error()
		users = nil
	}
	userResponses := response.UserResponse{
		Status:  status,
		Code:    code,
		Message: message,
		Err:     errMsg,
		Data:    users,
	}

	return userResponses
}
