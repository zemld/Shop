package handlers

import (
	"net/http"

	"github.com/zemld/Shop/api-gateway/domain/models"
	"github.com/zemld/Shop/api-gateway/internal"
)

// @description Makes delivery for item.
// @tags Items
// @param secret query string true "Secret auth code for admin"
// @param name query string true "Item name which you want to deliver"
// @param amount query int true "Amount of the item to deliver"
// @produce json
// @success 200 {object} models.ItemDeliveredResponse
// @failure 400 {object} models.StatusResponse
// @failure 403 {object} models.StatusResponse
// @failure 500 {object} models.StatusResponse
// @router /v1/items/deliver [post]
func DeliverItemHandler(w http.ResponseWriter, r *http.Request) {
	response, err := internal.SendRequestToService(internal.POST, internal.AdminServiceURL, r.URL.Path, r.URL.Query(), nil)
	if err != nil {
		internal.WriteResponse(w, models.StatusResponse{
			Message: "Failed to deliver item: " + err.Error(),
		}, http.StatusInternalServerError)
		return
	}
	internal.TryParseResponseBodyAndWriteResponse(w, response, models.DeliverItemResponseType)
}
