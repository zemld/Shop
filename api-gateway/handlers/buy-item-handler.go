package handlers

import (
	"net/http"

	"github.com/zemld/Shop/api-gateway/domain/models"
	"github.com/zemld/Shop/api-gateway/internal"
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
	response, err := internal.SendRequestToAdminService(internal.POST, r.URL.Path, r.URL.Query())
	if err != nil {
		internal.WriteResponse(w, models.StatusResponse{
			Message: "Failed to buy item: " + err.Error(),
		}, http.StatusInternalServerError)
		return
	}
	internal.TryParseResponseBodyAndWriteResponse(w, response, models.PurchaseItemResponseType)
}
