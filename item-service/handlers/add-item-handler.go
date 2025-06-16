package handlers

import (
	"net/http"

	"github.com/zemld/Shop/item-service/db"
	"github.com/zemld/Shop/item-service/domain/models"
	"github.com/zemld/Shop/item-service/internal"
)

// @description Add a new item.
// @tag.name Items operations
// @param name query string true "Item name which you want to add"
// @param price query float64 true "Cost of the item"
// @param amount query int true "Amount of the item in stock"
// @produce json
// @success 200 {object} models.ItemResponse
// @failure 400 {object} models.StatusResponse
// @failure 500 {object} models.StatusResponse
// @router /v1/items/add [post]
func AddItemHandler(w http.ResponseWriter, r *http.Request) {
	item, err := internal.ValidateItemFromRequest(r)
	if err != nil {
		internal.WriteResponse(w, models.StatusResponse{
			Name:    item.Name,
			Message: err.Error(),
		}, http.StatusBadRequest)
		return
	}
	updatedItem, err := db.AddItem(db.ItemsDB, item)
	if err != nil {
		internal.WriteResponse(w, models.StatusResponse{
			Name:    item.Name,
			Message: err.Error(),
		}, http.StatusInternalServerError)
		return
	}
	internal.WriteResponse(w, models.ItemResponse{
		Item:    updatedItem,
		Message: "Item added successfully",
	}, http.StatusOK)
}
