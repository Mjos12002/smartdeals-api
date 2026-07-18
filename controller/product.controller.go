package controller

// This file contains the controller functions for handling Product-related HTTP requests.

import (
	"mime/multipart"
	"net/http"
	"path/filepath"
	"strconv"

	"smartdeals.rw/dto"
	"smartdeals.rw/response"
	"smartdeals.rw/service"
	"smartdeals.rw/utils"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// CreateProduct handles the creation of a new product
func CreateProduct(context *gin.Context) {
	var productDTO dto.ProductDTO

	name := context.Request.FormValue("name")
	description := context.Request.FormValue("description")

	logoFile, logoFileHeader, fileError := context.Request.FormFile("logo")

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

	discountStartDate := context.Request.FormValue("discount_start_date")
	discountEndDate := context.Request.FormValue("discount_end_date")
	status := context.Request.FormValue("status")

	//Check the uploaded logo file details
	if fileError != nil {
		context.JSON(http.StatusBadRequest, response.GenericCreateResponse{
			Status:   "error",
			Code:     http.StatusBadRequest,
			Message:  "Invalid file",
			Err:      fileError.Error(),
			RecordID: 0,
		})
		return
	}
	logoPath, fileHandleError := utils.HandleFileUpload(logoFile, logoFileHeader, context)
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
	productDTO.DiscountStartDate = discountStartDate
	productDTO.DiscountEndDate = discountEndDate
	productDTO.Status = status
	productDTO.Logo = logoPath
	// Call the service layer to create the product
	productService := service.NewProductService(utils.DBInitialize())
	createdProductID, err := productService.CreateProduct(&productDTO)
	if err != nil {
		context.JSON(http.StatusInternalServerError, response.GenericCreateResponse{
			Status:   "error",
			Code:     http.StatusInternalServerError,
			Message:  "Failed to create product",
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

	

	c.JSON(http.StatusOK, response.ProductResponse{
		Status:  "success",
		Code:    http.StatusOK,
		Message: "Products retrieved successfully",
		Data:    products,
	})
}

func HandleFileUploading(logoFile multipart.File, logoFileHeader *multipart.FileHeader, context *gin.Context) error {

	defer logoFile.Close()

	logoFileExt := filepath.Ext(logoFileHeader.Filename)

	newFilename := uuid.New().String() + logoFileExt

	logoFileDest := filepath.Join("./resources", newFilename)

	if err := context.SaveUploadedFile(logoFileHeader, logoFileDest); err != nil {
		return err
	}

	return nil
}
