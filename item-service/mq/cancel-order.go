package mq

import (
	"encoding/json"
	"log"

	"github.com/nats-io/nats.go"
	"github.com/zemld/Shop/item-service/domain/models"
)

func HandleCanceledOrder() {
	nc, err := nats.Connect(URL)
	if err != nil {
		log.Fatalf("Error connecting mq: %v", err)
	}
	_, err = nc.Subscribe(PaymentCancel, func(msg *nats.Msg) {
		log.Printf("Recieved message: %s", string(msg.Data))
		var order models.OrderStatusResponse
		if err := json.Unmarshal(msg.Data, &order); err != nil {
			log.Printf("Error unmarshalling message: %v", err)
			return
		}
		orderMessage := models.OrderMsg{
			Id:    order.OrderID,
			User:  order.Order.User,
			Items: order.Order.Items,
		}
		log.Printf("Unmarshalled order message: %v", orderMessage)
		restoredOrder, err := updateItemsInStock(orderMessage, false)
		log.Printf("Restoring items in stock for order: %v", restoredOrder)
		if err != nil {
			log.Printf("Error updating items in stock: %v", err)
			return
		}
		order.Message = "Order canceled"
		encodedResponse, _ := json.Marshal(order)
		log.Printf("Response: %v", order)
		if err := nc.Publish(OrderHandled, encodedResponse); err != nil {
			log.Printf("Error publishing message: %v", err)
			return
		}
		log.Printf("Order processed: %v", restoredOrder)
	})
	if err != nil {
		log.Fatalf("Error mq subscription: %v", err)
	}
}
