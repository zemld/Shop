package models

type StatusResponse struct {
	Name    string `json:"name"`
	Message string `json:"message"`
}

type OrderStatusResponse struct {
	OrderID int64  `json:"order_id"`
	Order   Order  `json:"order"`
	Message string `json:"message"`
}
