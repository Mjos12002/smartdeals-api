package controller

import (
	"github.com/gin-gonic/gin"
	"smartdeals.rw/model"
	"smartdeals.rw/response"
	"smartdeals.rw/service"
	"smartdeals.rw/utils"
)

// RoleController provides handler functions for role-related operations

func CreateRole(c *gin.Context) {
	// Implement logic to create a new role using the RoleService
	// You can extract the role name from the request body and call the service method to create the role
	var roles model.Roles

	// Bind the JSON request body to the roles struct and handle any binding errors
	if err := c.ShouldBindJSON(&roles); err != nil {

		c.JSON(400, response.CreateResponse{
			Status:   "error",
			Code:     400,
			Message:  "Invalid request body",
			Err:      err.Error(),
			RecordID: 0,
		})
		return

	}
	// Create a new instance of RoleService and call the CreateRole method to create the role
	roleService := service.NewRoleService(utils.DBInitialize())

	role, err := roleService.CreateRole(&roles)
	status := "success"
	code := 200
	message := "Role created successfully"
	errMsg := ""
	recordID := 0

	if err != nil {
		code = 500
		status = "error"
		message = "Failed to create role"
		errMsg = err.Error()
	}
	recordID = int(role)

	c.JSON(code, response.CreateResponse{
		Status:   status,
		Code:     code,
		Message:  message,
		Err:      errMsg,
		RecordID: recordID,
	})
}

// GetAllRoles is a handler function to fetch all roles
func GetAllRoles(c *gin.Context) {
	// Implement logic to fetch all roles using the RoleService
	roleService := service.NewRoleService(utils.DBInitialize())
	roles, err := roleService.GetAllRoles()
	if err != nil {
		c.JSON(500, roles)
		return
	}
	c.JSON(200, roles)
}

// GetRole is a handler function to fetch a role by ID
func GetRole(c *gin.Context) {
	// Implement logic to fetch a role by ID using the RoleService
	// You can extract the role ID from the URL parameters and call the service method to get the role
	roleID := c.Param("id")
	roleService := service.NewRoleService(utils.DBInitialize())
	role, err := roleService.GetRole(roleID)
	if err != nil {
		c.JSON(404, role)
		return
	}
	c.JSON(200, role)
}
