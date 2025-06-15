package models

type ItemResponse struct {
	Item    Item   `json:"item"`
	Message string `json:"message"`
}
