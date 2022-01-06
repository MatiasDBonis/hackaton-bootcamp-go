package internal

import "fmt"

type Service interface {
	InsertAll(sales []Sales) (int, error)
}

type service struct {
	repo Repository
}

func NewService(repository Repository) Service {
	return &service{repo: repository}
}

func (s *service) InsertAll(sales []Sales) (int, error) {
	affectedRows := 0
	totalRows := len(sales)

	for _, sale := range sales {
		_, err := s.repo.Insert(sale)
		if err == nil {
			affectedRows++
		}
	}

	if affectedRows != totalRows {
		return affectedRows, fmt.Errorf("hubo error insertando %v ventas", totalRows-affectedRows)
	}
	return affectedRows, nil
}
