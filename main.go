package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/gin-gonic/gin"
)

type Transaction struct {
	ID       int
	Codigo   string
	Moneda   string
	Monto    int
	Emisor   string
	Receptor string
	Fecha    string
}

func main() {
	router := gin.Default()
	router.GET("/hola/:name", func(c *gin.Context) {
		name := c.Param("name")
		saludo := "Hola " + name
		c.JSON(200, gin.H{
			"message": saludo,
		})
	})

	router.GET("/transacciones", GetAll)
	router.Run()
}

func GetAll(c *gin.Context) {
	data, _ := ioutil.ReadFile("transactions.json")
	var transacciones []Transaction

	err := json.Unmarshal(data, &transacciones)

	if err != nil {
		fmt.Println(err)
	}

	c.JSON(200, transacciones)
}
