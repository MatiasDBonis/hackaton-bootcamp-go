package internal

import "fmt"

type Service interface {
	InsertAll(products []Products) (int, error)
}

type service struct {
	repo Repository
}

func NewService(repository Repository) Service {
	return &service{repo: repository}
}

func (s *service) InsertAll(products []Products) (int, error) {
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
