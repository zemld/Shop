package db

import (
	"database/sql"
)

const ()

func ConnectDB() (*sql.DB, error) {
	db, err := sql.Open("pgx", usersDB)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
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
