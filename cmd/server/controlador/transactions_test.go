package controlador

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/MiguelAngelCipamochaF/go-web/internal/transacciones"
	"github.com/MiguelAngelCipamochaF/go-web/pkg/store"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func createRequestTest(method string, url string, body string) (*http.Request, *httptest.ResponseRecorder) {
	req := httptest.NewRequest(method, url, bytes.NewBuffer([]byte(body)))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("token", "123456")

	return req, httptest.NewRecorder()
}

func createServer() *gin.Engine {
	_ = os.Setenv("TOKEN", "123456")
	db := store.New(store.FileType, "./transacciones.json")
	repo := transacciones.NewRepository(db)
	service := transacciones.NewService(repo)
	controller := NewTransaction(service)
	r := gin.Default()

	trg := r.Group("/transacciones")
	trg.PUT("/:id", controller.Update())
	trg.DELETE("/:id", controller.Delete())
	return r
}

func Test_UpdateTransaction_OK(t *testing.T) {
	r := createServer()
	req, rr := createRequestTest(http.MethodPut, "/transacciones/2", `{"ID": 2, "Codigo": "002", "Moneda": "USD",
	"Monto": 4500, "Emisor": "Miguel Cipamocha", "Receptor": "Miles Morales", "Fecha": "12-01-2022"}`)

	r.ServeHTTP(rr, req)

	assert.Equal(t, 200, rr.Code)
}

func Test_UpdateTransaction_Err(t *testing.T) {
	r := createServer()
	req, rr := createRequestTest(http.MethodPut, "/transacciones/10", `{"ID": 2, "Codigo": "002", "Moneda": "USD",
	"Monto": 4500, "Emisor": "Miguel Cipamocha", "Receptor": "Miles Morales", "Fecha": "12-01-2022"}`)

	r.ServeHTTP(rr, req)

	assert.Equal(t, 404, rr.Code)
}

func Test_DeleteTransaction_OK(t *testing.T) {
	r := createServer()
	req, rr := createRequestTest(http.MethodDelete, "/transacciones/2", "")

	r.ServeHTTP(rr, req)

	assert.Equal(t, 200, rr.Code)
}

func Test_DeleteTransaction_Err(t *testing.T) {
	r := createServer()
	req, rr := createRequestTest(http.MethodDelete, "/transacciones/10", "")

	r.ServeHTTP(rr, req)

	assert.Equal(t, 404, rr.Code)
}
