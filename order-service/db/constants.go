package db

const (
	OutboxDB               = "postgres://order_outbox:password@order-outboxdb:5432/order-outbox-db"
	createOutboxTableQuery = `CREATE TABLE IF NOT EXISTS outbox (
		id SERIAL NOT NULL,
		created_at TIMESTAMP NOT NULL DEFAULT NOW() PRIMARY KEY,
		order JSONB NOT NULL
		);`
	createOrderQuery         = "INSERT INTO outbox (order) VALUES ($1) RETURNING id;"
	handleOrderQuery         = "DELETE FROM outbox WHERE id = $1;"
	getUnhandeledOrdersQuery = `SELECT id, order FROM outbox WHERE created_at < NOW() ORDER BY created_at;`
)
