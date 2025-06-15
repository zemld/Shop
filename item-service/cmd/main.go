package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	httpSwagger "github.com/swaggo/http-swagger"
	"github.com/zemld/Shop/item-service/handlers"
)

// @title Item Service
// @version 1.0
// @description API for interacting with items in storage.
// @host localhost:8083
// @BasePath /
func main() {
	router := chi.NewRouter()

	router.Post("/v1/items/add", handlers.AddItemHandler)
	router.Post("/v1/items/remove", handlers.RemoveItemHandler)

	fs := http.FileServer(http.Dir("./docs"))
	router.Handle("/docs/*", http.StripPrefix("/docs/", fs))
	router.Get("/swagger/*", httpSwagger.Handler(httpSwagger.URL("http://localhost:8083/docs/swagger.json")))
	http.ListenAndServe(":8083", router)
}
