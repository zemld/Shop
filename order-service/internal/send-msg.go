package internal

import (
	"encoding/json"

	"github.com/nats-io/nats.go"
	"github.com/zemld/Shop/order-service/db"
	"github.com/zemld/Shop/order-service/models"
	"github.com/zemld/Shop/order-service/mq"
)

func SendMsg(order models.Order) (*nats.Conn, error) {
	database, tx, err := db.BeginTransaction(db.OutboxDB)
	if err != nil {
		return nil, err
	}
	defer db.CloseDB(database)
	encodedOrder, _ := json.Marshal(order)
	id, err := db.StoreNewOrder(tx, encodedOrder)
	if err != nil {
		db.RollbackTransaction(tx)
		return nil, err
	}
	orderMsg := models.OrderMsg{
		Id:    id,
		Items: order.Items,
		User:  order.User,
	}
	nc, err := nats.Connect(mq.URL)
	if err != nil {
		db.RollbackTransaction(tx)
		return nil, err
	}
	encodedOrderMsg, _ := json.Marshal(orderMsg)
	if err = nc.Publish(mq.NewOrders, encodedOrderMsg); err != nil {
		db.RollbackTransaction(tx)
		return nil, err
	}
	tx.Commit()
	return nc, nil
}
