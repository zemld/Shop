package handlers

import (
	"net/http"

	"github.com/zemld/Shop/api-gateway/domain/models"
	"github.com/zemld/Shop/api-gateway/internal"
)

// @description Change price for item.
// @tag.name Items operations
// @param secret query string true "Admin secret code for authorization"
// @param name query string true "Item name which you want to change price for"
// @param price query number true "New price of the item"
// @produce json
// @success 200 {object} models.ItemWithNewPriceResponse
// @failure 400 {object} models.StatusResponse
// @failure 403 {object} models.StatusResponse
// @failure 500 {object} models.StatusResponse
// @router /v1/items/update-price [post]
func UpdateItemPriceHandler(w http.ResponseWriter, r *http.Request) {
	response, err := internal.SendRequestToAdminService(internal.POST, r.URL.Path, r.URL.Query())
	if err != nil {
		internal.WriteResponse(w, models.StatusResponse{
			Message: "Failed to update item price: " + err.Error(),
		}, http.StatusInternalServerError)
		return
	}

	internal.TryParseResponseBodyAndWriteResponse(w, response, models.UpdatePriceResponseType)
}
