package handlers

import (
	"net/http"

	"github.com/zemld/Shop/user-service/db"
	"github.com/zemld/Shop/user-service/domain/dto"
	"github.com/zemld/Shop/user-service/domain/models"
	"github.com/zemld/Shop/user-service/internal"
)

// @description Returns user's balance.
// @tags Users
// @param name query string true "User whose balance you want to get"
// @produce json
// @success 200 {object} models.User
// @failure 400 {object} dto.StatusResponse
// @failure 500 {object} dto.StatusResponse
// @router /v1/users/balance [get]
func GetUserBalanceHandler(w http.ResponseWriter, r *http.Request) {
	user := r.URL.Query().Get("name")
	if user == "" {
		internal.WriteResponse(
			w,
			dto.StatusResponse{
				User:    user,
				Message: "You must provide a valid user name.",
			},
			http.StatusBadRequest)
		return
	}
	balance, err := db.CreateDBConnectionAndGetUserBalance(db.UsersDB, user)
	if err != nil {
		internal.WriteResponse(
			w,
			dto.StatusResponse{
				User:    user,
				Message: "Error while getting user balance: " + err.Error(),
			},
			http.StatusInternalServerError)
		return
	}
	internal.WriteResponse(
		w,
		models.User{
			Name:    user,
			Balance: balance,
		},
		http.StatusOK)
}
