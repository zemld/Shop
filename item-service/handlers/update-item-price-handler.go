package handlers

import (
	"net/http"

	"github.com/zemld/Shop/item-service/db"
	"github.com/zemld/Shop/item-service/domain/models"
	"github.com/zemld/Shop/item-service/internal"
)

// @description Change price for item.
// @tags Items
// @param name query string true "Item name which you want to change price for"
// @param price query number true "New price of the item"
// @produce json
// @success 200 {object} models.ItemWithNewPriceResponse
// @failure 400 {object} models.StatusResponse
// @failure 500 {object} models.StatusResponse
// @router /v1/items/update-price [post]
func UpdateItemPriceHandler(w http.ResponseWriter, r *http.Request) {
	name, err := internal.ValidateItemNameFromRequest(r)
	if err != nil {
		internal.WriteResponse(w, models.StatusResponse{
			Name:    name,
			Message: "Incorrect item name",
		}, http.StatusBadRequest)
		return
	}
	newPrice, err := internal.ValidateItemPriceFromRequest(r)
	if err != nil {
		internal.WriteResponse(w, models.StatusResponse{
			Name:    name,
			Message: "Incorrect item price",
		}, http.StatusBadRequest)
		return
	}
	database, tx, err := db.BeginTransaction(db.ItemsDB)
	if err != nil {
		internal.WriteResponse(w, models.StatusResponse{
			Name:    name,
			Message: "Failed to start transaction: " + err.Error(),
		}, http.StatusInternalServerError)
		return
	}
	defer db.CloseDB(database)
	itemInStore, err := db.SelectItem(database, name)
	if err != nil {
		internal.WriteResponse(w, models.StatusResponse{
			Name:    name,
			Message: "Failed to select item: " + err.Error(),
		}, http.StatusInternalServerError)
		db.RollbackTransaction(tx)
		return
	}
	oldPrice := itemInStore.Price
	updatedItem, err := db.UpdateItemPrice(database, itemInStore, newPrice)
	if err != nil {
		internal.WriteResponse(w, models.StatusResponse{
			Name:    itemInStore.Name,
			Message: err.Error(),
		}, http.StatusInternalServerError)
		db.RollbackTransaction(tx)
		return
	}
	internal.WriteResponse(w, models.ItemWithNewPriceResponse{
		Item:     updatedItem,
		Message:  "Item price updated successfully",
		OldPrice: oldPrice,
	}, http.StatusOK)
	db.CommitTransaction(tx)
}
