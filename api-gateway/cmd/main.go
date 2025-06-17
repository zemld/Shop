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

	// TODO: настроить теги
	router.Get("/v1/users/{user}", handlers.CheckUserRegisteredHanlder)
	router.Post("/v1/users/register", handlers.RegisterUserHandler)
	router.Post("/v1/users/change-balance", handlers.ChangeBalanceHandler)
	router.Get("/v1/users/balance", handlers.GetUserBalanceHandler)

	router.Post("/v1/admins/register", handlers.RegisterAdminHandler)
	router.Get("/v1/admins/auth", handlers.AuthentificateAdminHandler)

	router.Post("/v1/items/add", handlers.AddItemHandler)
	router.Post("/v1/items/buy", handlers.BuyItemHandler)
	router.Post("/v1/items/deliver", handlers.DeliverItemHandler)
	router.Post("/v1/items/remove", handlers.RemoveItemHandler)
	router.Post("/v1/items/update-price", handlers.UpdateItemPriceHandler)

	fs := http.FileServer(http.Dir("./docs"))
	router.Handle("/docs/*", http.StripPrefix("/docs/", fs))
	router.Get("/swagger/*", httpSwagger.Handler(httpSwagger.URL("http://localhost:8080/docs/swagger.json")))
	http.ListenAndServe(":8080", router)
}
