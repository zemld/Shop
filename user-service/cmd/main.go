package cmd

import (
	"net/http"

	"github.com/go-chi/chi"
)

func main() {
	router := chi.NewRouter()

	http.ListenAndServe(":8081", router)
}
