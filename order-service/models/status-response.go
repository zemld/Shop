package models

type StatusResponse struct {
	OrderID int    `json:"order_id"`
	Order   Order  `json:"order"`
	Message string `json:"message"`
}
