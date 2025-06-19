package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/zemld/Shop/order-service/db"
	"github.com/zemld/Shop/order-service/internal"
	"github.com/zemld/Shop/order-service/models"
)

// @description Create a new order
// @tags Orders
// @param order body models.Order true "User whose balance you want to change"
// @produce json
// @success 200 {object} models.Order
// @failure 400 {object} models.StatusResponse
// @failure 500 {object} models.StatusResponse
// @router /v1/orders/create-order [post]
func CreateOrderHandler(w http.ResponseWriter, r *http.Request) {
	var order models.Order
	if err := json.NewDecoder(r.Body).Decode(&order); err != nil {
		internal.WriteResponse(w, models.StatusResponse{Message: "Invalid request body: " + err.Error()}, http.StatusBadRequest)
		return
	}
	database, tx, err := db.BeginTransaction(db.OutboxDB)
	if err != nil {
		internal.WriteResponse(w, models.StatusResponse{Message: "Can't connect to db: " + err.Error()}, http.StatusInternalServerError)
		return
	}
	defer db.CloseDB(database)
	encodedOrder, _ := json.Marshal(order)
	id, err := db.StoreNewOrder(tx, encodedOrder)
	if err != nil {
		internal.WriteResponse(w, models.StatusResponse{Message: "Failed to store order: " + err.Error()}, http.StatusInternalServerError)
		db.RollbackTransaction(tx)
		return
	}
	_ = id
	// TODO: положить в очередь.
	tx.Commit()
}
