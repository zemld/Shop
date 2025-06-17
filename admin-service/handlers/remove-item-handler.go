package handlers

import (
	"net/http"

	"github.com/zemld/Shop/admin-service/domain/models"
	"github.com/zemld/Shop/admin-service/internal"
)

// @description Remove an item.
// @tags Items
// @param secret query string true "Admin secret code for authorization"
// @param name query string true "Item name which you want to remove"
// @produce json
// @success 200 {object} models.ItemResponse
// @failure 400 {object} models.StatusResponse
// @failure 403 {object} models.StatusResponse
// @failure 500 {object} models.StatusResponse
// @router /v1/items/remove [post]
func RemoveItemHandler(w http.ResponseWriter, r *http.Request) {
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
			Message: "Failed to remove item: " + err.Error(),
		}, http.StatusInternalServerError)
		return
	}
	internal.TryParseResponseBodyAndWriteResponse(w, response, models.ItemResponseType)
}
