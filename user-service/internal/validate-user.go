package internal

import (
	"net/http"
	"strconv"

	"github.com/zemld/Shop/user-service/domain/models"
)

func ValidateUserFromRequest(r *http.Request) (models.User, error) {
	name := r.URL.Query().Get("name")
	balance, err := strconv.ParseFloat(r.URL.Query().Get("balance"), 64)
	if name == "" || err != nil || balance < 0 {
		return models.User{}, &models.ValidationError{Message: "Invalid user data"}
	}
	return models.User{
		Name:    name,
		Balance: balance,
	}, nil
}
