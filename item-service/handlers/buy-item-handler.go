package handlers

import (
	"net/http"

	"github.com/zemld/Shop/item-service/db"
	"github.com/zemld/Shop/item-service/domain/models"
	"github.com/zemld/Shop/item-service/internal"
)

// @description Makes item purchase.
// @tag.name Items operations
// @param name query string true "Item name which you want to buy"
// @param amount query int true "Amount of the item to buy"
// @produce json
// @success 200 {object} models.ItemResponse
// @failure 400 {object} models.ItemResponse
// @failure 500 {object} models.ItemResponse
// @router /v1/items/buy [post]
func BuyItemHandler(w http.ResponseWriter, r *http.Request) {
	name, amount, err := internal.ValidateItemNameAndAmountFromRequest(r)
	if err != nil {
		internal.WriteResponse(w, models.ItemResponse{
			Item:    models.Item{},
			Message: err.Error(),
		}, http.StatusBadRequest)
		return
	}
	item, err := db.UpdateItemAmount(db.ItemsDB, name, -amount)
	if err != nil {
		internal.WriteResponse(w, models.ItemResponse{
			Item:    item,
			Message: err.Error(),
		}, http.StatusInternalServerError)
		return
	}
	internal.WriteResponse(w, models.ItemResponse{
		Item:    item,
		Message: "Item purchased",
	}, http.StatusOK)
}
