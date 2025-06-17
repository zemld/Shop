package db

import (
	"database/sql"
	"errors"
)

func StoreNewOrder(tx *sql.Tx, order []byte) (int64, error) {
	if tx == nil {
		return -1, errors.New("database connection is nil")
	}

	ctx, cancel := getContext()
	defer cancel()

	var id int64
	if err := tx.QueryRowContext(ctx, createOrderQuery, order).Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}
