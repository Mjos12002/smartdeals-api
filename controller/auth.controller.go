package controller

import (
	"fmt"
	"net/http"

	models "smartdeals.rw/model"

	"github.com/gin-gonic/gin"
)

// GetUser is a handler function to fetch a user by ID
func GetUser(c *gin.Context) {
	id := c.Param("id")
	fmt.Println("Fetching user with ID:", id)
	user := models.UserModel{
		Username: "alice",
		Password: "Alice",
		Email:    "",
		Phone:    "",
		Address:  "",
		Name:     "",
		LastName: "",
	} // Mock DB
	c.JSON(http.StatusOK, user)
}
