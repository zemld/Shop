package models

// TODO: по-хорошему здесь нужен какой-то статус операции задавать.
// Мб отдельный респонс для обновления количества товара.
type ItemResponse struct {
	Item    Item   `json:"item"`
	Message string `json:"message"`
}
