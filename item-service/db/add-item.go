package db

import (
	"database/sql"

	"github.com/jackc/pgx/v5"
	"github.com/zemld/Shop/item-service/domain/models"
)

func AddItem(dbConnection string, item models.Item, createTableQuery ...string) (models.Item, error) {
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

	updatedItem, err := addItem(db, item)
	if err != nil {
		return models.Item{}, err
	}
	return updatedItem, nil
}

func addItem(db *sql.DB, item models.Item) (models.Item, error) {
	if db == nil {
		return models.Item{}, pgx.ErrNoRows
	}

	ctx, cancel := getContext()
	defer cancel()

	var updatedItem models.Item
	err := db.QueryRowContext(ctx, addItemQuery, item.Name, item.Price, item.Amount).Scan(&updatedItem.Name, &updatedItem.Price, &updatedItem.Amount)
	if err != nil {
		return models.Item{}, err
	}

	return updatedItem, nil
}
