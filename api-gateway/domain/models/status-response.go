package models

type OrderStatusResponse struct {
	OrderID int    `json:"order_id"`
	Order   Order  `json:"order"`
	Message string `json:"message"`
}
