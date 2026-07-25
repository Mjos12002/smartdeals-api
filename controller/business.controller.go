package controller

// This file contains the controller functions for handling business-related API requests.
import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

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
	fmt.Printf("%s", processedToken)

	if processedToken.Status == "Invalid token" {
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
	street := context.Request.FormValue("street")
	popular_name := context.Request.FormValue("popular_name")
	email := context.Request.FormValue("email")
	phone_number := context.Request.FormValue("phone_number")
	twitter := context.Request.FormValue("twitter")
	facebook := context.Request.FormValue("facebook")
	instagram := context.Request.FormValue("instagram")
	province := context.Request.FormValue("province")
	district := context.Request.FormValue("district")
	fileURL := []string{}

	//logoFile, logoFileHeader, err := context.Request.FormFile("logo_url")

	multiPForm, err := context.MultipartForm()

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

	// Get the files uploaded
	files := multiPForm.File["logo_url"]

	for _, file := range files {

		logoFileExt := filepath.Ext(file.Filename)
		newFilename := uuid.New().String() + logoFileExt

		logoFileDest := filepath.Join("./resources", newFilename)
		if err := context.SaveUploadedFile(file, logoFileDest); err != nil {
			context.JSON(http.StatusInternalServerError, response.GenericCreateResponse{
				Status:   "Error",
				Code:     http.StatusInternalServerError,
				Message:  err.Error(),
				Err:      err.Error(),
				RecordID: 0,
			})
			return
		}
		fileURL = append(fileURL, logoFileDest)

	}

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
	business.Street = street
	business.Description = description
	business.LogoURL = strings.Join(fileURL, ", ")
	business.PopularName = phone_number
	business.Email = email
	business.PhoneNumber = phone_number
	business.Twitter = twitter
	business.Facebook = facebook
	business.Instagram = instagram
	business.Province = province
	business.District = district
	business.UserProfile = profileID.Data.ID
	business.PopularName = popular_name

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
