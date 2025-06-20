package handlers

import (
	"net/http"

	"github.com/zemld/Shop/api-gateway/domain/models"
	"github.com/zemld/Shop/api-gateway/internal"
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
	response, err := internal.SendRequestToService(internal.POST, internal.AdminServiceURL, r.URL.Path, r.URL.Query())
	if err != nil {
		internal.WriteResponse(w, models.StatusResponse{
			Message: "Failed to remove item: " + err.Error(),
		}, http.StatusInternalServerError)
		return
	}
	internal.TryParseResponseBodyAndWriteResponse(w, response, models.ItemResponseType)
}
