package db

import (
	"database/sql"
	"log"

	"github.com/jackc/pgx/v5"
)

func CreateDBConnectionAndCheckAdmin(dbConnection string, code string, createTableQuery ...string) (string, error) {
	createTableUsersQuery := createAdminTableQuery
	if len(createTableQuery) > 0 {
		log.Println("Using custom create table query for users")
		createTableUsersQuery = createTableQuery[0]
	}

	db, err := ConnectDB(dbConnection)
	if err != nil {
		return "", err
	}
	defer db.Close()

	if err := CreateTable(db, createTableUsersQuery); err != nil {
		return "", err
	}

	adminName, err := checkIsAdmin(db, code)
	if err != nil {
		return "", err
	}

	return adminName, nil
}

func checkIsAdmin(db *sql.DB, code string) (string, error) {
	if db == nil {
		return "", pgx.ErrNoRows
	}

	var adminName string
	ctx, cancel := getContext()
	defer cancel()
	err := db.QueryRowContext(ctx, checkIsAdminQuery, code).Scan(&adminName)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", nil
		}
		return "", err
	}

	return adminName, nil
}
