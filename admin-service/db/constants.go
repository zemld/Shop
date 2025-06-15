package db

const (
	AdminDB               = "postgres://admin:password@admin-db:5432/admindb"
	createAdminTableQuery = `CREATE TABLE IF NOT EXISTS admins (
		username TEXT PRIMARY KEY,
		code TEXT NOT NULL
		);`
	checkIsAdminQuery  = "SELECT 1 FROM admins WHERE username = $1 LIMIT 1;"
	registerAdminQuery = "INSERT INTO admins (username, code) VALUES ($1, $2) ON CONFLICT (username) DO UPDATE SET code = EXCLUDED.code RETURNING username, code;"
)
