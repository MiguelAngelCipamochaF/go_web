package transacciones

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"reflect"
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
	GetAll() []Transaction
	GetByID(id int) *Transaction
	GetField(v interface{}, name string) (interface{}, error)
	GenID() int
	Store(tRequest Transaction)
	Update(newT Transaction, id int) (Transaction, error)
	Delete(id int) error
	Patch(id int, codigo string, monto int) (Transaction, error)
}

type repository struct {
}

var transacciones []Transaction

func NewRepository() Repository {
	data, e := ioutil.ReadFile("./internal/transacciones/transactions.json")

	if e != nil {
		fmt.Println(e)
	}

	err := json.Unmarshal(data, &transacciones)

	if err != nil {
		fmt.Println(err)
	}
	return &repository{}
}

func (r *repository) GetAll() []Transaction {
	return transacciones
}

func (r *repository) GetByID(id int) *Transaction {
	for _, t := range transacciones {
		if t.ID == id {
			return &t
		}
	}
	return nil
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

func (r *repository) GenID() int {
	lastId := 0

	if len(transacciones) > 0 {
		lastId = transacciones[len(transacciones)-1].ID
	}

	return lastId + 1
}

func (r *repository) Store(tRequest Transaction) {
	transacciones = append(transacciones, tRequest)
}

func (r *repository) Update(newT Transaction, id int) (Transaction, error) {
	for i, _ := range transacciones {
		if transacciones[i].ID == id {
			transacciones[i] = newT
			transacciones[i].ID = id
			return transacciones[i], nil
		}
	}
	return Transaction{}, fmt.Errorf("error: no existe transaccion con ID: %d", id)
}

func (r *repository) Delete(id int) error {
	for i, _ := range transacciones {
		if transacciones[i].ID == id {
			transacciones = append(transacciones[:i], transacciones[i+1:]...)
			return nil
		}
	}

	return fmt.Errorf("error: no existe transaccion con ID: %v", id)
}

func (r *repository) Patch(id int, codigo string, monto int) (Transaction, error) {
	for i, _ := range transacciones {
		if transacciones[i].ID == id {
			transacciones[i].Codigo = codigo
			transacciones[i].Monto = monto
			return transacciones[i], nil
		}
	}
	return Transaction{}, fmt.Errorf("error: no existe transaccion con ID: %v", id)
}
