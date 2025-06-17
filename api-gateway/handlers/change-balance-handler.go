package handlers

import (
	"net/http"

	"github.com/zemld/Shop/api-gateway/domain/models"
	"github.com/zemld/Shop/api-gateway/internal"
)

// @description Changes the balance of a registered user.
// @tags Users
// @param name query string true "User whose balance you want to change"
// @param balance query float64 true "New balance of the user"
// @produce json
// @success 200 {object} models.User
// @failure 400 {object} models.StatusResponse
// @failure 500 {object} models.StatusResponse
// @router /v1/users/change-balance [post]
func ChangeBalanceHandler(w http.ResponseWriter, r *http.Request) {
	response, err := internal.SendRequestToUserService(internal.POST, r.URL.Path, r.URL.Query())
	if err != nil {
		internal.WriteResponse(w, models.StatusResponse{
			Name:    "",
			Message: "Can't receive response from user service.",
		}, http.StatusInternalServerError)
		return
	}
	internal.TryParseResponseBodyAndWriteResponse(w, response, models.UserResponseType)
}
