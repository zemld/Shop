package db

import (
	"database/sql"

	"github.com/jackc/pgx/v5"
)

func CreateDBConnectionAndRegisterAdmin(dbConnection string, username string, code string, createTableQuery ...string) error {
	createTableUsersQuery := createAdminTableQuery
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

	err = registerAdmin(db, username, code)
	if err != nil {
		return err
	}
	return nil
}

func registerAdmin(db *sql.DB, username string, code string) error {
	if db == nil {
		return pgx.ErrNoRows
	}

	ctx, cancel := getContext()
	defer cancel()

	_, err := db.ExecContext(ctx, registerAdminQuery, username, code)
	if err != nil {
		return err
	}

	return nil
}
