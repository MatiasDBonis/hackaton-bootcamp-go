package customers

import (
	"fmt"

	"github.com/MatiasDBonis/hackaton-bootcamp-go.git/internal/domain"
)

type Service interface {
	InsertAll([]domain.Customers) (int, error)
	Update(customer domain.Customers) (domain.Customers, error)
	GetTotals() ([]domain.DTOCustomerTotals, error)
}

type service struct {
	repo Repository
}

func NewService(repository Repository) Service {
	return &service{repo: repository}
}

func (s *service) InsertAll(customers []domain.Customers) (int, error) {
	affectedRows := 0
	totalRows := len(customers)

	for _, customer := range customers {
		_, err := s.repo.Insert(customer)
		if err == nil {
			affectedRows++
		}
	}

	if affectedRows != totalRows {
		return affectedRows, fmt.Errorf("hubo error insertando %v customers", totalRows-affectedRows)
	}
	return affectedRows, nil
}

func (s *service) Update(customer domain.Customers) (domain.Customers, error) {
	return s.repo.Update(customer)
}

func (s *service) GetTotals() ([]domain.DTOCustomerTotals, error) {
	data, err := s.repo.GetTotals()
	return data, err
}
