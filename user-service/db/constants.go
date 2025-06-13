package db

const (
	usersDB               = "postgres://user:password@user-db:5432/userdb"
	createUsersTableQuery = `CREATE TABLE IF NOT EXISTS users (
		username TEXT PRIMARY KEY,
		balance FLOAT NOT NULL DEFAULT 0.0`
	checkUserRegisteredQuery = "SELECT 1 FROM users WHERE username = $1 LIMIT 1;"
)
