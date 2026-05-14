package controller

// This file contains handler functions related to user operations

import (
	"github.com/gin-gonic/gin"
	"smartdeals.rw/model"
	"smartdeals.rw/response"
	"smartdeals.rw/service"
	"smartdeals.rw/utils"
)

// CreateUser is a handler function to create a new user
func CreateUser(c *gin.Context) {
	var userDetails model.UserDetails

	// Bind the JSON request body to the userDetails struct and handle any binding errors
	if err := c.ShouldBindJSON(&userDetails); err != nil {
		c.JSON(400, response.CreateResponse{
			Status:   "error",
			Code:     400,
			Message:  "Invalid request body",
			Err:      err.Error(),
			RecordID: 0,
		})
		return
	}

	// Create a new instance of UserService and call the CreateUser method to create the user
	userService := service.NewUserService(utils.DBInitialize())

	userID, err := userService.CreateUser(&userDetails)
	status := "success"
	code := 200
	message := "User created successfully"
	errMsg := ""
	recordID := 0

	if err != nil {
		code = 500
		status = "error"
		message = "Failed to create user"
		errMsg = err.Error()
	}
	recordID = userID

	c.JSON(code, response.CreateResponse{
		Status:   status,
		Code:     code,
		Message:  message,
		Err:      errMsg,
		RecordID: recordID,
	})
}

func GetAllUsers(c *gin.Context) {
	userService := service.NewUserService(utils.DBInitialize())
	userResponses := userService.GetAllUsers()

	c.JSON(userResponses.Code, userResponses)
}
