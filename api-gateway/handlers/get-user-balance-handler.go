package handlers

import (
	"net/http"

	"github.com/zemld/Shop/api-gateway/domain/dto"
	"github.com/zemld/Shop/api-gateway/internal"
)

// @description Returns user's balance.
// @tag.name Users operations
// @param user query string true "User whose balance you want to get"
// @produce json
// @success 200 {object} models.User
// @failure 400 {object} dto.StatusResponse
// @failure 500 {object} dto.StatusResponse
// @router /v1/users/balance [get]
func GetUserBalanceHandler(w http.ResponseWriter, r *http.Request) {
	response, err := internal.SendRequestToUserService(internal.GET, r.URL.Path, r.URL.Query())
	if err != nil {
		internal.WriteResponse(w, dto.StatusResponse{
			User:    r.URL.Query().Get("user"),
			Message: "Can't receive response from user service.",
		}, http.StatusInternalServerError)
		return
	}
	internal.TryParseResponseBodyAndWriteResponse(w, response, dto.UserResponseType)
}
