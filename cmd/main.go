package main

import (
	"github.com/MiguelAngelCipamochaF/go-web/cmd/server/controlador"
	"github.com/MiguelAngelCipamochaF/go-web/internal/transacciones"
	"github.com/MiguelAngelCipamochaF/go-web/pkg/store"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()
	db := store.New(store.FileType, "./transacciones.json")
	repo := transacciones.NewRepository(db)
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

	trns := router.Group("/transacciones")
	{
		trns.GET("/", controller.GetAll())
		trns.GET("/:id", controller.GetByID())
		trns.POST("/", controller.Store())
		trns.PUT("/:id", controller.Update())
		trns.DELETE("/:id", controller.Delete())
		trns.PATCH("/:id", controller.Patch())
	}
	router.Run()
}
