package handlers

import (
	"net/http"

	"github.com/zemld/Shop/api-gateway/domain/models"
	"github.com/zemld/Shop/api-gateway/internal"
)

// @description Adds item to store.
// @tags Items
// @param secret query string true "Secret auth code"
// @param name query string true "Item name which you want to add"
// @param price query float64 true "Cost of the item"
// @param amount query int true "Amount of the item in stock"
// @produce json
// @success 200 {object} models.ItemResponse
// @failure 400 {object} models.StatusResponse
// @failure 403 {object} models.StatusResponse
// @failure 500 {object} models.StatusResponse
// @router /v1/items/add [post]
func AddItemHandler(w http.ResponseWriter, r *http.Request) {
	response, err := internal.SendRequestToService(internal.POST, internal.AdminServiceURL, r.URL.Path, r.URL.Query(), nil)
	if err != nil {
		internal.WriteResponse(w, models.StatusResponse{
			Message: "Failed to add item: " + err.Error(),
		}, http.StatusInternalServerError)
		return
	}
	internal.TryParseResponseBodyAndWriteResponse(w, response, models.AddItemResponseType)
}
