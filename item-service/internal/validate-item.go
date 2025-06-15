package internal

import (
	"net/http"
	"strconv"

	"github.com/zemld/Shop/item-service/domain/models"
)

func ValidateItemFromRequest(r *http.Request) (models.Item, error) {
	name, amount, err := ValidateItemNameAndAmountFromRequest(r)
	if err != nil {
		return models.Item{}, err
	}

	priceStr := r.URL.Query().Get("price")

	if priceStr == "" {
		return models.Item{}, &models.ValidationError{Message: "Invalid item data"}
	}

	price, err := strconv.ParseFloat(priceStr, 64)
	if err != nil || price < 0 {
		return models.Item{}, &models.ValidationError{Message: "Invalid price"}
	}

	return models.Item{
		Name:   name,
		Price:  price,
		Amount: amount,
	}, nil
}

func ValidateItemNameAndAmountFromRequest(r *http.Request) (string, int, error) {
	name := r.URL.Query().Get("name")
	amountStr := r.URL.Query().Get("amount")

	if name == "" || amountStr == "" {
		return "", 0, &models.ValidationError{Message: "Invalid item data"}
	}

	amount, err := strconv.Atoi(amountStr)
	if err != nil || amount < 0 {
		return "", 0, &models.ValidationError{Message: "Invalid amount"}
	}

	return name, amount, nil
}
