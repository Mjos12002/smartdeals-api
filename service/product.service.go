package service

// This file contains the service layer for handling business logic related to products.

import (
	"gorm.io/gorm"
	"smartdeals.rw/model"
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
func (s *ProductService) CreateProduct(product *model.Products) (int, error) {
	// Return the created product and any error encountered
	productCreated := s.db.Create(&product)
	if productCreated.Error != nil {
		return 0, productCreated.Error
	}
	return int(product.ID), nil

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

// Get all products
func (s *ProductService) GetAllProducts() ([]model.Products, error) {
	var products []model.Products
	if err := s.db.Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}
