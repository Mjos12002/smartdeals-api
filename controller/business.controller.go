package controller

// This file contains the controller functions for handling business-related API requests.

import (
	"github.com/gin-gonic/gin"
	"smartdeals.rw/dto"
	"smartdeals.rw/response"
	"smartdeals.rw/service"
	"smartdeals.rw/utils"
)

// CreateBusiness is a handler function to create a new business
func CreateBusiness(c *gin.Context) {

	var business dto.BusinessesDTO

	// Bind the JSON request body to the business struct and handle any binding errors
	if err := c.ShouldBindJSON(&business); err != nil {
		c.JSON(400, response.GenericCreateResponse{
			Status:   "error",
			Code:     400,
			Message:  "Invalid request body",
			Err:      err.Error(),
			RecordID: 0,
		})
		return
	}

	// Create a new instance of BusinessService and call the CreateBusiness method to create the business
	businessService := service.NewBusinessService(utils.DBInitialize())

	businessID, err := businessService.CreateBusiness(&business)
	status := "success"
	code := 200
	message := "Business created successfully"
	errMsg := ""
	recordID := 0

	if err != nil {
		code = 500
		status = "error"
		message = "Failed to create business"
		errMsg = err.Error()
	}
	recordID = businessID

	c.JSON(code, response.GenericCreateResponse{
		Status:   status,
		Code:     code,
		Message:  message,
		Err:      errMsg,
		RecordID: recordID,
	})
}

// GetAllBusinesses is a handler function to retrieve all businesses
func GetAllBusinesses(c *gin.Context) {
	businessService := service.NewBusinessService(utils.DBInitialize())
	businessResponse := businessService.GetAllBusinesses()

	c.JSON(businessResponse.Code, businessResponse)
}
