package service

// This file contains the service layer for handling business logic related to products.

import (
	"fmt"

	"gorm.io/gorm"
	"smartdeals.rw/dto"
	"smartdeals.rw/model"
	"smartdeals.rw/response"
)

// ProductService provides methods for product-related operations
type ProductService struct {
	db *gorm.DB // Placeholder for the database connection
}

// NewProductService creates a new instance of ProductService
func NewProductService(db *gorm.DB) *ProductService {
	return &ProductService{db: db}
}

// CreateProduct creates a new product in the database & return the created record id and any error encountered
func (s *ProductService) CreateProduct(productDTO *dto.ProductDTO) (int, error) {

	// Return the created product and any error encountered
	productModel := model.Products{
		Name:                productDTO.Name,
		Description:         productDTO.Description,
		Price:               float64(productDTO.Price),
		Discount:            float64(productDTO.Discount),
		DiscountedPrice:     float64(productDTO.DiscountedPrice),
		Status:              productDTO.Status,
		Logo:                productDTO.Logo,
		ProductCategoriesID: productDTO.ProductCategory,
		BusinessesID:        productDTO.BusinessID,
	}

	productCreated := s.db.Create(&productModel)
	if productCreated.Error != nil {
		return 0, productCreated.Error
	}
	return int(productModel.ID), nil

}

// GetProduct retrieves a product by its ID from the database
func (s *ProductService) GetProduct(id string) (*model.Products, error) {
	// Return the retrieved product and any error encountered
	var product model.Products
	if err := s.db.First(&product, id).Error; err != nil {
		return nil, err
	}
	return &product, nil
}

// Get all products and convert to the response.ProductResponseData format
func (s *ProductService) GetAllProducts() ([]response.ProductResponseData, error) {

	var products []response.Products
	if err := s.db.Model(&response.Products{}).Preload("Businesses").Preload("Categories").Find(&products).Error; err != nil {
		fmt.Printf("Error %s", err.Error())
		return nil, err
	}

	productResponses := make([]response.ProductResponseData, len(products))
	for i, product := range products {
		productResponses[i] = response.ProductResponseData{
			ID:                  product.ID,
			Name:                product.Name,
			Description:         product.Description,
			Price:               product.Price,
			Discount:            product.Discount,
			DiscountedPrice:     product.DiscountedPrice,
			Status:              product.Status,
			ProductCategoriesID: product.ProductCategoriesID,
			Logo:                product.Logo,
			Business:            product.Businesses,
			Cateegories:         product.Categories,
		}
	}
	return productResponses, nil
}

// Get all products and convert to the response.ProductResponseData format
func (s *ProductService) GetAllProductByID(id int) ([]response.ProductResponseData, error) {

	var products []response.Products
	if err := s.db.Model(&response.Products{}).Preload("Businesses").Preload("Categories").Where("id", id).Find(&products).Error; err != nil {
		fmt.Printf("Error %s", err.Error())
		return nil, err
	}

	productResponses := make([]response.ProductResponseData, len(products))
	for i, product := range products {
		productResponses[i] = response.ProductResponseData{
			ID:                  product.ID,
			Name:                product.Name,
			Description:         product.Description,
			Price:               product.Price,
			Discount:            product.Discount,
			DiscountedPrice:     product.DiscountedPrice,
			Status:              product.Status,
			ProductCategoriesID: product.ProductCategoriesID,
			Logo:                product.Logo,
			Business:            product.Businesses,
			Cateegories:         product.Categories,
		}
	}
	return productResponses, nil
}
