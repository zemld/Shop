package main

import (
	"net/http"

	"github.com/go-chi/chi"
	httpSwagger "github.com/swaggo/http-swagger"
	"github.com/zemld/Shop/user-service/handlers"
)

// @title User Service
// @version 1.0
// @description API for interacting with users.
// @host localhost:8081
// @BasePath /
func main() {
	router := chi.NewRouter()

	router.Get("/v1/users/{userID}", handlers.CheckUserRegistered)
	router.Post("/v1/users/register", handlers.RegisterUserHandler)
	router.Post("/v1/users/change-balance", handlers.ChangeBalanceHandler)
	router.Get("/v1/users/balance", handlers.GetUserBalanceHandler)

	fs := http.FileServer(http.Dir("./docs"))
	router.Handle("/docs/*", http.StripPrefix("/docs/", fs))
	router.Get("/swagger/*", httpSwagger.Handler(httpSwagger.URL("http://localhost:8081/docs/swagger.json")))
	http.ListenAndServe(":8081", router)
}
