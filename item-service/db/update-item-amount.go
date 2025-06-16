package db

import (
	"database/sql"

	"github.com/zemld/Shop/item-service/domain/models"
)

func ConnectToDBAndUpdateItemAmount(dbConnection string, name string, diff int, createTableQuery ...string) (models.Item, error) {
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

	updatedItem, err := updateItemAmount(db, name, diff)
	if err != nil {
		return models.Item{}, err
	}
	return updatedItem, nil
}

func UpdateItemAmount(db *sql.DB, name string, diff int, createTableQuery ...string) (models.Item, error) {
	creationQuery := createItemsTableQuery
	if len(createTableQuery) > 0 {
		creationQuery = createTableQuery[0]
	}

	if db == nil {
		return models.Item{}, sql.ErrNoRows
	}

	if err := CreateTable(db, creationQuery); err != nil {
		return models.Item{}, err
	}

	updatedItem, err := updateItemAmount(db, name, diff)
	if err != nil {
		return models.Item{}, err
	}
	return updatedItem, nil
}

func updateItemAmount(db *sql.DB, name string, diff int) (models.Item, error) {
	if db == nil {
		return models.Item{}, sql.ErrNoRows
	}

	ctx, cancel := getContext()
	defer cancel()

	var updatedItem models.Item
	err := db.QueryRowContext(ctx, updateItemAmountQuery, name, diff).Scan(&updatedItem.Name, &updatedItem.Price, &updatedItem.Amount)
	if err != nil {
		return models.Item{}, err
	}

	return updatedItem, nil
}
