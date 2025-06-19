package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/zemld/Shop/order-service/internal"
	"github.com/zemld/Shop/order-service/models"
	"github.com/zemld/Shop/order-service/mq"
)

// @description Create a new order
// @tags Orders
// @param order body models.Order true "User whose balance you want to change"
// @produce json
// @success 200 {object} models.Order
// @failure 400 {object} models.OrderStatusResponse
// @failure 500 {object} models.OrderStatusResponse
// @router /v1/orders/create-order [post]
func CreateOrderHandler(w http.ResponseWriter, r *http.Request) {
	var order models.Order
	if err := json.NewDecoder(r.Body).Decode(&order); err != nil {
		internal.WriteResponse(w, models.OrderStatusResponse{Message: "Invalid request body: " + err.Error()}, http.StatusBadRequest)
		return
	}
	defer r.Body.Close()
	nc, err := internal.SendMsg(order)
	if err != nil {
		internal.WriteResponse(w, models.OrderStatusResponse{Message: fmt.Sprintf("Can't handle new order: %s", err.Error())}, http.StatusInternalServerError)
		return
	}
	defer nc.Close()
	msg, err := internal.WaitForMessage(nc, mq.OrderHandled, mq.DefaultTimeout)
	if err != nil {
		internal.WriteResponse(w, models.OrderStatusResponse{Message: fmt.Sprintf("Can't handle new order: %s", err.Error())}, http.StatusInternalServerError)
		return
	}
	var createdOrder models.Order
	if err := json.Unmarshal(msg, &createdOrder); err != nil {
		internal.WriteResponse(w, models.OrderStatusResponse{Message: fmt.Sprintf("Can't handle new order: %s", err.Error())}, http.StatusInternalServerError)
		return
	}
	internal.WriteResponse(w, createdOrder, http.StatusOK)
}
