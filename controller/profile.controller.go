package controller

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"smartdeals.rw/dto"
	"smartdeals.rw/model"
	"smartdeals.rw/response"
	"smartdeals.rw/service"
	"smartdeals.rw/utils"
)

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
	fmt.Printf("token: %s\n", token)
	userToken := utils.ProcessToken(token)

	// Initialize the profile service
	profileService := service.NewProfileService(utils.DBInitialize())
	fmt.Printf("User ID: %d", userToken.Id)
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
