package handlers

import (
	"net/http"

	"github.com/zemld/Shop/admin-service/domain/models"
	"github.com/zemld/Shop/admin-service/internal"
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
	isAdmin := internal.AuthAdmin(r.URL.Query().Get("secret"))
	if !isAdmin {
		internal.WriteResponse(w, models.StatusResponse{
			Message: "You are not authorized to add items.",
		}, http.StatusForbidden)
		return
	}
	response, err := internal.SendRequestToItemService(internal.POST, r.URL.Path, r.URL.Query())
	if err != nil {
		internal.WriteResponse(w, models.StatusResponse{
			Message: "Failed to add item: " + err.Error(),
		}, http.StatusInternalServerError)
		return
	}
	internal.TryParseResponseBodyAndWriteResponse(w, response, models.AddItemResponseType)
}
