package main

import (
	"fmt"

	authController "smartdeals.rw/controller"
	"smartdeals.rw/utils"

	"github.com/gin-gonic/gin"
)

// main function to set up the Gin router and define routes
func main() {

	dbInit := utils.DBInitialize()
	// role := model.Roles{
	// 	RoleName: "Super Admin",
	// }

	dbInit.Exec("INSERT INTO roles (created_at, updated_at, deleted_at, role_name) VALUES ('2026-05-12 20:11:59.84','2026-05-12 20:11:59.84',NULL,'Another Role')")

	//dbInit.Create(&role)
	fmt.Println("Database connection initialized and models migrated successfully.")
	r := gin.Default()

	//Route to get the user by ID
	r.GET("/api/users/:id", authController.GetUser) // Get a user

	r.Run(":8080") // Run on port 8080
}
