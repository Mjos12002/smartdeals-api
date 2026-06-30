package main

import (
	"github.com/gin-gonic/gin"
	restcontroller "smartdeals.rw/controller"
)

// main function to set up the Gin router and define routes
func main() {

	//dbInit.Create(&role)
	router := gin.Default()
	router.Use(CORSMiddleware())

	secure := router.Group("/api/secure")
	{
		//Route to get the user by ID
		secure.GET("/v1/user/:id", restcontroller.GetUser) // Get a user

		//Get all roles
		secure.GET("/v1/role", restcontroller.GetAllRoles) // Get all roles

		// Route to get a role by ID
		secure.GET("/v1/role/:id", restcontroller.GetRole) // Get a role

		// Route to create a new role
		secure.POST("/v1/role", restcontroller.CreateRole) // Create role

		secure.POST("/v1/user", restcontroller.CreateUser) // Create a new user

		// Route to get all users
		secure.GET("/v1/user", restcontroller.GetAllUsers) // Get all users
		// Route to handle user sign-up
		secure.POST("/v1/signup", restcontroller.SignUp) // User sign-up

		// Route to handle user signin
		secure.POST("/v1/signin", restcontroller.SignIn) // User sign-in

		// Route to create a new address
		secure.POST("/v1/address", restcontroller.CreateAddress) // Create a new address

		// Route to get all addresses
		secure.GET("/v1/address", restcontroller.GetAllAddresses) // Get all addresses

		// Route to create a new business
		secure.POST("/v1/business", restcontroller.CreateBusiness) // Create a new business

		// Route to get all businesses
		secure.GET("/v1/business", restcontroller.GetAllBusinesses) // Get all businesses

		// Route to create a new product
		secure.POST("/v1/product", restcontroller.CreateProduct) // Create a new product

		// Route to get all products
		secure.GET("/v1/product", restcontroller.GetAllProducts) // Get all products``

		// Route to create a new product category
		secure.POST("/v1/product-category", restcontroller.CreateProductCategory) // Create product category

		// Route to get all product categories
		secure.GET("/v1/product-category", restcontroller.GetAllProductCategories) // Get all product categoris

		// Route to create profile
		secure.POST("/v1/profile", restcontroller.CreateProfile) // Create profile

		secure.GET("/v1/profile", restcontroller.GetProfile)

	}
	// Start the server on a specific port
	router.Run(":8090") // Run on port 8090
}

// CORSMiddleware is used to allow CORS
func CORSMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {

		context.Header("Access-Control-Allow-Origin", "*")
		context.Header("Access-Control-Allow-Credentials", "true")
		context.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		context.Header("Access-Control-Allow-Methods", "POST,HEAD,PATCH, OPTIONS, GET, PUT")

		if context.Request.Method == "OPTIONS" {
			context.AbortWithStatus(204)
			return
		}

		context.Next()
	}
}

// SecureAPIMiddleWare is a middleware used to secure APIs by requiring the Authentication token
func SecureAPIMiddleWare() gin.HandlerFunc {
	return func(context *gin.Context) {

		context.Next()
	}
}
