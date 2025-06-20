package handlers

import (
	"log"
	"net/http"

	"github.com/zemld/Shop/api-gateway/domain/constants"
	"github.com/zemld/Shop/api-gateway/domain/models"
	"github.com/zemld/Shop/api-gateway/internal"
)

// @description Checks if user is registered.
// @tags Users
// @param user path string true "User which you want to check"
// @produce json
// @success 200 {object} models.UserRegistered
// @failure 400 {object} models.StatusResponse
// @failure 500 {object} models.StatusResponse
// @router /v1/users/{user} [get]
func CheckUserRegisteredHanlder(w http.ResponseWriter, r *http.Request) {
	user, err := internal.TryParseURLPathParam(r.URL.Path, constants.UserRegisteredPath, constants.PathParam)
	if err != nil {
		log.Println("Can't parse user from request URL")
		internal.WriteResponse(
			w,
			models.StatusResponse{
				Name:    "",
				Message: "Can't parse user from request URL"},
			http.StatusBadRequest)
		return
	}
	log.Printf("Parsed user: %s\n", user)
	response, err := internal.SendRequestToService(internal.GET, internal.UserServiceURL, r.URL.Path, nil)
	if err != nil {
		internal.WriteResponse(w, models.StatusResponse{
			Name:    user,
			Message: "Can't recieve response from user service.",
		}, http.StatusInternalServerError)
		return
	}
	internal.TryParseResponseBodyAndWriteResponse(w, response, models.UserRegisteredResponseType)
}
