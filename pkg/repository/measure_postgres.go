package repository

import (
	"fmt"

	"github.com/SJMaks/products-api"
	"github.com/jmoiron/sqlx"
)

type MeasurePostgres struct {
	db *sqlx.DB
}

func NewMeasurePostgres(db *sqlx.DB) *MeasurePostgres {
	return &MeasurePostgres{db: db}
}

func (r *MeasurePostgres) CreateMeasure(measure products.Measure) (int, error) {
	var id int

	query := fmt.Sprintf("INSERT INTO %s (name) values ($1) RETURNING id", measuresTable)
	row := r.db.QueryRow(query, measure.Name)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *MeasurePostgres) GetMeasure(id int) (products.Measure, error) {
	var measure products.Measure

	query := fmt.Sprintf("SELECT * FROM %s WHERE id=$1", measuresTable)
	err := r.db.Get(&measure, query, id)

	return measure, err
}

func (r *MeasurePostgres) GetMeasures() ([]products.Measure, error) {
	var measures []products.Measure

	query := fmt.Sprintf("SELECT * FROM %s", measuresTable)
	err := r.db.Select(&measures, query)

	return measures, err
}

func (r *MeasurePostgres) DeleteMeasure(id int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id=$1", measuresTable)
	_, err := r.db.Exec(query, id)

	return err
}

func (r *MeasurePostgres) UpdateMeasure(measure products.Measure, id int) error {
	query := fmt.Sprintf("UPDATE %s SET name = $1 WHERE id = $2", measuresTable)
	_, err := r.db.Exec(query, measure.Name, id)

	return err
}
