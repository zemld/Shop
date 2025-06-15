package db

import (
	"database/sql"

	"github.com/jackc/pgx/v5"
	"github.com/zemld/Shop/item-service/domain/models"
)

func CreateDBConnectionAndAddItem(dbConnection string, item models.Item, createTableQuery ...string) error {
	creationQuery := createItemsTableQuery
	if len(createTableQuery) > 0 {
		creationQuery = createTableQuery[0]
	}

	db, err := ConnectDB(dbConnection)
	if err != nil {
		return err
	}
	defer db.Close()

	if err := CreateTable(db, creationQuery); err != nil {
		return err
	}

	err = addItem(db, item)
	if err != nil {
		return err
	}
	return nil
}

func addItem(db *sql.DB, item models.Item) error {
	if db == nil {
		return pgx.ErrNoRows
	}

	ctx, cancel := getContext()
	defer cancel()

	_, err := db.ExecContext(ctx, addItemQuery, item.Name, item.Price, item.Amount)
	if err != nil {
		return err
	}

	return nil
}
