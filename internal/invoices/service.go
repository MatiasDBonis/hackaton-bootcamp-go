package invoices

import (
	"fmt"

	"github.com/MatiasDBonis/hackaton-bootcamp-go.git/internal/domain"
)

type Service interface {
	InsertAll(invoices []domain.Invoices) (int, error)
	Update(invoice domain.Invoices) (domain.Invoices, error)
}

type service struct {
	repo Repository
}

func NewService(repository Repository) Service {
	return &service{repo: repository}
}

func (s *service) InsertAll(invoices []domain.Invoices) (int, error) {
	affectedRows := 0
	totalRows := len(invoices)

	for _, invoice := range invoices {
		_, err := s.repo.Insert(invoice)
		if err == nil {
			affectedRows++
		}
	}

	if affectedRows != totalRows {
		return affectedRows, fmt.Errorf("hubo error insertando %v invoices", totalRows-affectedRows)
	}
	return affectedRows, nil
}

func (s *service) Update(invoice domain.Invoices) (domain.Invoices, error) {
	return s.repo.Update(invoice)
}
