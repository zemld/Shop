package mq

import (
	"encoding/json"
	"log"

	"github.com/nats-io/nats.go"
	"github.com/zemld/Shop/item-service/domain/models"
)

func HandleNewOrder() {
	nc, err := nats.Connect(URL)
	if err != nil {
		log.Fatalf("Error connecting mq: %v", err)
	}
	_, err = nc.Subscribe(NewOrders, func(msg *nats.Msg) {
		log.Printf("Recieved message: %s", string(msg.Data))
		var orderMsg models.OrderMsg
		if err := json.Unmarshal(msg.Data, &orderMsg); err != nil {
			log.Printf("Error unmarshalling message: %v", err)
			return
		}
		checkedOrder, err := updateItemsInStock(orderMsg, true)
		if err != nil {
			log.Printf("Error checking items in stock: %v", err)
			return
		}
		encodedOrder, _ := json.Marshal(checkedOrder)
		if err := nc.Publish(Storage, encodedOrder); err != nil {
			log.Printf("Error publishing message: %v", err)
			return
		}
		log.Printf("Order processed: %v", checkedOrder)
	})
	if err != nil {
		log.Fatalf("Error mq subscription: %v", err)
	}
}
