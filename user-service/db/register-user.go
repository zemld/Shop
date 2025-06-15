package db

import (
	"database/sql"

	"github.com/jackc/pgx/v5"
)

func CreateDBConnectionAndRegisterUser(dbConnection string, username string, balance float64, createTableQuery ...string) error {
	createTableUsersQuery := createUsersTableQuery
	if len(createTableQuery) > 0 {
		createTableUsersQuery = createTableQuery[0]
	}

	db, err := ConnectDB(dbConnection)
	if err != nil {
		return err
	}
	defer db.Close()

	if err := CreateTable(db, createTableUsersQuery); err != nil {
		return err
	}

	err = registerUser(db, username, balance)
	if err != nil {
		return err
	}
	return nil
}

func registerUser(db *sql.DB, username string, balance float64) error {
	if db == nil {
		return pgx.ErrNoRows
	}

	ctx, cancel := getContext()
	defer cancel()

	_, err := db.ExecContext(ctx, RegisterUserQuery, username, balance)
	if err != nil {
		return err
	}

	return nil
}
