package internal

import (
	"net/http"
	"strconv"

	"github.com/zemld/Shop/item-service/domain/models"
)

func ValidateItemFromRequest(r *http.Request) (models.Item, error) {
	name, err := ValidateItemNameFromRequest(r)
	if err != nil {
		return models.Item{}, &models.ValidationError{Message: "Invalid item name"}
	}
	price, err := ValidateItemPriceFromRequest(r)
	if err != nil {
		return models.Item{}, &models.ValidationError{Message: "Invalid item price"}
	}
	amount, err := ValidateItemAmountFromRequest(r)
	if err != nil {
		return models.Item{}, &models.ValidationError{Message: "Invalid item amount"}
	}

	return models.Item{
		Name:   name,
		Price:  price,
		Amount: amount,
	}, nil
}

func ValidateItemNameFromRequest(r *http.Request) (string, error) {
	name := r.URL.Query().Get("name")
	if name == "" {
		return "", &models.ValidationError{Message: "Item name is required"}
	}
	return name, nil
}

func ValidateItemPriceFromRequest(r *http.Request) (float64, error) {
	priceStr := r.URL.Query().Get("price")
	if priceStr == "" {
		return 0, &models.ValidationError{Message: "Item price is required"}
	}

	price, err := strconv.ParseFloat(priceStr, 64)
	if err != nil || price < 0 {
		return 0, &models.ValidationError{Message: "Invalid item price"}
	}

	return price, nil
}

func ValidateItemAmountFromRequest(r *http.Request) (int, error) {
	amountStr := r.URL.Query().Get("amount")
	if amountStr == "" {
		return 0, &models.ValidationError{Message: "Item amount is required"}
	}

	amount, err := strconv.Atoi(amountStr)
	if err != nil || amount < 0 {
		return 0, &models.ValidationError{Message: "Invalid item amount"}
	}

	return amount, nil
}
