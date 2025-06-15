package internal

import (
	"net/http"
	"strconv"

	"github.com/zemld/Shop/item-service/domain/models"
)

func ValidateItemFromRequest(r *http.Request) (models.Item, error) {
	name := r.URL.Query().Get("name")
	priceStr := r.URL.Query().Get("price")
	amountStr := r.URL.Query().Get("amount")

	if name == "" || priceStr == "" || amountStr == "" {
		return models.Item{}, &models.ValidationError{Message: "Invalid item data"}
	}

	price, err := strconv.ParseFloat(priceStr, 64)
	if err != nil || price < 0 {
		return models.Item{}, &models.ValidationError{Message: "Invalid price"}
	}

	amount, err := strconv.Atoi(amountStr)
	if err != nil || amount < 0 {
		return models.Item{}, &models.ValidationError{Message: "Invalid amount"}
	}

	return models.Item{
		Name:   name,
		Price:  price,
		Amount: amount,
	}, nil
}
