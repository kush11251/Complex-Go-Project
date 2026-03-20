package repository

import (
	"example.com/project/pkg/models"
	"database/sql"
)

// ProductRepository provides product database operations
// @Description product repository
// @receiver pr
// @return ProductRepository
func NewProductRepository() *ProductRepository {
	return &ProductRepository{db: db.GetDB()}
}

type ProductRepository struct {
	db *sql.DB
}

func (pr *ProductRepository) FindAllProducts() ([]models.Product, error) {
	// implement product database query
	return nil, nil
}