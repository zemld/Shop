package handlers

import (
	"net/http"

	"github.com/zemld/Shop/api-gateway/domain/models"
	"github.com/zemld/Shop/api-gateway/internal"
)

// @description Checks if user is admin.
// @tag.name Admin operations
// @param code query string true "Secret auth code"
// @produce json
// @success 200 {object} models.IsAdminResponse
// @failure 400 {object} models.StatusResponse
// @failure 403 {object} models.StatusResponse
// @failure 500 {object} models.StatusResponse
// @router /v1/admins/auth [get]
func AuthentificateAdminHandler(w http.ResponseWriter, r *http.Request) {
	response, err := internal.SendRequestToAdminService(internal.GET, r.URL.Path, r.URL.Query())
	if err != nil {
		internal.WriteResponse(w, models.StatusResponse{
			Message: "Failed to authenticate admin: " + err.Error(),
		}, http.StatusInternalServerError)
		return
	}
	internal.TryParseResponseBodyAndWriteResponse(w, response, models.IsAdminResponseType)
}
