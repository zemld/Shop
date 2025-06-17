package db

import (
	"database/sql"
	"errors"
)

func ConfirmOrder(tx *sql.Tx, id int64) error {
	if tx == nil {
		return errors.New("database connection is nil")
	}

	ctx, cancel := getContext()
	defer cancel()

	if _, err := tx.ExecContext(ctx, handleOrderQuery, id); err != nil {
		return err
	}

	return nil
}
