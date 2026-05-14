package main

import (
	authController "smartdeals.rw/controller"

	"github.com/gin-gonic/gin"
)

// main function to set up the Gin router and define routes
func main() {

	//dbInit.Create(&role)
	r := gin.Default()

	//Route to get the user by ID
	r.GET("/api/user/:id", authController.GetUser) // Get a user

	//Get all roles
	r.GET("/api/role", authController.GetAllRoles) // Get all roles

	// Route to get a role by ID
	r.GET("/api/role/:id", authController.GetRole) // Get a role

	// Route to create a new role
	r.POST("/api/role", authController.CreateRole)

	// Start the server
	r.Run(":8090") // Run on port 8080
}
