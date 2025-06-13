package main

import (
	"net/http"

	"github.com/go-chi/chi"
	httpSwagger "github.com/swaggo/http-swagger"
	"github.com/zemld/Shop/api-gateway/handlers"
)

// @title Gateway
// @version 1.0
// @description API gateway for Shop application.
// @host localhost:8080
// @BasePath /
func main() {
	router := chi.NewRouter()

	router.Get("/v1/users/{userID}", handlers.CheckUserRegistered)

	fs := http.FileServer(http.Dir("./docs"))
	router.Handle("/docs/*", http.StripPrefix("/docs/", fs))
	router.Get("/swagger/*", httpSwagger.Handler(httpSwagger.URL("http://localhost:8080/docs/swagger.json")))
	http.ListenAndServe(":8080", router)
}
