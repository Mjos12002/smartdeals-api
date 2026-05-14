package controller

import (
	"github.com/gin-gonic/gin"
)

// GetUser is a handler function to fetch a user by ID
func GetUser(c *gin.Context) {

	// validate := validator.New()

	// dbInit := utils.DBInitialize()
	// role := model.Roles{
	// 	RoleName: "Su",
	// }

	// err := validate.Struct(role)
	// if err != nil {
	// 	response := response.CreateResponse{
	// 		Status:  "error",
	// 		Code:    http.StatusBadRequest,
	// 		Message: "Validation failed",
	// 		Err:     err.Error(),
	// 	}
	// 	c.JSON(http.StatusBadRequest, gin.H{"response": response})
	// 	return
	// }

	// if err := dbInit.Create(&role).Error; err != nil {
	// 	response := response.CreateResponse{
	// 		Status:  "error",
	// 		Code:    http.StatusBadRequest,
	// 		Message: "Failed to create role",
	// 		Err:     err.Error(),
	// 	}
	// 	c.JSON(http.StatusBadRequest, gin.H{"response": response})
	// 	return
	// }

	// dbInit.Create(&role)
	// response := response.CreateResponse{
	// 	Status:   "success",
	// 	Code:     http.StatusOK,
	// 	Message:  "User fetched successfully",
	// 	RecordID: role.ID,
	// }
	// c.JSON(http.StatusOK, gin.H{"response": response})

}
