package handlers

import (
	"net/http"

	"github.com/zemld/Shop/user-service/db"
	"github.com/zemld/Shop/user-service/domain/dto"
	"github.com/zemld/Shop/user-service/internal"
)

// @description Registers a new user.
// @tags Users
// @param name query string true "User which you want to register"
// @param balance query float64 true "Balance of the user"
// @produce json
// @success 200 {object} models.User
// @failure 400 {object} dto.StatusResponse
// @failure 500 {object} dto.StatusResponse
// @router /v1/users/register [post]
func RegisterUserHandler(w http.ResponseWriter, r *http.Request) {
	user, err := internal.ValidateUserFromRequest(r)
	if err != nil {
		internal.WriteResponse(
			w,
			dto.StatusResponse{
				User:    user.Name,
				Message: "You must provide a valid name and balance for the user.",
			},
			http.StatusBadRequest)
		return
	}

	err = db.CreateDBConnectionAndRegisterUser(db.UsersDB, user.Name, user.Balance)
	if err != nil {
		internal.WriteResponse(
			w,
			dto.StatusResponse{
				User:    user.Name,
				Message: "Error while registering user: " + err.Error(),
			},
			http.StatusInternalServerError)
		return
	}
	internal.WriteResponse(w, user, http.StatusOK)
}
