package controlador

import (
	"github.com/MiguelAngelCipamochaFigueredo/internal/transacciones"
	"github.com/gin-gonic/gin"
)

type Transaction struct {
	service transacciones.Service
}

func newTransaction(t transacciones.Service) *Transaction {
	return &Transaction{
		service: t,
	}
}

func (t *Transaction) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		transactions := ctx.service.GetAll()
		ctx.JSON(200, transactions)
	}
}
