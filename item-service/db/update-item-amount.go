package db

import (
	"database/sql"
)

func CreateDBConnectionAndUpdateItemAmount(dbConnection string, name string, diff int, createTableQuery ...string) error {
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

	err = updateItemAmount(db, name, diff)
	if err != nil {
		return err
	}
	return nil
}

func updateItemAmount(db *sql.DB, name string, diff int) error {
	if db == nil {
		return sql.ErrNoRows
	}

	ctx, cancel := getContext()
	defer cancel()

	_, err := db.ExecContext(ctx, updateItemAmountQuery, name, diff)
	if err != nil {
		return err
	}

	return nil
}
