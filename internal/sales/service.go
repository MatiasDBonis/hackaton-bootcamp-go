package sales

import (
	"fmt"

	"github.com/MatiasDBonis/hackaton-bootcamp-go.git/internal/domain"
)

type Service interface {
	InsertAll(sales []domain.Sales) (int, error)
	Update(sale domain.Sales) (domain.Sales, error)
}

type service struct {
	repo Repository
}

func NewService(repository Repository) Service {
	return &service{repo: repository}
}

func (s *service) InsertAll(sales []domain.Sales) (int, error) {
	affectedRows := 0
	totalRows := len(sales)

	for _, sale := range sales {
		_, err := s.repo.Insert(sale)
		fmt.Println(err.Error())
		if err == nil {
			affectedRows++
		}
	}

	if affectedRows != totalRows {
		return affectedRows, fmt.Errorf("hubo error insertando %v ventas", totalRows-affectedRows)
	}
	return affectedRows, nil
}

func (s *service) Update(sale domain.Sales) (domain.Sales, error) {
	return s.repo.Update(sale)
}
