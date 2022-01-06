package internal

import "fmt"

type Service interface {
	InsertAll(invoices []Invoices) (int, error)
}

type service struct {
	repo Repository
}

func NewService(repository Repository) Service {
	return &service{repo: repository}
}

func (s *service) InsertAll(invoices []Invoices) (int, error) {
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
