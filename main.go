package main

import (
	authController "smartdeals.rw/controller"

	"github.com/gin-gonic/gin"
)

// main function to set up the Gin router and define routes
func main() {
	r := gin.Default()

	//Route to get the user by ID
	r.GET("/api/users/:id", authController.GetUser) // Get a user

	r.Run(":8080") // Run on port 8080
}
