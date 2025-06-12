package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/zemld/Shop/api-gateway/handlers"
)

func main() {
	router := chi.NewRouter()

	router.Get("/users/{id}", handlers.CheckUserExistsHandler)

	http.ListenAndServe(":8080", router)
}
