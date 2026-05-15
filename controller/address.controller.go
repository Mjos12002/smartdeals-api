package controller

// This file contains the controller functions for handling address-related API requests.

import (
	"github.com/gin-gonic/gin"
	"smartdeals.rw/dto"
	"smartdeals.rw/model"
	"smartdeals.rw/response"
	"smartdeals.rw/service"
	"smartdeals.rw/utils"
)

// CreateAddress is a handler function to create a new address
func CreateAddress(c *gin.Context) {
	var address dto.AddressDTO

	// Bind the JSON request body to the address struct and handle any binding errors
	if err := c.ShouldBindJSON(&address); err != nil {
		c.JSON(400, response.GenericCreateResponse{
			Status:   "error",
			Code:     400,
			Message:  "Invalid request body",
			Err:      err.Error(),
			RecordID: 0,
		})
		return
	}

	// Create a new instance of AddressService and call the CreateAddress method to create the address
	addressService := service.NewAddressService(utils.DBInitialize())

	// Generate the address model from the DTO and call the CreateAddress method to create the address in the database
	addressModel := model.Addresses{
		Street:      address.Street,
		PopularName: address.PopularName,
		Province:    address.Province,
		District:    address.District,
		Sector:      address.Sector,
		LongLat:     address.LongLat,
	}
	addressID, err := addressService.CreateAddress(&addressModel)
	status := "success"
	code := 200
	message := "Address created successfully"
	errMsg := ""
	recordID := 0

	if err != nil {
		code = 500
		status = "error"
		message = "Failed to create address"
		errMsg = err.Error()
	}
	recordID = addressID

	c.JSON(code, response.GenericCreateResponse{
		Status:   status,
		Code:     code,
		Message:  message,
		Err:      errMsg,
		RecordID: recordID,
	})
}

// GetAllAddresses is a handler function to retrieve all addresses
func GetAllAddresses(c *gin.Context) {
	addressService := service.NewAddressService(utils.DBInitialize())
	addressResponse, err := addressService.GetAllAddresses()

	if err != nil {
		c.JSON(500, response.GenericCreateResponse{
			Status:   "error",
			Code:     500,
			Message:  "Failed to retrieve addresses",
			Err:      err.Error(),
			RecordID: 0,
		})
		return
	}

	c.JSON(addressResponse.Code, addressResponse)
}
