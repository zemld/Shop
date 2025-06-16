package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	httpSwagger "github.com/swaggo/http-swagger"
	"github.com/zemld/Shop/admin-service/handlers"
)

// @title Gateway
// @version 1.0
// @description Admin service for managing administrators and items in the storage.
// @host localhost:8082
// @BasePath /
func main() {
	router := chi.NewRouter()

	router.Post("/v1/admins/register", handlers.RegisterAdminHandler)
	router.Get("/v1/admins/auth", handlers.AuthentificateAdminHandler)
	router.Post("/v1/items/add", handlers.AddItemHandler)

	fs := http.FileServer(http.Dir("./docs"))
	router.Handle("/docs/*", http.StripPrefix("/docs/", fs))
	router.Get("/swagger/*", httpSwagger.Handler(httpSwagger.URL("http://localhost:8082/docs/swagger.json")))
	http.ListenAndServe(":8082", router)
}
