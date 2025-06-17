package handlers

import (
	"net/http"

	"github.com/zemld/Shop/item-service/db"
	"github.com/zemld/Shop/item-service/domain/models"
	"github.com/zemld/Shop/item-service/internal"
)

// @description Remove an item.
// @tags Items
// @param name query string true "Item name which you want to remove"
// @produce json
// @success 200 {object} models.ItemResponse
// @failure 400 {object} models.StatusResponse
// @failure 500 {object} models.StatusResponse
// @router /v1/items/remove [post]
func RemoveItemHandler(w http.ResponseWriter, r *http.Request) {
	itemName := r.URL.Query().Get("name")
	if itemName == "" {
		internal.WriteResponse(w, models.StatusResponse{
			Name:    itemName,
			Message: "Item name is required",
		}, http.StatusBadRequest)
		return
	}
	item := models.Item{Name: itemName}
	updatedItem, err := db.RemoveItem(db.ItemsDB, item)
	if err != nil {
		internal.WriteResponse(w, models.StatusResponse{
			Name:    item.Name,
			Message: err.Error(),
		}, http.StatusInternalServerError)
		return
	}
	internal.WriteResponse(w, models.ItemResponse{
		Item:    updatedItem,
		Message: "Item removed successfully",
	}, http.StatusOK)
}
