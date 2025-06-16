package db

import "database/sql"

func BeginTransaction(dbConnection string) (*sql.DB, *sql.Tx, error) {
	db, err := ConnectDB(dbConnection)
	if err != nil {
		return nil, nil, err
	}

	ctx, cancel := getContext()
	defer cancel()
	tx, err := db.BeginTx(ctx, nil)
	return db, tx, err
}

func CommitTransaction(tx *sql.Tx) error {
	if tx == nil {
		return sql.ErrTxDone
	}
	return tx.Commit()
}

func RollbackTransaction(tx *sql.Tx) error {
	if tx == nil {
		return sql.ErrTxDone
	}
	return tx.Rollback()
}
