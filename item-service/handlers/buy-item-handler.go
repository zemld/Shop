package handlers

import (
	"net/http"

	"github.com/zemld/Shop/item-service/db"
	"github.com/zemld/Shop/item-service/domain/models"
	"github.com/zemld/Shop/item-service/internal"
)

// @description Makes item purchase.
// @tags Items
// @param name query string true "Item name which you want to buy"
// @param amount query int true "Amount of the item to buy"
// @produce json
// @success 200 {object} models.ItemBoughtResponse
// @failure 400 {object} models.StatusResponse
// @failure 500 {object} models.StatusResponse
// @router /v1/items/buy [post]
func BuyItemHandler(w http.ResponseWriter, r *http.Request) {
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
	canBuyCnt := min(itemInStore.Amount, amount)
	item, err := db.UpdateItemAmount(database, name, -canBuyCnt)
	if err != nil {
		internal.WriteResponse(w, models.StatusResponse{
			Name:    item.Name,
			Message: err.Error(),
		}, http.StatusInternalServerError)
		db.RollbackTransaction(tx)
		return
	}
	db.CommitTransaction(tx)
	defer db.CloseDB(database)
	internal.WriteResponse(w, models.ItemBoughtResponse{
		Item:    item,
		Message: "Item purchased",
		Bought:  canBuyCnt,
	}, http.StatusOK)
}
