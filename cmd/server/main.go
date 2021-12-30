package main

import (
	"os"

	"github.com/MiguelAngelCipamochaF/go-web/cmd/server/controlador"
	"github.com/MiguelAngelCipamochaF/go-web/docs"
	"github.com/MiguelAngelCipamochaF/go-web/internal/transacciones"
	"github.com/MiguelAngelCipamochaF/go-web/pkg/store"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

// @title MELI Bootcamp GO API
// @version 1.0
// @description This API Handle MELI Transactions.
// @termsOfService https://developers.mercadolibre.com.ar/es_ar/terminos-y-condiciones
// @contact.name API Support
// @contact.url https://developers.mercadolibre.com.ar/support
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
func main() {
	_ = godotenv.Load()
	db := store.New(store.FileType, "./transacciones.json")
	repo := transacciones.NewRepository(db)
	service := transacciones.NewService(repo)
	controller := controlador.NewTransaction(service)
	router := gin.Default()

	docs.SwaggerInfo.Host = os.Getenv("HOST")
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

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
