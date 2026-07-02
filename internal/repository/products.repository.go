package repository

import (
	"my-diet-server/internal/models"

	"gorm.io/gorm"
)

type ProductsRepository interface {
	CreateProduct(product *models.ProductModel) (models.ProductModel, error)
}

type productsRepo struct {
	db *gorm.DB
}

// CreateProduct implements ProductsRepository.
func (p *productsRepo) CreateProduct(product *models.ProductModel) (models.ProductModel, error) {
	panic("unimplemented")
}

func NewProductsRepository(db *gorm.DB) ProductsRepository {
	return &productsRepo{db}
}
