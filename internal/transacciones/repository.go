package transacciones

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/MiguelAngelCipamochaF/go-web/pkg/store"
)

type Transaction struct {
	ID       int    `json:"id"`
	Codigo   string `json:"codigo"`
	Moneda   string `json:"moneda"`
	Monto    int    `json:"monto"`
	Emisor   string `json:"emisor"`
	Receptor string `json:"receptor"`
	Fecha    string `json:"fecha"`
}

type Repository interface {
	GetAll() ([]Transaction, error)
	GetByID(id int) (Transaction, error)
	GetField(v interface{}, name string) (interface{}, error)
	GenID() (int, error)
	Store(id int, codigo string, moneda string, monto int, emisor string, receptor string, fecha string) (Transaction, error)
	Update(newT Transaction, id int) (Transaction, error)
	Delete(id int) error
	Patch(id int, codigo string, monto int) (Transaction, error)
}

type repository struct {
	db store.Store
}

func NewRepository(db store.Store) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) GetAll() ([]Transaction, error) {
	var ts []Transaction
	err := r.db.Read(&ts)
	if err != nil {
		return nil, err
	}
	return ts, nil
}

func (r *repository) GetByID(id int) (Transaction, error) {
	var transacciones []Transaction

	if err := r.db.Read(&transacciones); err != nil {
		return Transaction{}, err
	}

	for _, t := range transacciones {
		if t.ID == id {
			return t, nil
		}
	}
	return Transaction{}, fmt.Errorf("error: unknown transaction with ID %d", id)
}

func (r *repository) GetField(v interface{}, name string) (interface{}, error) {
	rv := reflect.ValueOf(v)
	if rv.Kind() != reflect.Ptr || rv.Elem().Kind() != reflect.Struct {
		return nil, errors.New("v debe ser un puntero a una estructura")
	}

	rv = rv.Elem()

	fv := rv.FieldByName(name)

	if !fv.IsValid() {
		return nil, fmt.Errorf("%s no existe en la estructura", name)
	}

	if !fv.CanSet() {
		return nil, fmt.Errorf("no es posible acceder al campo %s", name)
	}

	if fv.IsZero() {
		return nil, fmt.Errorf("el campo %s esta vacio", name)
	}

	return fv, nil
}

func (r *repository) GenID() (int, error) {
	var ts []Transaction
	if err := r.db.Read(&ts); err != nil {
		return 0, err
	}

	lastID := 0

	if len(ts) == 0 {
		return lastID, nil
	}

	lastID = ts[len(ts)-1].ID

	return lastID + 1, nil
}

func (r *repository) Store(id int, codigo string, moneda string, monto int, emisor string, receptor string, fecha string) (Transaction, error) {
	var ts []Transaction
	_ = r.db.Read(&ts)
	transaction := Transaction{
		ID:       id,
		Codigo:   codigo,
		Moneda:   moneda,
		Monto:    monto,
		Emisor:   emisor,
		Receptor: receptor,
		Fecha:    fecha,
	}
	ts = append(ts, transaction)

	if err := r.db.Write(ts); err != nil {
		return Transaction{}, err
	}
	return transaction, nil
}

func (r *repository) Update(newT Transaction, id int) (Transaction, error) {
	var transacciones []Transaction
	if err := r.db.Read(&transacciones); err != nil {
		return Transaction{}, err
	}

	for i := range transacciones {
		if transacciones[i].ID == id {
			transacciones[i] = newT
			transacciones[i].ID = id
			_ = r.db.Write(transacciones)
			return transacciones[i], nil
		}
	}
	return Transaction{}, errors.New("error: no existe transaccion con ese ID")
}

func (r *repository) Delete(id int) error {
	var transacciones []Transaction

	if err := r.db.Read(&transacciones); err != nil {
		return err
	}

	for i := range transacciones {
		if transacciones[i].ID == id {
			transacciones = append(transacciones[:i], transacciones[i+1:]...)
			_ = r.db.Write(transacciones)
			return nil
		}
	}

	return errors.New("error: no existe transaccion con ese ID")
}

func (r *repository) Patch(id int, codigo string, monto int) (Transaction, error) {
	var transacciones []Transaction

	if err := r.db.Read(&transacciones); err != nil {
		return Transaction{}, err
	}

	for i := range transacciones {
		if transacciones[i].ID == id {
			transacciones[i].Codigo = codigo
			transacciones[i].Monto = monto
			_ = r.db.Write(transacciones)
			return transacciones[i], nil
		}
	}
	return Transaction{}, fmt.Errorf("error: no existe transaccion con ID: %v", id)
}
