package db

import (
	"database/sql"
	"log"

	"github.com/jackc/pgx/v5"
)

func CreateDBConnectionAndCheckUserRegistered(dbConnection, username string, createTableQuery ...string) (bool, error) {
	createTableUsersQuery := createUsersTableQuery
	if len(createTableQuery) > 0 {
		log.Println("Using custom create table query for users")
		createTableUsersQuery = createTableQuery[0]
	}

	db, err := ConnectDB(dbConnection)
	if err != nil {
		return false, err
	}
	defer db.Close()

	if err := CreateTable(db, createTableUsersQuery); err != nil {
		return false, err
	}

	exists, err := checkUserRegistered(db, username)
	if err != nil {
		return false, err
	}

	return exists, nil
}

func checkUserRegistered(db *sql.DB, username string) (bool, error) {
	if db == nil {
		return false, pgx.ErrNoRows
	}

	var exists int
	ctx, cancel := getContext()
	defer cancel()
	err := db.QueryRowContext(ctx, CheckUserRegisteredQuery, username).Scan(&exists)
	doesExist := exists == 1
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, err
	}

	return doesExist, nil
}
