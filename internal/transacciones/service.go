package transacciones

type Service interface {
	GetAll() []Transaction
}

type service struct {
	repo Repository
}

func NewService(s Repository) Service {
	return &service{
		repo: s,
	}
}

func (s *service) GetAll() []Transaction {
	ts := s.repo.GetAll()
	return ts
}
