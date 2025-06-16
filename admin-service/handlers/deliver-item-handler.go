package handlers

import (
	"net/http"

	"github.com/zemld/Shop/admin-service/domain/models"
	"github.com/zemld/Shop/admin-service/internal"
)

// @description Makes delivery for item.
// @tag.name Items operations
// @param secret query string true "Secret auth code for admin"
// @param name query string true "Item name which you want to deliver"
// @param amount query int true "Amount of the item to deliver"
// @produce json
// @success 200 {object} models.ItemDeliveredResponse
// @failure 400 {object} models.StatusResponse
// @failure 500 {object} models.StatusResponse
// @router /v1/items/deliver [post]
func DeliverItemHandler(w http.ResponseWriter, r *http.Request) {
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
			Message: "Failed to deliver item: " + err.Error(),
		}, http.StatusInternalServerError)
		return
	}
	internal.TryParseResponseBodyAndWriteResponse(w, response, models.DeliverItemResponseType)
}
