package handlers

import (
	"net/http"

	"github.com/zemld/Shop/api-gateway/domain/models"
	"github.com/zemld/Shop/api-gateway/internal"
)

// @description Registers a new admin.
// @tags Admins
// @param name query string true "Admin which you want to register"
// @produce json
// @success 200 {object} models.Admin
// @failure 400 {object} models.StatusResponse
// @failure 500 {object} models.StatusResponse
// @router /v1/admins/register [post]
func RegisterAdminHandler(w http.ResponseWriter, r *http.Request) {
	response, err := internal.SendRequestToService(internal.POST, internal.AdminServiceURL, r.URL.Path, r.URL.Query(), nil)
	if err != nil {
		internal.WriteResponse(w, models.StatusResponse{
			Message: "Failed to register admin: " + err.Error(),
		}, http.StatusInternalServerError)
		return
	}
	internal.TryParseResponseBodyAndWriteResponse(w, response, models.AdminType)
}
