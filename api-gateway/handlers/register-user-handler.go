package handlers

import (
	"net/http"

	"github.com/zemld/Shop/api-gateway/domain/models"
	"github.com/zemld/Shop/api-gateway/internal"
)

// @description Registers a new user.
// @tag.name Users operations
// @param name query string true "User which you want to register"
// @param balance query float64 true "Balance of the user"
// @produce json
// @success 200 {object} models.User
// @failure 400 {object} models.StatusResponse
// @failure 500 {object} models.StatusResponse
// @router /v1/users/register [post]
func RegisterUserHandler(w http.ResponseWriter, r *http.Request) {
	response, err := internal.SendRequestToUserService(internal.POST, r.URL.Path, r.URL.Query())
	if err != nil {
		internal.WriteResponse(w, models.StatusResponse{
			Name:    r.URL.Query().Get("name"),
			Message: "Can't receive response from user service.",
		}, http.StatusInternalServerError)
		return
	}
	internal.TryParseResponseBodyAndWriteResponse(w, response, models.UserResponseType)
}
