package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/zemld/Shop/order-service/handlers"
)

func main() {
	router := chi.NewRouter()

	router.Post("/v1/orders/create-order", handlers.CreateOrderHandler)

	http.ListenAndServe(":8084", router)
}
