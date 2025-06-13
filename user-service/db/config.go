package db

import (
	"context"
	"database/sql"
	"time"
)

func getContext() (context.Context, context.CancelFunc) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	return ctx, cancel
}

func ConnectDB(dbConnection string) (*sql.DB, error) {
	db, err := sql.Open("pgx", dbConnection)
	if err != nil {
		return nil, err
	}

	ctx, cancel := getContext()
	defer cancel()
	if err = db.PingContext(ctx); err != nil {
		return nil, err
	}

	return db, nil
}

func CreateTable(db *sql.DB, createTableQuery string) error {
	if db == nil {
		return nil
	}
	ctx, cancel := getContext()
	defer cancel()
	if _, err := db.ExecContext(ctx, createTableQuery); err != nil {
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
