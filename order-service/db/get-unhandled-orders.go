package db

import (
	"database/sql"
	"errors"
)

func GetUnhandledOrders(db *sql.DB) (map[int64][]byte, error) {
	if db == nil {
		return nil, errors.New("database connection is nil")
	}

	ctx, cancel := getContext()
	defer cancel()

	rows, err := db.QueryContext(ctx, getUnhandeledOrdersQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	unhandledOrders := make(map[int64][]byte)
	for rows.Next() {
		var id int64
		var order []byte
		if err := rows.Scan(&id, &order); err != nil {
			return nil, err
		}
		unhandledOrders[id] = order
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return unhandledOrders, nil
}
