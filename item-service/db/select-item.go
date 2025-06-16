package db

import (
	"database/sql"

	"github.com/jackc/pgx/v5"
	"github.com/zemld/Shop/item-service/domain/models"
)

func SelectItem(db *sql.DB, name string, createTableQuery ...string) (models.Item, error) {
	creationQuery := createItemsTableQuery
	if len(createTableQuery) > 0 {
		creationQuery = createTableQuery[0]
	}

	if err := CreateTable(db, creationQuery); err != nil {
		return models.Item{}, err
	}

	updatedItem, err := selectItem(db, name)
	if err != nil {
		return models.Item{}, err
	}
	return updatedItem, nil
}

func selectItem(db *sql.DB, name string) (models.Item, error) {
	if db == nil {
		return models.Item{}, pgx.ErrNoRows
	}

	ctx, cancel := getContext()
	defer cancel()

	var item models.Item
	err := db.QueryRowContext(ctx, getItemQuery, name).Scan(&item.Name, &item.Price, &item.Amount)
	if err != nil {
		return models.Item{}, err
	}

	return item, nil
}
