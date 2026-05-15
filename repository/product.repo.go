package repository

import "gorm.io/gorm"

// ProductRepository structures the product repository
type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{db: db}
}
