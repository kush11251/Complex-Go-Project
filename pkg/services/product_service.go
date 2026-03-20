package services

import (
	"example.com/project/pkg/models"
	"example.com/project/pkg/repository"
)

// ProductService provides product related functionality
// @Description product service
// @receiver ps
// @return ProductService
func NewProductService() *ProductService {
	return &ProductService{ProductRepo: repository.NewProductRepository()}
}

// GetProductSvc provides product svc instance
func GetProductSvc() *ProductService {
	return NewProductService()
}

type ProductService struct {
	ProductRepo repository.ProductRepository
}

func (ps *ProductService) GetProducts() ([]models.Product, error) {
	return ps.ProductRepo.FindAllProducts()
}