package controlador

import (
	"fmt"
	"os"
	"strconv"

	"github.com/MiguelAngelCipamochaF/go-web/internal/transacciones"
	"github.com/MiguelAngelCipamochaF/go-web/pkg/web"
	"github.com/gin-gonic/gin"
)

type request struct {
	ID       int    `json:"id"`
	Codigo   string `json:"codigo"`
	Moneda   string `json:"moneda"`
	Monto    int    `json:"monto"`
	Emisor   string `json:"emisor"`
	Receptor string `json:"receptor"`
	Fecha    string `json:"fecha"`
}

type Transaction struct {
	service transacciones.Service
}

func NewTransaction(t transacciones.Service) *Transaction {
	return &Transaction{
		service: t,
	}
}

// ListTransactions godoc
// @Summary List transactions
// @Tags Transactions
// @Description get transactions
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Success 200 {object} web.Response
// @Router /transacciones [get]
func (t *Transaction) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		transactions, err := t.service.GetAll()

		if err != nil {
			fmt.Println(err)
			return
		}

		id := ctx.Query("id")
		intID, _ := strconv.Atoi(id)
		codigo := ctx.Query("codigo")
		moneda := ctx.Query("moneda")
		monto := ctx.Query("monto")
		intMonto, _ := strconv.Atoi(monto)
		emisor := ctx.Query("emisor")
		receptor := ctx.Query("receptor")
		fecha := ctx.Query("fecha")

		if id != "" || codigo != "" || moneda != "" || monto != "" || emisor != "" || receptor != "" || fecha != "" {
			var filtrados []transacciones.Transaction
			fmt.Println(moneda)
			for _, t := range transactions {
				if intID == t.ID || codigo == t.Codigo || intMonto == t.Monto || moneda == t.Moneda || emisor == t.Emisor || receptor == t.Receptor || fecha == t.Fecha {
					filtrados = append(filtrados, t)
					fmt.Println(filtrados)
				}
			}

			ctx.JSON(200, web.NewResponse(200, filtrados, ""))
			return
		}
		ctx.JSON(200, web.NewResponse(200, transactions, ""))
	}
}

// ListTransactionsWithID godoc
// @Summary List transactions with ID
// @Tags Transactions
// @Description get transactions with the given ID
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Param id path int true "ID"
// @Success 200 {object} web.Response
// @Router /transacciones/{id} [get]
func (t *Transaction) GetByID() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		intId, _ := strconv.Atoi(id)

		filtrado, err := t.service.GetByID(intId)

		if err != nil {
			ctx.JSON(404, web.NewResponse(404, nil, err.Error()))
			return
		}

		ctx.JSON(200, web.NewResponse(200, filtrado, ""))
	}
}

// StoreTransactions godoc
// @Summary Store transactions
// @Tags Transactions
// @Description store transactions
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Param transaction body request true "Transaction to store"
// @Success 200 {object} web.Response
// @Router /transacciones [post]
func (t *Transaction) Store() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("token")

		if token != os.Getenv("TOKEN") {
			c.JSON(401, web.NewResponse(401, nil, "error: no tiene permisos para realizar la peticion solicitada"))
			return
		}

		var trnsRequest request
		if err := c.ShouldBindJSON(&trnsRequest); err != nil {
			c.JSON(404, web.NewResponse(404, nil, err.Error()))
			return
		}

		camposRequeridos := []string{"Codigo", "Moneda", "Monto", "Emisor", "Receptor", "Fecha"}

		for _, campo := range camposRequeridos {
			value, err := t.service.GetField(&trnsRequest, campo)

			if err != nil {
				fmt.Println(err)
				return
			}

			if value == "" {
				c.String(400, "el campo %s es requerido", value)
				return
			}
		}

		trnsRequest.ID, _ = t.service.GenID()

		c.JSON(200, web.NewResponse(200, trnsRequest, ""))

		_, _ = t.service.Store(trnsRequest.ID, trnsRequest.Codigo, trnsRequest.Moneda, trnsRequest.Monto, trnsRequest.Emisor, trnsRequest.Receptor, trnsRequest.Fecha)
	}
}

// UpdateTransactions godoc
// @Summary Update transactions
// @Tags Transactions
// @Description update the entire transaction with the desired ID
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Param id path int true "ID"
// @Param transaction body request true "Transaction to update"
// @Success 200 {object} web.Response
// @Router /transacciones/{id} [put]
func (t *Transaction) Update() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		intId, _ := strconv.Atoi(id)

		camposRequeridos := []string{"Codigo", "Moneda", "Monto", "Emisor", "Receptor", "Fecha"}
		var modified transacciones.Transaction

		if err := c.ShouldBindJSON(&modified); err != nil {
			c.JSON(404, web.NewResponse(404, nil, err.Error()))
			return
		}

		for _, campo := range camposRequeridos {
			value, err := t.service.GetField(&modified, campo)

			if err != nil {
				fmt.Println(err)
				return
			}

			if value == nil {
				c.String(400, "el campo %s es requerido", value)
				return
			}
		}

		newT, err := t.service.Update(modified, intId)

		if err != nil {
			c.JSON(404, web.NewResponse(404, nil, err.Error()))
			return
		}

		c.JSON(200, web.NewResponse(200, newT, ""))
	}
}

// DeleteTransactions godoc
// @Summary Delete transactions
// @Tags Transactions
// @Description delete the entire transaction with the desired ID
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Param id path int true "ID"
// @Param transaction body request true "Transaction to delete"
// @Success 200 {object} web.Response
// @Router /transacciones/{id} [delete]
func (t *Transaction) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		intId, _ := strconv.Atoi(id)

		if t.service.Delete(intId) != nil {
			c.JSON(404, web.NewResponse(404, nil, t.service.Delete(intId).Error()))
			return
		}
		c.JSON(200, web.NewResponse(200, "Transaction deleted!", ""))
	}
}

func (t *Transaction) Patch() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		intId, _ := strconv.Atoi(id)

		var modified request

		if err := c.ShouldBindJSON(&modified); err != nil {
			c.JSON(404, web.NewResponse(404, nil, err.Error()))
			return
		}

		if modified.Codigo == "" {
			c.JSON(400, web.NewResponse(400, nil, "error: el codigo de la transaccion es requerido"))
			return
		}

		if modified.Monto == 0 {
			c.JSON(400, web.NewResponse(400, nil, "error: el monto de la transaccion es requerido"))
			return
		}

		newT, err := t.service.Patch(intId, modified.Codigo, modified.Monto)

		if err != nil {
			c.JSON(404, web.NewResponse(404, nil, err.Error()))
			return
		}

		c.JSON(200, web.NewResponse(200, newT, ""))
	}
}
