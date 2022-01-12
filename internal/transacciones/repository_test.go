package transacciones

import (
	"encoding/json"
	"errors"
	"testing"

	"github.com/MiguelAngelCipamochaF/go-web/pkg/store"
	"github.com/stretchr/testify/assert"
)

func TestGetAllError(t *testing.T) {
	expectedError := errors.New("error for GetAll")
	dbMock := store.Mock{
		Data: nil,
		Err:  expectedError,
	}
	storeMocked := store.FileStore{
		FileName: "",
		Mock:     &dbMock,
	}
	myRepo := NewRepository(&storeMocked)
	_, err := myRepo.GetAll()
	assert.Equal(t, err, expectedError)
}

func TestGetAll(t *testing.T) {
	input := []Transaction{
		{
			ID:       1,
			Codigo:   "001",
			Moneda:   "USD",
			Monto:    1500,
			Emisor:   "Miguel Cipamocha",
			Receptor: "Miles Morales",
			Fecha:    "12-01-2022",
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
	trans, _ := myRepo.GetAll()
	assert.Equal(t, trans, input)
}
