package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	httpSwagger "github.com/swaggo/http-swagger"
	"github.com/zemld/Shop/order-service/handlers"
)

// @title Orders Service
// @version 1.0
// @description API for creating orders.
// @host localhost:8084
// @BasePath /
func main() {
	router := chi.NewRouter()

	router.Post("/v1/orders/create-order", handlers.CreateOrderHandler)

	fs := http.FileServer(http.Dir("./docs"))
	router.Handle("/docs/*", http.StripPrefix("/docs/", fs))
	router.Get("/swagger/*", httpSwagger.Handler(httpSwagger.URL("http://localhost:8084/docs/swagger.json")))

	http.ListenAndServe(":8084", router)
}
