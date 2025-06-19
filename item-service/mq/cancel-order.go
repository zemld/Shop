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
		var orderMsg models.OrderMsg
		if err := json.Unmarshal(msg.Data, &orderMsg); err != nil {
			log.Printf("Error unmarshalling message: %v", err)
			return
		}
		restoredOrder, err := updateItemsInStock(orderMsg, false)
		if err != nil {
			log.Printf("Error updating items in stock: %v", err)
			return
		}
		encodedOrder, _ := json.Marshal(restoredOrder)
		if err := nc.Publish(OrderHandled, encodedOrder); err != nil {
			log.Printf("Error publishing message: %v", err)
			return
		}
		log.Printf("Order processed: %v", restoredOrder)
	})
	if err != nil {
		log.Fatalf("Error mq subscription: %v", err)
	}
}
