package main

import (
	"github.com/MiguelAngelCipamochaFigueredo/go-web/internal/transacciones"
	"github.com/gin-gonic/gin"
)

func main() {

	repo := transacciones.NewRepository()
	service := transacciones.NewService(repo)
	controller := transacciones.NewController(service)

	router := gin.Default()
	router.GET("/hola/:name", func(c *gin.Context) {
		name := c.Param("name")
		saludo := "Hola " + name
		c.JSON(200, gin.H{
			"message": saludo,
		})
	})

	router.GET("/transacciones", controller.GetAll)
	// router.POST("/transacciones/", Guardar())
	// router.GET("/transacciones/:id", GetByID)
	router.Run()
}

// func GetAll(c *gin.Context) {
// 	id := c.Query("id")
// 	intID, _ := strconv.Atoi(id)
// 	codigo := c.Query("codigo")
// 	moneda := c.Query("moneda")
// 	monto := c.Query("monto")
// 	intMonto, _ := strconv.Atoi(monto)
// 	emisor := c.Query("emisor")
// 	receptor := c.Query("receptor")
// 	fecha := c.Query("fecha")

// 	if id != "" || codigo != "" || moneda != "" || monto != "" || emisor != "" || receptor != "" || fecha != "" {
// 		var filtrados []Transaction

// 		for _, t := range transacciones {
// 			if intID == t.ID || codigo == t.Codigo || intMonto == t.Monto || moneda == t.Moneda || emisor == t.Emisor || receptor == t.Receptor || fecha == t.Fecha {
// 				filtrados = append(filtrados, t)
// 			}
// 		}

// 		c.JSON(200, filtrados)
// 		return
// 	}

// 	c.JSON(200, transacciones)
// }

// func GetByID(c *gin.Context) {
// 	id := c.Param("id")
// 	intId, _ := strconv.Atoi(id)
// 	var filtrado *Transaction

// 	for _, t := range transacciones {
// 		fmt.Println(t.ID)
// 		if intId == t.ID {
// 			filtrado = &t
// 			break
// 		}
// 	}

// 	if filtrado != nil {
// 		c.JSON(200, filtrado)
// 	} else {
// 		c.JSON(404, gin.H{
// 			"message": "Error",
// 		})
// 	}

// }

// func getField(v interface{}, name string) (interface{}, error) {
// 	rv := reflect.ValueOf(v)
// 	if rv.Kind() != reflect.Ptr || rv.Elem().Kind() != reflect.Struct {
// 		return nil, errors.New("v debe ser un puntero a una estructura")
// 	}

// 	rv = rv.Elem()

// 	fv := rv.FieldByName(name)

// 	if !fv.IsValid() {
// 		return nil, fmt.Errorf("%s no existe en la estructura\n", name)
// 	}

// 	if !fv.CanSet() {
// 		return nil, fmt.Errorf("no es posible acceder al campo %s\n", name)
// 	}

// 	if fv.IsZero() {
// 		return nil, fmt.Errorf("el campo %s esta vacio\n", name)
// 	}

// 	return fv, nil
// }

// func Guardar() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		token := c.GetHeader("token")

// 		if token != "123456" {
// 			c.JSON(401, gin.H{
// 				"error": "no tiene permisos para realizar la peticion solicitada",
// 			})
// 			return
// 		}

// 		var trnsRequest Transaction
// 		if err := c.ShouldBindJSON(&trnsRequest); err != nil {
// 			c.JSON(404, gin.H{
// 				"error": err.Error(),
// 			})
// 			return
// 		}

// 		camposRequeridos := []string{"Codigo", "Moneda", "Monto", "Emisor", "Receptor", "Fecha"}

// 		for _, campo := range camposRequeridos {
// 			value, err := getField(&trnsRequest, campo)

// 			if err != nil {
// 				fmt.Println(err)
// 				return
// 			}

// 			if value == "" {
// 				c.String(400, "el campo %s es requerido", value)
// 				return
// 			}
// 		}

// 		trnsRequest.ID = genID()
// 		c.JSON(200, trnsRequest)

// 		transacciones = append(transacciones, trnsRequest)
// 	}
// }

// func genID() int {
// 	lastId := 0

// 	if len(transacciones) > 0 {
// 		lastId = transacciones[len(transacciones)-1].ID
// 	}

// 	return lastId + 1
// }
