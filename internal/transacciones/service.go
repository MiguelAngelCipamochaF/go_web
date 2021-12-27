package transacciones

import "fmt"

type Service interface {
	GetAll() []Transaction
	GetByID(id int) *Transaction
	GetField(v interface{}, name string) (interface{}, error)
	GenID() int
	Store(tRequest Transaction)
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

func (s *service) GetByID(id int) *Transaction {
	tid := s.repo.GetByID(id)
	return tid
}

func (s *service) GetField(v interface{}, name string) (interface{}, error) {
	tGetField, err := s.repo.GetField(v, name)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return tGetField, nil
}

func (s *service) GenID() int {
	tGenId := s.repo.GenID()
	return tGenId
}

func (s *service) Store(tRequest Transaction) {
	s.repo.Store(tRequest)
}
