package main

import (
	"github.com/zemld/Shop/payment-service/internal"
)

func main() {
	go internal.PayForOrder()
	select {}
}
