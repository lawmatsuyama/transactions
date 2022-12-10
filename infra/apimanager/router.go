package apimanager

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	_ "github.com/lawmatsuyama/transactions/docs"
	httpSwagger "github.com/swaggo/http-swagger"
)

// NewHandler create and return new handler
func NewHandler(transactionAPI TransactionAPI) (handler *chi.Mux) {
	handler = chi.NewRouter()
	handler.Use(middleware.Heartbeat("/transactions/v1"))
	handler.Post("/transactions/v1/save", transactionAPI.Save)
	handler.Post("/transactions/v1/get", transactionAPI.Get)
	handler.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8080/swagger/doc.json")))
	return
}
