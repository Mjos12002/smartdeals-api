package controller

// This file contains the controller functions for handling business-related API requests.

import (
	"net/http"
	"strconv"

	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"smartdeals.rw/dto"
	"smartdeals.rw/response"
	"smartdeals.rw/service"
	"smartdeals.rw/utils"
)

// CreateBusiness is a handler function to create a new business
func CreateBusiness(context *gin.Context) {

	// Process the authorization jwt token
	token := context.GetHeader("Authorization")
	processedToken := utils.ProcessToken(token)

	if processedToken.Id == "" {
		context.JSON(http.StatusBadRequest, response.GenericCreateResponse{
			Status:   "Error",
			Code:     http.StatusUnauthorized,
			Message:  "Unauthorized, contact admin",
			Err:      "Unauthorized, contact admin",
			RecordID: 0,
		})
		return
	}

	uID, _ := strconv.Atoi(processedToken.Id)

	var business dto.BusinessesDTO
	name := context.Request.FormValue("name")
	description := context.Request.FormValue("description")
	logoFile, logoFileHeader, err := context.Request.FormFile("logo_url")

	// Processing the logo file
	if err != nil {
		context.JSON(http.StatusBadRequest, response.GenericCreateResponse{
			Status:   "Error",
			Code:     http.StatusInternalServerError,
			Message:  err.Error(),
			Err:      err.Error(),
			RecordID: 0,
		})
		return
	}
	defer logoFile.Close()

	logoFileExt := filepath.Ext(logoFileHeader.Filename)
	newFilename := uuid.New().String() + logoFileExt

	logoFileDest := filepath.Join("./resources", newFilename)

	// Get user profile id
	userAuthService := service.NewProfileService(utils.DBInitialize())

	profileID, err := userAuthService.GetProfileID(uID)

	// Get the profile
	if err != nil {
		context.JSON(http.StatusBadRequest, response.GenericCreateResponse{
			Status:   "Error",
			Code:     http.StatusNoContent,
			Message:  err.Error(),
			Err:      err.Error(),
			RecordID: 0,
		})
		return
	}

	business.Name = name
	business.Address = 1
	business.Description = description
	business.LogoURL = logoFileDest
	business.UserProfile = profileID.Data.ID

	if err := context.SaveUploadedFile(logoFileHeader, logoFileDest); err != nil {
		context.JSON(http.StatusInternalServerError, response.GenericCreateResponse{
			Status:   "Error",
			Code:     http.StatusInternalServerError,
			Message:  err.Error(),
			Err:      err.Error(),
			RecordID: 0,
		})
		return
	}

	// Create a new instance of BusinessService and call the CreateBusiness method to create the business
	businessService := service.NewBusinessService(utils.DBInitialize())

	businessID, err := businessService.CreateBusiness(&business)
	status := "success"
	code := http.StatusCreated
	message := "Business created successfully"
	errMsg := ""
	recordID := 0

	if err != nil {
		code = http.StatusInternalServerError
		status = "error"
		message = "Failed to create business"
		errMsg = err.Error()
	}
	recordID = businessID

	context.JSON(http.StatusBadRequest, response.GenericCreateResponse{
		Status:   status,
		Code:     code,
		Message:  message,
		Err:      errMsg,
		RecordID: recordID,
	})
}

// GetAllBusinesses is a handler function to retrieve all businesses
func GetAllBusinesses(context *gin.Context) {

	token := context.GetHeader("Authorization")
	processedToken := utils.ProcessToken(token)

	if processedToken.Id == "" || processedToken.Status == "Invalid token" {
		context.JSON(http.StatusUnauthorized, response.GenericCreateResponse{
			Status:   "Error",
			Code:     http.StatusUnauthorized,
			Message:  "Unauthorized, contact admin",
			Err:      "Unauthorized, contact admin",
			RecordID: 0,
		})
		return
	}
	uID, _ := strconv.Atoi(processedToken.Id)

	// Get profile
	userAuthService := service.NewProfileService(utils.DBInitialize())
	profileID, err := userAuthService.GetProfileID(uID)
	if err != nil {
		context.JSON(http.StatusBadRequest, response.GenericCreateResponse{
			Status:   "Error",
			Code:     http.StatusNoContent,
			Message:  err.Error(),
			Err:      err.Error(),
			RecordID: 0,
		})
		return
	}

	businessService := service.NewBusinessService(utils.DBInitialize())

	businessResponse := businessService.GetAllBusinesses(profileID.Data.ID)

	context.JSON(businessResponse.Code, businessResponse)
}
