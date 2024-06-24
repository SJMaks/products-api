package service

import (
	"github.com/SJMaks/products-api"
	"github.com/SJMaks/products-api/pkg/repository"
)

type ProductService struct {
	repo repository.Product
}

func NewProductService(repo repository.Product) *ProductService {
	return &ProductService{repo: repo}
}

func (s *ProductService) CreateProduct(product products.Product) (int, error) {
	return s.repo.CreateProduct(product)
}

func (s *ProductService) GetProduct(id int) (products.Product, error) {
	return s.repo.GetProduct(id)
}

func (s *ProductService) GetProducts() ([]products.Product, error) {
	return s.repo.GetProducts()
}

func (s *ProductService) DeleteProduct(id int) error {
	return s.repo.DeleteProduct(id)
}

func (s *ProductService) UpdateProduct(product products.Product, id int) error {
	return s.repo.UpdateProduct(product, id)
}
