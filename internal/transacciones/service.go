package transacciones

import "fmt"

type Service interface {
	GetAll() []Transaction
	GetByID(id int) *Transaction
	GetField(v interface{}, name string) (interface{}, error)
	GenID() int
	Store(tRequest Transaction)
	Update(newT Transaction, id int) (Transaction, error)
	Delete(id int) error
	Patch(id int, codigo string, monto int) (Transaction, error)
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

func (s *service) Update(newT Transaction, id int) (Transaction, error) {
	modT, err := s.repo.Update(newT, id)

	if err != nil {
		return Transaction{}, err
	}

	return modT, nil
}

func (s *service) Delete(id int) error {
	return s.repo.Delete(id)
}

func (s *service) Patch(id int, codigo string, monto int) (Transaction, error) {
	transaction, err := s.repo.Patch(id, codigo, monto)

	if err != nil {
		return Transaction{}, err
	}

	return transaction, nil
}
