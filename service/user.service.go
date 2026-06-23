package service

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/golang-jwt/jwt/v5"
	jwtutil "github.com/kittipat1413/go-common/util/jwt"
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

// MyCustomClaims is the structure of the jwt claims
type MyCustomClaims struct {
	jwt.RegisteredClaims
	UserID   string `json:"uid"`
	UserRole string `json:"user_role"`
}

// NewUserService creates a new instance of UserService
func NewUserService(db *gorm.DB) *UserService {
	return &UserService{db: db}
}

// GetUserByID fetches a user by their ID
func (s *UserService) GetUserByID(id uint) (*model.UserProfiles, error) {
	// Implement logic to fetch user from the database using the provided ID
	// Return the user and any error encountered
	return nil, nil // Placeholder return statement
}

// CreateProfile creates a new user profile in the database
func (s *UserService) CreateProfile(user *model.UserProfiles) (int, error) {
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
	// Return the list of users and any error encountered

	status := "success"
	code := 200
	message := "Users retrieved successfully"
	errMsg := ""

	// Fetch all users from the database
	users := []model.UserProfiles{}
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

func (s *UserService) CreateUserAccount(user *model.UserAuth) (int, error) {
	// Return the created user account and any error encountered
	userValidation := utils.ValidateUserAuth(*user)
	if userValidation != nil {
		return 0, userValidation
	}
	userCreated := s.db.Create(&user)
	if userCreated.Error != nil {
		return 0, userCreated.Error
	}
	return int(user.ID), nil
}

// SignIn is a method to handle user sign-in using the provided credentials
func (s *UserService) SignIn(user *model.SignInModel) (*response.SigninResponse, error) {
	// Return a token or session information and any error encountered
	authenticatedUser := &response.UserAuths{}

	s.db.Model(&response.UserAuths{}).Preload("Roles").Where("username = ?", user.Username).Where("password = ?", user.Password).First(authenticatedUser)
	fmt.Printf("%v", authenticatedUser)
	tokenClaims, err := CreateJWT(fmt.Sprintf("%d", authenticatedUser.ID), authenticatedUser.Roles.RoleName, authenticatedUser.Username)
	if err != nil {
		return nil, err
	}
	roleName := authenticatedUser.Roles.RoleName
	userID := authenticatedUser.ID
	userName := authenticatedUser.Username
	signinResponse := response.SigninResponse{
		Token:    tokenClaims,
		Username: userName,
		ID:       int(userID),
		Role:     roleName,
	}
	return &signinResponse, nil
}

// CrateJWT creates a JWT token for the authenticated user
func CreateJWT(userID, role, username string) (string, error) {
	ctx := context.Background()
	signingKey := []byte("7b8ebceb27141aacfbc79027")
	manager, err := jwtutil.NewJWTManager(jwtutil.HS256, signingKey)
	if err != nil {
		log.Fatalf("Failed to create JWTManager: %v", err)
	}

	claims := &MyCustomClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			Issuer:    "smartdeals.rw",
			Subject:   "smartdeals-token-creator",
		},
		UserID:   userID,
		UserRole: role,
	}

	tokenStringHS256, err := manager.CreateToken(ctx, claims)

	if err != nil {
		//log.Fatalf("Failed to create token: %v", err)
		return "Failure", nil
	}

	parsedClaims := &MyCustomClaims{}
	err = manager.ParseAndValidateToken(ctx, tokenStringHS256, parsedClaims)
	if err != nil {
		//log.Fatalf("Failed to validate token: %v", err)
		return "Failure", nil
	}

	return tokenStringHS256, nil
}
