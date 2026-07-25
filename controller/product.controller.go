package controller

// This file contains the controller functions for handling Product-related HTTP requests.

import (
	"net/http"
	"strconv"

	"smartdeals.rw/dto"
	"smartdeals.rw/response"
	"smartdeals.rw/service"
	"smartdeals.rw/utils"

	"github.com/gin-gonic/gin"
)

// CreateProduct handles the creation of a new product
func CreateProduct(context *gin.Context) {

	// Get user id from the token
	token := context.GetHeader("Authorization")
	processedToken := utils.ProcessToken(token)

	if processedToken.Status == "Invalid token" {
		context.JSON(http.StatusBadRequest, response.GenericCreateResponse{
			Status:   "Error",
			Code:     http.StatusUnauthorized,
			Message:  "Unauthorized, contact admin",
			Err:      "Unauthorized, contact admin",
			RecordID: 0,
		})
		return
	}

	userID, _ := strconv.Atoi(processedToken.Id)
	// End getting user id from the token
	var productDTO dto.ProductDTO

	name := context.Request.FormValue("name")
	description := context.Request.FormValue("description")

	price, err := strconv.Atoi(context.Request.FormValue("price"))
	// Check the conversion of the price to integer type
	if err != nil {
		context.JSON(http.StatusBadRequest, response.GenericCreateResponse{
			Status:   "error",
			Code:     http.StatusBadRequest,
			Message:  "Invalid price value",
			Err:      err.Error(),
			RecordID: 0,
		})
		return
	}

	product_categories_id, err := strconv.Atoi(context.Request.FormValue("product_categories_id"))
	if err != nil {
		context.JSON(http.StatusBadRequest, response.GenericCreateResponse{
			Status:   "error",
			Code:     http.StatusBadRequest,
			Message:  "Invalid product category",
			Err:      err.Error(),
			RecordID: 0,
		})
		return
	}

	discount, err := strconv.Atoi(context.Request.FormValue("discount"))
	// Check the conversion of the discount to integer type
	if err != nil {
		context.JSON(http.StatusBadRequest, response.GenericCreateResponse{
			Status:   "error",
			Code:     http.StatusBadRequest,
			Message:  "Invalid discount value",
			Err:      err.Error(),
			RecordID: 0,
		})
		return
	}

	discountedPrice, err := strconv.Atoi(context.Request.FormValue("discounted_price"))
	// Check the conversion of the discounted price to integer type
	if err != nil {
		context.JSON(http.StatusBadRequest, response.GenericCreateResponse{
			Status:   "error",
			Code:     http.StatusBadRequest,
			Message:  "Invalid discounted price value",
			Err:      err.Error(),
			RecordID: 0,
		})
		return
	}
	status := context.Request.FormValue("status")

	//Check the uploaded logo file details
	logoPath, fileHandleError := utils.HandleFileUpload(context, "products")
	if fileHandleError != nil {
		context.JSON(http.StatusBadRequest, response.GenericCreateResponse{
			Status:   "error",
			Code:     http.StatusBadRequest,
			Message:  fileHandleError.Error(),
			Err:      fileHandleError.Error(),
			RecordID: 0,
		})
		return
	}

	//Populate the product DTO
	productDTO.Name = name
	productDTO.Description = description
	productDTO.Price = price
	productDTO.Discount = discount
	productDTO.DiscountedPrice = discountedPrice
	productDTO.Status = status
	productDTO.Logo = logoPath
	productDTO.ProductCategory = product_categories_id

	// Call the service layer to create the product
	productService := service.NewProductService(utils.DBInitialize())
	businessService := service.NewBusinessService(utils.DBInitialize())
	businessDetails := businessService.GetUserBusiness(userID)
	productDTO.BusinessID = int(businessDetails.Data[0].ID)

	createdProductID, err := productService.CreateProduct(&productDTO)
	if err != nil {
		context.JSON(http.StatusInternalServerError, response.GenericCreateResponse{
			Status:   "error",
			Code:     http.StatusInternalServerError,
			Message:  err.Error(),
			Err:      err.Error(),
			RecordID: 0,
		})
		return
	}

	// Return the created product ID in the response
	context.JSON(http.StatusOK, response.CreateProductResponse{
		Status:   "success",
		Code:     http.StatusOK,
		Message:  "Product created successfully",
		RecordID: createdProductID,
	})
}

// GetAllProducts handles the retrieval of all products
func GetAllProducts(c *gin.Context) {
	productService := service.NewProductService(utils.DBInitialize())
	products, err := productService.GetAllProducts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ProductResponse{
			Status:  "error",
			Code:    http.StatusInternalServerError,
			Message: "Failed to retrieve products",
			Err:     err.Error(),
			Data:    nil,
		})
		return
	}
	// Run this when the operation (select records) is done successfully
	c.JSON(http.StatusOK, response.ProductResponse{
		Status:  "success",
		Code:    http.StatusOK,
		Message: "Products retrieved successfully",
		Data:    products,
	})
}
