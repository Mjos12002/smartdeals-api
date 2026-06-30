package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"smartdeals.rw/dto"
	"smartdeals.rw/model"
	"smartdeals.rw/response"
	"smartdeals.rw/service"
	"smartdeals.rw/utils"
)

// Get profile
func GetProfile(context *gin.Context) {

	// Check the Authorization header
	var token = context.GetHeader("Authorization")
	// Initialize the profile
	if token == "" {
		context.AbortWithStatusPureJSON(http.StatusUnauthorized, response.ProfileResponse{
			Status:  "Error",
			Code:    http.StatusUnauthorized,
			Message: "Unauthorized",
		})
		return
	}

	userToken := utils.ProcessToken(token)
	userID := 0
	// Process invalid token
	if userToken.Status == "Invalid token" {
		context.JSON(http.StatusUnauthorized, response.ProfileResponse{
			Status:  "Error",
			Code:    http.StatusUnauthorized,
			Message: "Unauthorized user, contact admin",
		})
		return
	}

	// Process valid token
	userID, err := strconv.Atoi(userToken.Id)
	if err != nil {
		context.JSON(http.StatusUnauthorized, response.ProfileResponse{
			Status:  "Error",
			Code:    http.StatusUnauthorized,
			Message: "Invalid user ID",
		})
		return
	}

	profileService := service.NewProfileService(utils.DBInitialize())
	pr, err := profileService.GetProfile(userID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, pr)
		return
	}
	context.JSON(200, pr)

}

// Create profile
func CreateProfile(c *gin.Context) {

	// Get Authorization header
	var token = c.GetHeader("Authorization")
	// Initialize the profile
	var profileDTO dto.ProfileDTO
	if err := c.ShouldBindJSON(&profileDTO); err != nil {
		c.JSON(422, response.GenericCreateResponse{
			Status:   "error",
			Code:     422,
			Message:  "Invalid request body",
			Err:      err.Error(),
			RecordID: 0,
		})
		return
	}

	// Get the user id
	userToken := utils.ProcessToken(token)

	// Initialize the profile service
	profileService := service.NewProfileService(utils.DBInitialize())
	if userToken.Status != "Invalid token" {
		userID, err := strconv.Atoi(userToken.Id)
		if err != nil {
			c.JSON(500, response.GenericCreateResponse{
				Status:   "error",
				Code:     500,
				Message:  "Unknown account user",
				Err:      err.Error(),
				RecordID: 0,
			})
			return
		}
		profileModel := model.UserProfiles{
			FirstName:   profileDTO.FirstName,
			LastName:    profileDTO.LastName,
			UserAuthsId: userID,
		}
		id, err := profileService.CreateProfile(&profileModel)
		if err != nil {
			c.JSON(500, response.GenericCreateResponse{
				Status:   "error",
				Code:     500,
				Message:  "Error creating profile",
				Err:      err.Error(),
				RecordID: 0,
			})
			return
		}
		c.JSON(200, response.GenericCreateResponse{
			Status:   "error",
			Code:     200,
			Message:  "Profile created successfully",
			RecordID: id,
		})
		return
	}
	c.JSON(500, response.GenericCreateResponse{
		Status:   "error",
		Code:     500,
		Message:  "Invalid token",
		RecordID: 0,
	})

}
