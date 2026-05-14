package main

import (
	"github.com/gin-gonic/gin"
	restcontroller "smartdeals.rw/controller"
)

// main function to set up the Gin router and define routes
func main() {

	//dbInit.Create(&role)
	r := gin.Default()

	//Route to get the user by ID
	r.GET("/api/user/:id", restcontroller.GetUser) // Get a user

	//Get all roles
	r.GET("/api/role", restcontroller.GetAllRoles) // Get all roles

	// Route to get a role by ID
	r.GET("/api/role/:id", restcontroller.GetRole) // Get a role

	// Route to create a new role
	r.POST("/api/role", restcontroller.CreateRole)

	r.POST("/api/user", restcontroller.CreateUser) // Create a new user

	// Route to get all users
	r.GET("/api/user", restcontroller.GetAllUsers) // Get all users

	// Route to create a new address
	r.POST("/api/address", restcontroller.CreateAddress) // Create a new address

	// Route to get all addresses
	r.GET("/api/address", restcontroller.GetAllAddresses) // Get all addresses

	// Route to create a new business
	r.POST("/api/business", restcontroller.CreateBusiness) // Create a new business

	// Route to get all businesses
	r.GET("/api/business", restcontroller.GetAllBusinesses) // Get all businesses

	// Start the server
	r.Run(":8090") // Run on port 8090
}
