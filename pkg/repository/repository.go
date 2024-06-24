package repository

import (
	"github.com/SJMaks/products-api"
	"github.com/jmoiron/sqlx"
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

type Repository struct {
	Product
	Measure
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Product: NewProductPostgres(db),
		Measure: NewMeasurePostgres(db),
	}
}
