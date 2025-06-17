package handlers

import (
	"net/http"

	"github.com/zemld/Shop/item-service/db"
	"github.com/zemld/Shop/item-service/domain/models"
	"github.com/zemld/Shop/item-service/internal"
)

// @description Makes delivery for item.
// @tags Items
// @param name query string true "Item name which you want to deliver"
// @param amount query int true "Amount of the item to deliver"
// @produce json
// @success 200 {object} models.ItemDeliveredResponse
// @failure 400 {object} models.StatusResponse
// @failure 500 {object} models.StatusResponse
// @router /v1/items/deliver [post]
func DeliverItemHandler(w http.ResponseWriter, r *http.Request) {
	name, err := internal.ValidateItemNameFromRequest(r)
	if err != nil {
		internal.WriteResponse(w, models.StatusResponse{
			Name:    name,
			Message: "Incorrect item name",
		}, http.StatusBadRequest)
		return
	}
	amount, err := internal.ValidateItemAmountFromRequest(r)
	if err != nil {
		internal.WriteResponse(w, models.StatusResponse{
			Name:    name,
			Message: "Incorrect item amount",
		}, http.StatusBadRequest)
		return
	}
	item, err := db.ConnectToDBAndUpdateItemAmount(db.ItemsDB, name, amount)
	if err != nil {
		internal.WriteResponse(w, models.StatusResponse{
			Name:    item.Name,
			Message: err.Error(),
		}, http.StatusInternalServerError)
		return
	}
	internal.WriteResponse(w, models.ItemDeliveredResponse{
		Item:      item,
		Message:   "Item delivered",
		Delivered: amount,
	}, http.StatusOK)
}
