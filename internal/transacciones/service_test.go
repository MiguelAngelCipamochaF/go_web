package transacciones

import (
	"encoding/json"
	"errors"
	"testing"

	"github.com/MiguelAngelCipamochaF/go-web/pkg/store"
	"github.com/stretchr/testify/assert"
)

func TestServiceUpdateError(t *testing.T) {
	newTrans := Transaction{
		ID:       4,
		Codigo:   "001",
		Moneda:   "USD",
		Monto:    1500,
		Emisor:   "Miguel Cipamocha",
		Receptor: "Miles Morales",
		Fecha:    "12-01-2022",
	}
	input := []Transaction{
		{
			ID:       1,
			Codigo:   "001",
			Moneda:   "USD",
			Monto:    5000,
			Emisor:   "Miguel",
			Receptor: "Miles",
			Fecha:    "02-01-2022",
		},
		{
			ID:       2,
			Codigo:   "002",
			Moneda:   "USD",
			Monto:    4500,
			Emisor:   "Miles Morales",
			Receptor: "Miguel Cipamocha",
			Fecha:    "12-01-2022",
		},
	}

	dataJson, _ := json.Marshal(input)
	dbMock := store.Mock{
		Data: dataJson,
		Err:  nil,
	}
	storeMocked := store.FileStore{
		FileName: "",
		Mock:     &dbMock,
	}
	myRepo := NewRepository(&storeMocked)
	myService := NewService(myRepo)

	_, err := myService.Update(newTrans, newTrans.ID)

	assert.Equal(t, err, errors.New("error: no existe transaccion con ese ID"))
}

func TestServiceUpdate(t *testing.T) {
	newTrans := Transaction{
		ID:       1,
		Codigo:   "001",
		Moneda:   "USD",
		Monto:    1500,
		Emisor:   "Miguel Cipamocha",
		Receptor: "Miles Morales",
		Fecha:    "12-01-2022",
	}
	input := []Transaction{
		{
			ID:       1,
			Codigo:   "001",
			Moneda:   "USD",
			Monto:    5000,
			Emisor:   "Miguel",
			Receptor: "Miles",
			Fecha:    "02-01-2022",
		},
		{
			ID:       2,
			Codigo:   "002",
			Moneda:   "USD",
			Monto:    4500,
			Emisor:   "Miles Morales",
			Receptor: "Miguel Cipamocha",
			Fecha:    "12-01-2022",
		},
	}

	dataJson, _ := json.Marshal(input)
	dbMock := store.Mock{
		Data: dataJson,
		Err:  nil,
	}
	storeMocked := store.FileStore{
		FileName: "",
		Mock:     &dbMock,
	}
	myRepo := NewRepository(&storeMocked)
	myService := NewService(myRepo)

	result, _ := myService.Update(newTrans, newTrans.ID)

	assert.Equal(t, result, newTrans)
}

func TestServiceDeleteError(t *testing.T) {
	idToDelete := 4
	input := []Transaction{
		{
			ID:       1,
			Codigo:   "001",
			Moneda:   "USD",
			Monto:    5000,
			Emisor:   "Miguel",
			Receptor: "Miles",
			Fecha:    "02-01-2022",
		},
		{
			ID:       2,
			Codigo:   "002",
			Moneda:   "USD",
			Monto:    4500,
			Emisor:   "Miles Morales",
			Receptor: "Miguel Cipamocha",
			Fecha:    "12-01-2022",
		},
	}

	dataJson, _ := json.Marshal(input)
	dbMock := store.Mock{
		Data: dataJson,
		Err:  nil,
	}
	storeMocked := store.FileStore{
		FileName: "",
		Mock:     &dbMock,
	}
	myRepo := NewRepository(&storeMocked)
	myService := NewService(myRepo)

	err := myService.Delete(idToDelete)
	assert.Equal(t, err, errors.New("error: no existe transaccion con ese ID"))
}

func TestServiceDelete(t *testing.T) {
	idToDelete := 2
	input := []Transaction{
		{
			ID:       1,
			Codigo:   "001",
			Moneda:   "USD",
			Monto:    5000,
			Emisor:   "Miguel",
			Receptor: "Miles",
			Fecha:    "02-01-2022",
		},
		{
			ID:       2,
			Codigo:   "002",
			Moneda:   "USD",
			Monto:    4500,
			Emisor:   "Miles Morales",
			Receptor: "Miguel Cipamocha",
			Fecha:    "12-01-2022",
		},
	}

	dataJson, _ := json.Marshal(input)
	dbMock := store.Mock{
		Data: dataJson,
		Err:  nil,
	}
	storeMocked := store.FileStore{
		FileName: "",
		Mock:     &dbMock,
	}
	myRepo := NewRepository(&storeMocked)
	myService := NewService(myRepo)

	err := myService.Delete(idToDelete)
	_, err2 := myService.GetByID(idToDelete)

	assert.Equal(t, err, nil)
	assert.Equal(t, err2, nil)
}
