package db

import (
	"database/sql"
)

const ()

func ConnectDB(dbConnection string) (*sql.DB, error) {
	db, err := sql.Open("pgx", dbConnection)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}

func CreateTable(db *sql.DB, createTableQuery string) error {
	if db == nil {
		return nil
	}
	if _, err := db.Exec(createTableQuery); err != nil {
		return err
	}
	return nil
}

func CloseDB(db *sql.DB) error {
	if db == nil {
		return nil
	}
	if err := db.Close(); err != nil {
		return err
	}
	return nil
}
