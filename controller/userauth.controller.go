package controller

import (
	"github.com/gin-gonic/gin"
	"smartdeals.rw/dto"
	"smartdeals.rw/model"
	"smartdeals.rw/response"
	"smartdeals.rw/service"
	"smartdeals.rw/utils"
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

// Signup is a handler function to handle user sign-up
func SignUp(c *gin.Context) {
	// Implement logic to handle user sign-up using the AuthDTO and AuthService
	var signUpData dto.SignUpDTO

	// Instantiate the AuthService to handle the sign-up logic
	userService := service.NewUserService(utils.DBInitialize())

	// Bind the JSON request body to the signUpData struct and handle any binding errors
	if err := c.ShouldBindJSON(&signUpData); err != nil {
		c.JSON(400, response.GenericCreateResponse{
			Status:   "error",
			Code:     400,
			Message:  "Invalid request body",
			Err:      err.Error(),
			RecordID: 0,
		})
		return
	}

	// User account creation model based on the sign-up data received from the request
	encryptedPassword, err := utils.EncryptAES(signUpData.Password)

	if err != nil {
		c.JSON(500, response.GenericCreateResponse{
			Status:   "error",
			Code:     500,
			Message:  "Failed to encrypt password",
			Err:      err.Error(),
			RecordID: 0,
		})
		return
	}
	userAccountModel := model.UserAuth{
		Username: signUpData.Username,
		Password: encryptedPassword,
		Status:   "active", // Set default status to active
		RolesID:  2,        // Set default role ID (e.g., 2 for regular users)
	}

	// Call the CreateUserAccount method of the UserService to create a new user account in the database
	resp, err := userService.CreateUserAccount(&userAccountModel)
	status := "success"
	code := 200
	message := "User account created successfully"
	errMsg := ""
	recordID := 0

	// Handle the response from the CreateUserAccount method and set the appropriate response values based on success or failure
	if err != nil {
		code = 500
		status = "error"
		message = "Failed to create user account"
		errMsg = err.Error()
	} else {
		recordID = resp
	}

	// Send the JSON response back to the client with the appropriate status, code, message, error (if any), and record ID of the created user account
	c.JSON(code, response.GenericCreateResponse{
		Status:   status,
		Code:     code,
		Message:  message,
		Err:      errMsg,
		RecordID: recordID,
	})
}

// SignIn is a handler function to handle user sign-in
func SignIn(c *gin.Context) {
	// Implement logic to handle user sign-in using the AuthDTO and AuthService
	var signInDTO dto.SignInDTO

	// Instantiate the AuthService to handle the sign-in logic
	userService := service.NewUserService(utils.DBInitialize())
	// Bind the JSON request body to the signInDTO struct and handle any binding errors
	if err := c.ShouldBindJSON(&signInDTO); err != nil {
		c.JSON(400, response.GenericCreateResponse{
			Status:   "error",
			Code:     400,
			Message:  "Invalid request body",
			Err:      err.Error(),
			RecordID: 0,
		})
		return
	}
	encryptedPassword, err := utils.EncryptAES(signInDTO.Password)
	if err != nil {
		c.JSON(500, response.GenericCreateResponse{
			Status:   "error",
			Code:     500,
			Message:  "Failed to encrypt password",
			Err:      err.Error(),
			RecordID: 0,
		})
		return
	}
	// Call the SignIn method of the UserService to authenticate the user and handle the sign-in logic
	userAuthModel := &model.SignInModel{
		Username: signInDTO.Username,
		Password: encryptedPassword,
	}
	resp, err := userService.SignIn(userAuthModel)
	status := "success"
	code := 200
	message := "User signed in successfully"
	errMsg := ""
	recordID := 0
	userRole := ""
	userID := 0
	token := ""
	username := ""
	// Handle the response from the SignIn method and set the appropriate response values based on success or failure
	if err != nil {
		code = 401
		status = "error"
		message = "Invalid username or password"
		errMsg = err.Error()
	}

	if resp.ID == 0 {
		code = 404
		status = "Not Found"
		message = "User Not found"
		errMsg = "User Not Found"
		recordID = 0
		userRole = ""
		userID = 0
		token = ""
		username = ""
	}

	if resp.ID > 0 {
		code = 200
		status = "Success"
		message = "User found"
		errMsg = "User found"
		recordID = resp.ID
		userRole = resp.Role
		userID = resp.ID
		token = resp.Token
		username = resp.Username
	}

	// Send the JSON response back to the client with the appropriate status, code, message, error (if any), and any relevant data (e.g., user details, authentication token, etc.)
	c.JSON(code, response.GenericCreateResponse{
		Status:   status,
		Code:     code,
		Message:  message,
		Err:      errMsg,
		RecordID: recordID,
		UserRole: userRole,
		UserID:   userID,
		Token:    token,
		Username: username,
	})

}
