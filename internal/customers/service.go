package internal

import "fmt"

type Service interface {
	InsertAll([]Customers) (int, error)
}

type service struct {
	repo Repository
}

func NewService(repository Repository) Service {
	return &service{repo: repository}
}

func (s *service) InsertAll(customers []Customers) (int, error) {
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
