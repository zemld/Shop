package handlers

import (
	"net/http"

	"github.com/zemld/Shop/user-service/db"
	"github.com/zemld/Shop/user-service/domain/dto"
	"github.com/zemld/Shop/user-service/domain/models"
	"github.com/zemld/Shop/user-service/internal"
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
	doesRegistered, err := db.CreateDBConnectionAndCheckUserRegistered(db.UsersDB, user.Name)
	if err != nil || !doesRegistered {
		internal.WriteResponse(
			w,
			dto.StatusResponse{
				User:    user.Name,
				Message: "User is not registered.",
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
				Message: "Error while changing user balance: " + err.Error(),
			},
			http.StatusInternalServerError)
		return
	}
	internal.WriteResponse(
		w,
		models.User{
			Name:    user.Name,
			Balance: user.Balance,
		},
		http.StatusOK,
	)
}
