package apimanager

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/rs/cors"

	httpSwagger "github.com/swaggo/http-swagger"
)

// NewHandler create and return new handler
func NewHandler(transactionAPI TransactionAPI) (handler *chi.Mux) {
	handler = chi.NewRouter()

	handler.Use(cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
		AllowedMethods:   []string{"POST", "OPTIONS", "GET"},
		AllowedHeaders:   []string{"Accept", "Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization"},
	}).Handler)

	handler.Use(middleware.Heartbeat("/transactions/v1"))
	handler.Post("/transactions/v1/save", transactionAPI.Save)
	handler.Post("/transactions/v1/get", transactionAPI.Get)
	handler.Get("/transactions/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("http://127.0.0.1:8080/transactions/swagger/doc.json")))
	return
}
