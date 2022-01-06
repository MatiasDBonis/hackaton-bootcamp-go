package internal

type Service interface {
	InsertAll([]Products)
}

type service struct {
	repo Repository
}

func NewService(repository Repository) Service {
	return &service{repo: repository}
}

func (s *service) InsertAll([]Products) {}
