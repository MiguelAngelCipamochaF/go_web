package main

import (
	"github.com/MiguelAngelCipamochaF/go-web/internal/transacciones"
	"github.com/MiguelAngelCipamochaF/go-web/server/controlador"
	"github.com/gin-gonic/gin"
)

func main() {

	repo := transacciones.NewRepository()
	service := transacciones.NewService(repo)
	controller := controlador.NewTransaction(service)

	router := gin.Default()
	router.GET("/hola/:name", func(c *gin.Context) {
		name := c.Param("name")
		saludo := "Hola " + name
		c.JSON(200, gin.H{
			"message": saludo,
		})
	})

	router.GET("/transacciones", controller.GetAll())
	router.POST("/transacciones/", controller.Store())
	router.GET("/transacciones/:id", controller.GetByID())
	router.Run()
}
