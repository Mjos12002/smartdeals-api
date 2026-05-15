package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"smartdeals.rw/dto"
	"smartdeals.rw/model"
	"smartdeals.rw/response"
	"smartdeals.rw/service"
	"smartdeals.rw/utils"
)

// CreateProductCategory is used to create a new product category record
func CreateProductCategory(c *gin.Context) {

	var productCategoryDTO dto.ProductCategoryDTO

	// Check availability of the request fields
	if err := c.ShouldBindJSON(&productCategoryDTO); err != nil {
		c.JSON(http.StatusBadRequest, response.GenericCreateResponse{
			Status:   "error",
			Code:     http.StatusBadRequest,
			Message:  "Invalid Request Data",
			Err:      err.Error(),
			RecordID: 0,
		})
		return
	}

	// Validate data
	if err := utils.ValidateProductCategory(&productCategoryDTO); err != nil {
		c.JSON(http.StatusInternalServerError, response.GenericCreateResponse{
			Status:   "error",
			Code:     http.StatusInternalServerError,
			Message:  "Validation failed",
			Err:      err.Error(),
			RecordID: 0,
		})
		return
	}

	// Construct the model
	productCategoryModel := model.ProductCategory{
		Name:        productCategoryDTO.Name,
		Description: productCategoryDTO.Description,
	}
	// Create a product category service instance and record to the database
	productCategoryService := service.NewProductCategoryService(utils.DBInitialize())
	productCategoryID, err := productCategoryService.CreateProductCategory(&productCategoryModel)
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
	c.JSON(http.StatusOK, response.GenericCreateResponse{
		Status:   "success",
		Code:     http.StatusOK,
		Message:  "Product category created successfully",
		RecordID: productCategoryID,
	})
}

// GetAllProductCategory is used to get the lislt of
func GetAllProductCategories(c *gin.Context) {
	productCategoryService := service.NewProductCategoryService(utils.DBInitialize())
	productCategories, err := productCategoryService.GetAllProductCategories()

	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ProductCategoryResponse{
			Status:  "error",
			Code:    http.StatusInternalServerError,
			Message: "Failed to retrieve product category",
			Err:     err.Error(),
			Data:    nil,
		})
		return
	}

	c.JSON(http.StatusOK, response.ProductCategoryResponse{
		Status:  "success",
		Code:    http.StatusOK,
		Message: "Products retrieved successfully",
		Data:    productCategories,
	})
}
