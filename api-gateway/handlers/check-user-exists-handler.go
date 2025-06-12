package handlers

import (
	"log"
	"net/http"

	"github.com/zemld/Shop/api-gateway/domain/constants"
	"github.com/zemld/Shop/api-gateway/domain/dto"
	"github.com/zemld/Shop/api-gateway/internal"
)

// @description Checks if user is registered.
// @tag.name File operations
// @param userID path int true "User ID which you want to check"
// @produce json
// @success 200 {object} dto.UserRegistered
// @failure 400 {object} dto.StatusResponse
// @failure 500 {object} dto.StatusResponse
// @router /v1/users/{userID} [get]
func CheckUserRegistered(w http.ResponseWriter, r *http.Request) {
	userID, err := internal.TryParseURLPathParamAndConvertToInt(r.URL.Path, constants.UserRegisteredPath, constants.IDPathParam)
	if err != nil {
		log.Println("Can't parse user ID from request URL")
		internal.WriteResponse(
			w,
			dto.StatusResponse{
				UserID:  -1,
				Message: "Can't parse user ID from request URL"},
			http.StatusBadRequest)
		return
	}
	log.Printf("Parsed user ID: %d\n", userID)
	userRegisteredResponse := dto.UserRegistered{
		UserID:       userID,
		IsRegistered: true, // TODO: написать запрос в user-service.
	}
	internal.WriteResponse(w, userRegisteredResponse, http.StatusOK)
}
