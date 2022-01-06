package internal

type Service interface{}

type service struct {
	repo Repository
}

func NewService(repository Repository) Service {
	return &service{repo: repository}
}
