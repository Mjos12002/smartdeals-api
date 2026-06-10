package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	restcontroller "smartdeals.rw/controller"
)

// main function to set up the Gin router and define routes
func main() {

	//dbInit.Create(&role)
	r := gin.Default()
	r.Use(cors.Default())

	//Route to get the user by ID
	r.GET("/api/v1/user/:id", restcontroller.GetUser) // Get a user

	//Get all roles
	r.GET("/api/v1/role", restcontroller.GetAllRoles) // Get all roles

	// Route to get a role by ID
	r.GET("/api/v1/role/:id", restcontroller.GetRole) // Get a role

	// Route to create a new role
	r.POST("/api/v1/role", restcontroller.CreateRole)

	r.POST("/api/v1/user", restcontroller.CreateUser) // Create a new user

	// Route to get all users
	r.GET("/api/v1/user", restcontroller.GetAllUsers) // Get all users

	// Route to handle user sign-up
	r.POST("/api/v1/signup", restcontroller.SignUp) // User sign-up

	// Route to handle user signin
	r.POST("/api/v1/signin", restcontroller.SignIn) // User sign-in

	// Route to create a new address
	r.POST("/api/v1/address", restcontroller.CreateAddress) // Create a new address

	// Route to get all addresses
	r.GET("/api/v1/address", restcontroller.GetAllAddresses) // Get all addresses

	// Route to create a new business
	r.POST("/api/v1/business", restcontroller.CreateBusiness) // Create a new business

	// Route to get all businesses
	r.GET("/api/v1/business", restcontroller.GetAllBusinesses) // Get all businesses

	// Route to create a new product
	r.POST("/api/v1/product", restcontroller.CreateProduct) // Create a new product

	// Route to get all products
	r.GET("/api/v1/product", restcontroller.GetAllProducts) // Get all products``

	// Route to create a new product category
	r.POST("/api/v1/product-category", restcontroller.CreateProductCategory)

	// Route to get all product categories
	r.GET("/api/v1/product-category", restcontroller.GetAllProductCategories)

	// Start the server on a specific port
	r.Run(":8090") // Run on port 8090
}
