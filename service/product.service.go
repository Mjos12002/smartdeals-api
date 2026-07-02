package service

// This file contains the service layer for handling business logic related to products.

import (
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
		DiscountStartDate:   productDTO.DiscountStartDate,
		DiscountEndDate:     productDTO.DiscountEndDate,
		Status:              productDTO.Status,
		Logo:                productDTO.Logo,
		ProductCategoriesID: 1,
		BusinessesID:        1,
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
	var products []model.Products
	if err := s.db.Find(&products).Error; err != nil {
		return nil, err
	}
	productResponses := make([]response.ProductResponseData, len(products))
	for i, product := range products {
		productResponses[i] = response.ProductResponseData{
			Name:                product.Name,
			Description:         product.Description,
			Price:               product.Price,
			Discount:            product.Discount,
			DiscountedPrice:     product.DiscountedPrice,
			DiscountStartDate:   product.DiscountStartDate,
			DiscountEndDate:     product.DiscountEndDate,
			Status:              product.Status,
			ProductCategoriesID: product.ProductCategoriesID,
			Logo:                product.Logo,
			BusinessesID:        product.BusinessesID,
		}
	}
	return productResponses, nil
}
