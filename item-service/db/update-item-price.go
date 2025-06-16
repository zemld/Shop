package db

import (
	"database/sql"

	"github.com/zemld/Shop/item-service/domain/models"
)

func UpdateItemPrice(dbConnection string, item models.Item, newPrice float64, createTableQuery ...string) (models.Item, error) {
	creationQuery := createItemsTableQuery
	if len(createTableQuery) > 0 {
		creationQuery = createTableQuery[0]
	}

	db, err := ConnectDB(dbConnection)
	if err != nil {
		return models.Item{}, err
	}
	defer db.Close()

	if err := CreateTable(db, creationQuery); err != nil {
		return models.Item{}, err
	}

	updatedItem, err := updateItemPrice(db, item, newPrice)
	if err != nil {
		return models.Item{}, err
	}
	return updatedItem, nil
}

func updateItemPrice(db *sql.DB, item models.Item, newPrice float64) (models.Item, error) {
	if db == nil {
		return models.Item{}, sql.ErrNoRows
	}

	ctx, cancel := getContext()
	defer cancel()

	var updatedItem models.Item
	err := db.QueryRowContext(ctx, updateItemPriceQuery, item.Name, newPrice).Scan(&updatedItem.Name, &updatedItem.Price, &updatedItem.Amount)
	if err != nil {
		return models.Item{}, err
	}

	return updatedItem, nil
}
