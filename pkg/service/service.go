package service

import (
	"github.com/SJMaks/products-api"
	"github.com/SJMaks/products-api/pkg/repository"
)

type Product interface {
	CreateProduct(product products.Product) (int, error)
	GetProduct(id int) (products.Product, error)
	GetProducts() ([]products.Product, error)
	DeleteProduct(id int) error
	UpdateProduct(product products.Product, id int) error
}

type Measure interface {
	CreateMeasure(measure products.Measure) (int, error)
	GetMeasure(id int) (products.Measure, error)
	GetMeasures() ([]products.Measure, error)
	DeleteMeasure(id int) error
	UpdateMeasure(measure products.Measure, id int) error
}

type Service struct {
	Product
	Measure
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Product: NewProductService(repos.Product),
		Measure: NewMeasureService(repos.Measure),
	}
}
