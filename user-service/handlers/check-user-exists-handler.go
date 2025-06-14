package handlers

import (
	"log"
	"net/http"

	"github.com/zemld/Shop/user-service/db"
	"github.com/zemld/Shop/user-service/domain/constants"
	"github.com/zemld/Shop/user-service/domain/dto"
	"github.com/zemld/Shop/user-service/internal"
)

// @description Checks if user is registered.
// @tag.name File operations
// @param user path string true "User which you want to check"
// @produce json
// @success 200 {object} dto.UserRegistered
// @failure 400 {object} dto.StatusResponse
// @failure 500 {object} dto.StatusResponse
// @router /v1/users/{user} [get]
func CheckUserRegistered(w http.ResponseWriter, r *http.Request) {
	user, err := internal.TryParseURLPathParam(r.URL.Path, constants.UserRegisteredPath, constants.PathParam)
	if err != nil {
		log.Println("Can't parse user from request URL")
		internal.WriteResponse(
			w,
			dto.StatusResponse{
				User:    "",
				Message: "Can't parse user from request URL"},
			http.StatusBadRequest)
		return
	}
	log.Printf("Parsed user: %s\n", user)
	doesExist, err := db.CreateDBConnectionAndCheckUserRegistered(db.UsersDB, db.CheckUserRegisteredQuery)
	if err != nil {
		internal.WriteResponse(w, dto.StatusResponse{
			User:    user,
			Message: "Error while checking user registration: " + err.Error(),
		}, http.StatusInternalServerError)
		return
	}
	userRegisteredResponse := dto.UserRegistered{
		User:         user,
		IsRegistered: doesExist,
	}
	internal.WriteResponse(w, userRegisteredResponse, http.StatusOK)
}
