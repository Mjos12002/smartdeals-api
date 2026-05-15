package service

// This file contains the service layer for handling business logic related to product categories.

import (
	"gorm.io/gorm"
	"smartdeals.rw/model"
)

// ProductCategory creates a new product category in the database
type ProductCategoryService struct {
	db *gorm.DB
}

// NewProductCategoryService is used to create a new object of the ProductCategoryService
func NewProductCategoryService(db *gorm.DB) *ProductCategoryService {
	return &ProductCategoryService{db: db}
}

// CreateProductCategory is used to create a new product record
func (s *ProductCategoryService) CreateProductCategory(productCategory *model.ProductCategory) (int, error) {
	productCategoryCreated := s.db.Create(&productCategory)
	if productCategoryCreated.Error != nil {
		return 0, productCategoryCreated.Error
	}
	return int(productCategory.ID), nil
}

// GetAllProductCategories is used to return the list of product categories
func (s *ProductCategoryService) GetAllProductCategories() ([]model.ProductCategory, error) {

	var productCategoryResponse []model.ProductCategory
	if err := s.db.Find(&productCategoryResponse).Error; err != nil {
		return nil, err
	}
	return productCategoryResponse, nil
}
