package handlers

import (
	"net/http"
	"strconv"

	"github.com/zemld/Shop/user-service/db"
	"github.com/zemld/Shop/user-service/domain/dto"
	"github.com/zemld/Shop/user-service/domain/models"
	"github.com/zemld/Shop/user-service/internal"
)

// @description Registers a new user.
// @tag.name Users operations
// @param name query string true "User which you want to register"
// @param balance query float64 true "Balance of the user"
// @produce json
// @success 200 {object} models.User
// @failure 400 {object} dto.StatusResponse
// @failure 500 {object} dto.StatusResponse
// @router /v1/users/register [post]
func RegisterUserHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	balance, err := strconv.ParseFloat(r.URL.Query().Get("balance"), 64)
	if name == "" || err != nil || balance < 0 {
		internal.WriteResponse(
			w,
			dto.StatusResponse{
				User:    name,
				Message: "You must provide a valid name and balance for the user.",
			},
			http.StatusBadRequest)
		return
	}

	err = db.CreateDBConnectionAndRegisterUser(db.UsersDB, name, balance)
	if err != nil {
		internal.WriteResponse(
			w,
			dto.StatusResponse{
				User:    name,
				Message: "Error while registering user: " + err.Error(),
			},
			http.StatusInternalServerError)
		return
	}
	user := models.User{
		Name:    name,
		Balance: balance,
	}
	internal.WriteResponse(w, user, http.StatusOK)
}
