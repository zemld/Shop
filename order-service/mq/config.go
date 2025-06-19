package mq

import "time"

const (
	URL = "nats://localhost:4222"
)

const (
	NewOrders     = "orders.new"
	Storage       = "confirm.new"
	Payment       = "pay.new"
	PaymentDone   = "pay.done"
	PaymentCancel = "pay.cancel"
	Shipping      = "orders.shipping"
	CancelOrder   = "orders.cancel"
)

const (
	DefaultTimeout = 5 * time.Second
)
