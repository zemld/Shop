package handlers

import (
	"net/http"

	"github.com/zemld/Shop/api-gateway/domain/models"
	"github.com/zemld/Shop/api-gateway/internal"
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
	response, err := internal.SendRequestToService(internal.POST, internal.OrderServiceURL, r.URL.Path, r.URL.Query())
	if err != nil {
		internal.WriteResponse(w, models.OrderStatusResponse{
			Message: "Failed to create order: " + err.Error(),
		}, http.StatusInternalServerError)
		return
	}
	internal.TryParseResponseBodyAndWriteResponse(w, response, models.OrderResponseType)
}
