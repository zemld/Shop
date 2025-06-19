package mq

import "time"

const (
	URL = "nats://nats:4222"
)

const (
	NewOrders     = "orders.new"
	Storage       = "confirm.new"
	Payment       = "pay.new"
	PaymentDone   = "pay.done"
	PaymentCancel = "pay.cancel"
	OrderHandled  = "orders.handled"
)

const (
	DefaultTimeout = 5 * time.Second
)
