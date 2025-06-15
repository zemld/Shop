package handlers

import (
	"net/http"

	"github.com/zemld/Shop/api-gateway/domain/dto"
	"github.com/zemld/Shop/api-gateway/internal"
)

// @description Changes the balance of a registered user.
// @tag.name Users operations
// @param name query string true "User whose balance you want to change"
// @param balance query float64 true "New balance of the user"
// @produce json
// @success 200 {object} models.User
// @failure 400 {object} dto.StatusResponse
// @failure 500 {object} dto.StatusResponse
// @router /v1/users/change-balance [post]
func ChangeBalanceHandler(w http.ResponseWriter, r *http.Request) {
	response, err := internal.SendRequestToUserService(internal.POST, r.URL.Path, r.URL.Query())
	if err != nil {
		internal.WriteResponse(w, dto.StatusResponse{
			User:    "",
			Message: "Can't receive response from user service.",
		}, http.StatusInternalServerError)
		return
	}
	internal.TryParseResponseBodyAndWriteResponse(w, response, dto.UserResponseType)
}
