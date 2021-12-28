package transacciones

import "fmt"

type Service interface {
	GetAll() ([]Transaction, error)
	GetByID(id int) (Transaction, error)
	GetField(v interface{}, name string) (interface{}, error)
	GenID() (int, error)
	Store(id int, codigo string, moneda string, monto int, emisor string, receptor string, fecha string) (Transaction, error)
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

func (s *service) GetAll() ([]Transaction, error) {
	ts, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}
	return ts, nil
}

func (s *service) GetByID(id int) (Transaction, error) {
	tid, err := s.repo.GetByID(id)
	if err != nil {
		return Transaction{}, err
	}
	return tid, nil
}

func (s *service) GetField(v interface{}, name string) (interface{}, error) {
	tGetField, err := s.repo.GetField(v, name)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return tGetField, nil
}

func (s *service) GenID() (int, error) {
	tGenId, err := s.repo.GenID()

	if err != nil {
		return 0, err
	}

	return tGenId, nil
}

func (s *service) Store(id int, codigo string, moneda string, monto int, emisor string, receptor string, fecha string) (Transaction, error) {
	transaction, err := s.repo.Store(id, codigo, moneda, monto, emisor, receptor, fecha)

	if err != nil {
		return Transaction{}, err
	}

	return transaction, nil
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
