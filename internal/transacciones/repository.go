package transacciones

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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
}

type repository struct {
}

var transacciones []Transaction

func NewRepository() Repository {
	return &repository{}
}

func (r *repository) GetAll() []Transaction {
	data, _ := ioutil.ReadFile("transactions.json")

	err := json.Unmarshal(data, &transacciones)

	if err != nil {
		fmt.Println(err)
	}

	return transacciones
}
