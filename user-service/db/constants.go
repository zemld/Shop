package db

const (
	UsersDB               = "postgres://user:password@user-db:5432/userdb"
	createUsersTableQuery = `CREATE TABLE IF NOT EXISTS users (
		username TEXT PRIMARY KEY,
		balance FLOAT NOT NULL DEFAULT 0.0 CHECK (balance >= 0.0)
		);`
	CheckUserRegisteredQuery = "SELECT 1 FROM users WHERE username = $1 LIMIT 1;"
	RegisterUserQuery        = "INSERT INTO users (username, balance) VALUES ($1, $2) ON CONFLICT (username) DO UPDATE SET balance = EXCLUDED.balance RETURNING username, balance;"
)
