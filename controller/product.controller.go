package controller

// This file contains the controller functions for handling Product-related HTTP requests.

import (
	"net/http"

	"smartdeals.rw/dto"
	"smartdeals.rw/model"
	"smartdeals.rw/response"
	"smartdeals.rw/service"
	"smartdeals.rw/utils"

	"github.com/gin-gonic/gin"
)

// CreateProduct handles the creation of a new product
func CreateProduct(c *gin.Context) {
	var productDTO dto.ProductDTO

	// Bind the incoming JSON to the ProductDTO struct
	if err := c.ShouldBindJSON(&productDTO); err != nil {
		c.JSON(http.StatusBadRequest, response.GenericCreateResponse{
			Status:   "error",
			Code:     http.StatusBadRequest,
			Message:  "Invalid request data",
			Err:      err.Error(),
			RecordID: 0,
		})
		return
	}

	// Validate the product data
	if err := utils.ValidateProduct(&productDTO); err != nil {
		c.JSON(http.StatusBadRequest, response.CreateProductResponse{
			Status:   "error",
			Code:     http.StatusBadRequest,
			Message:  "Validation failed",
			Err:      err.Error(),
			RecordID: 0,
		})
		return
	}

	// Create a new product model instance from the DTO
	productModel := model.Products{
		Name:              productDTO.Name,
		Description:       productDTO.Description,
		Price:             productDTO.Price,
		Discount:          productDTO.Discount,
		DiscountedPrice:   productDTO.DiscountedPrice,
		DiscountStartDate: productDTO.DiscountStartDate,
		DiscountEndDate:   productDTO.DiscountEndDate,
		Status:            productDTO.Status,
		ProductCategory:   productDTO.ProductCategory,
		Logo:              productDTO.Logo,
		BusinessID:        productDTO.BusinessID,
	}

	// Call the service layer to create the product
	productService := service.NewProductService(utils.DBInitialize())
	createdProductID, err := productService.CreateProduct(&productModel)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.GenericCreateResponse{
			Status:   "error",
			Code:     http.StatusInternalServerError,
			Message:  "Failed to create product",
			Err:      err.Error(),
			RecordID: 0,
		})
		return
	}

	// Return the created product ID in the response
	c.JSON(http.StatusOK, response.CreateProductResponse{
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

	c.JSON(http.StatusOK, response.ProductResponse{
		Status:  "success",
		Code:    http.StatusOK,
		Message: "Products retrieved successfully",
		Data:    products,
	})
}
