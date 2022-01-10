package products

import (
	"fmt"

	"github.com/MatiasDBonis/hackaton-bootcamp-go.git/internal/domain"
)

type Service interface {
	InsertAll(products []domain.Products) (int, error)
	Update(product domain.Products) (domain.Products, error)
}

type service struct {
	repo Repository
}

func NewService(repository Repository) Service {
	return &service{repo: repository}
}

func (s *service) InsertAll(products []domain.Products) (int, error) {
	affectedRows := 0
	totalRows := len(products)

	for _, product := range products {
		_, err := s.repo.Insert(product)
		if err == nil {
			affectedRows++
		}
	}

	if affectedRows != totalRows {
		return affectedRows, fmt.Errorf("hubo error insertando %v productos", totalRows-affectedRows)
	}
	return affectedRows, nil
}

func (s *service) Update(product domain.Products) (domain.Products, error) {
	return s.repo.Update(product)
}
