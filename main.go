package main

import (
	authController "smartdeals.rw/controller"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/api/users/:id", authController.GetUser) // Get a user

	r.Run(":8080") // Run on port 8080
}
