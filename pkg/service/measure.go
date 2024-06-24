package service

import (
	"github.com/SJMaks/products-api"
	"github.com/SJMaks/products-api/pkg/repository"
)

type MeasureService struct {
	repo repository.Measure
}

func NewMeasureService(repo repository.Measure) *MeasureService {
	return &MeasureService{repo: repo}
}

func (s *MeasureService) CreateMeasure(measure products.Measure) (int, error) {
	return s.repo.CreateMeasure(measure)
}

func (s *MeasureService) GetMeasure(id int) (products.Measure, error) {
	return s.repo.GetMeasure(id)
}

func (s *MeasureService) GetMeasures() ([]products.Measure, error) {
	return s.repo.GetMeasures()
}

func (s *MeasureService) DeleteMeasure(id int) error {
	return s.repo.DeleteMeasure(id)
}

func (s *MeasureService) UpdateMeasure(measure products.Measure, id int) error {
	return s.repo.UpdateMeasure(measure, id)
}
