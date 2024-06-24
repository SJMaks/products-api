package repository

import (
	"fmt"

	"github.com/SJMaks/products-api"
	"github.com/jmoiron/sqlx"
)

type ProductPostgres struct {
	db *sqlx.DB
}

func NewProductPostgres(db *sqlx.DB) *ProductPostgres {
	return &ProductPostgres{db: db}
}

func (r *ProductPostgres) CreateProduct(product products.Product) (int, error) {
	var id int

	query := fmt.Sprintf("INSERT INTO %s (name, quantity, unit_coast, measure) values ($1, $2, $3, $4) RETURNING id", productsTable)
	row := r.db.QueryRow(query, product.Name, product.Quantity, product.Unit_coast, product.Measure)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *ProductPostgres) GetProduct(id int) (products.Product, error) {
	var product products.Product

	query := fmt.Sprintf("SELECT * FROM %s WHERE id=$1", productsTable)
	err := r.db.Get(&product, query, id)

	return product, err
}

func (r *ProductPostgres) GetProducts() ([]products.Product, error) {
	var products []products.Product

	query := fmt.Sprintf("SELECT * FROM %s", productsTable)
	err := r.db.Select(&products, query)

	return products, err
}

func (r *ProductPostgres) DeleteProduct(id int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id=$1", productsTable)
	_, err := r.db.Exec(query, id)

	return err
}

func (r *ProductPostgres) UpdateProduct(product products.Product, id int) error {
	query := fmt.Sprintf("UPDATE %s SET name = $1, quantity = $2, unit_coast = $3, measure = $4 WHERE id = $5", productsTable)
	_, err := r.db.Exec(query, product.Name, product.Quantity, product.Unit_coast, product.Measure, id)

	return err
}
