package db

import (
	"database/sql"

	"github.com/jackc/pgx/v5"
)

func CreateDBConnectionAndGetUserBalance(dbConnection string, username string, createTableQuery ...string) (float64, error) {
	createTableUsersQuery := createUsersTableQuery
	if len(createTableQuery) > 0 {
		createTableUsersQuery = createTableQuery[0]
	}

	db, err := ConnectDB(dbConnection)
	if err != nil {
		return 0.0, err
	}
	defer db.Close()

	if err := CreateTable(db, createTableUsersQuery); err != nil {
		return 0.0, err
	}

	balance, err := getUserBalance(db, username)
	if err != nil {
		return 0.0, err
	}

	return balance, nil
}

func getUserBalance(db *sql.DB, username string) (float64, error) {
	if db == nil {
		return 0.0, pgx.ErrNoRows
	}

	ctx, cancel := getContext()
	defer cancel()

	var balance float64
	err := db.QueryRowContext(ctx, GetUserBalanceQuery, username).Scan(&balance)
	if err != nil {
		return 0.0, err
	}

	return balance, nil
}
