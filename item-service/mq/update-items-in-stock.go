package mq

import (
	"github.com/zemld/Shop/item-service/db"
	"github.com/zemld/Shop/item-service/domain/models"
)

func updateItemsInStock(order models.OrderMsg, isPurchase bool) (models.OrderMsg, error) {
	dbConn, err := db.ConnectDB(db.ItemsDB)
	if err != nil {
		return models.OrderMsg{}, err
	}
	defer dbConn.Close()
	var updatedOrder models.OrderMsg
	var updatedItems []models.ItemInOrder
	for _, item := range order.Items {
		var updatedItem models.Item
		var err error
		if isPurchase {
			updatedItem, err = db.UpdateItemAmount(dbConn, item.Name, -item.Amount)
		} else {
			updatedItem, err = db.UpdateItemAmount(dbConn, item.Name, item.Amount)
		}
		if err != nil {
			continue
		}
		updatedItems = append(updatedItems, models.ItemInOrder{
			Name:   updatedItem.Name,
			Amount: updatedItem.Amount,
		})
	}
	updatedOrder.Id = order.Id
	updatedOrder.User = order.User
	updatedOrder.Items = updatedItems
	return updatedOrder, nil
}
