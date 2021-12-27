package controlador

import (
	"fmt"
	"strconv"

	"github.com/MiguelAngelCipamochaF/go-web/internal/transacciones"
	"github.com/gin-gonic/gin"
)

type Request struct {
}

type Transaction struct {
	service transacciones.Service
}

func NewTransaction(t transacciones.Service) *Transaction {
	return &Transaction{
		service: t,
	}
}

func (t *Transaction) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		transactions := t.service.GetAll()
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

			for _, t := range transactions {
				if intID == t.ID || codigo == t.Codigo || intMonto == t.Monto || moneda == t.Moneda || emisor == t.Emisor || receptor == t.Receptor || fecha == t.Fecha {
					filtrados = append(filtrados, t)
				}
			}

			ctx.JSON(200, filtrados)
			return
		}
		ctx.JSON(200, transactions)
	}
}

func (t *Transaction) GetByID() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		intId, _ := strconv.Atoi(id)

		filtrado := t.service.GetByID(intId)

		if filtrado != nil {
			ctx.JSON(200, filtrado)
		} else {
			ctx.JSON(404, gin.H{
				"message": "Error",
			})
		}
	}
}

func (t *Transaction) Store() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("token")

		if token != "123456" {
			c.JSON(401, gin.H{
				"error": "no tiene permisos para realizar la peticion solicitada",
			})
			return
		}

		var trnsRequest transacciones.Transaction
		if err := c.ShouldBindJSON(&trnsRequest); err != nil {
			c.JSON(404, gin.H{
				"error": err.Error(),
			})
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

		trnsRequest.ID = t.service.GenID()
		c.JSON(200, trnsRequest)

		t.service.Store(trnsRequest)
	}
}

func (t *Transaction) Update() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		intId, _ := strconv.Atoi(id)

		camposRequeridos := []string{"Codigo", "Moneda", "Monto", "Emisor", "Receptor", "Fecha"}
		var modified transacciones.Transaction

		if err := c.ShouldBindJSON(&modified); err != nil {
			c.JSON(404, gin.H{
				"error": err.Error(),
			})
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
			c.JSON(404, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(200, newT)
	}
}

func (t *Transaction) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		intId, _ := strconv.Atoi(id)

		if t.service.Delete(intId) != nil {
			c.JSON(404, gin.H{
				"error": t.service.Delete(intId).Error(),
			})
			return
		}
		c.JSON(200, gin.H{
			"success": true,
		})
	}
}

func (t *Transaction) Patch() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		intId, _ := strconv.Atoi(id)

		var modified transacciones.Transaction

		if err := c.ShouldBindJSON(&modified); err != nil {
			c.JSON(404, gin.H{
				"error": err.Error(),
			})
			return
		}

		if modified.Codigo == "" {
			c.JSON(400, gin.H{
				"error": "el codigo de la transaccion es requerido",
			})
			return
		}

		if modified.Monto == 0 {
			c.JSON(400, gin.H{
				"error": "el codigo de la transaccion es requerido",
			})
			return
		}

		newT, err := t.service.Patch(intId, modified.Codigo, modified.Monto)

		if err != nil {
			c.JSON(404, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(200, newT)
	}
}
