package db

const (
	ItemsDB               = "postgres://item:password@item-db:5432/itemdb"
	createItemsTableQuery = `CREATE TABLE IF NOT EXISTS items (
		name TEXT PRIMARY KEY,
		price FLOAT NOT NULL CHECK (price >= 0.0),
		amount INT NOT NULL CHECK (amount >= 0)
		);`
	addItemQuery          = "INSERT INTO items (name, price, amount) VALUES ($1, $2, $3) ON CONFLICT (name) DO UPDATE SET price = EXCLUDED.price RETURNING username, balance;"
	updateItemPriceQuery  = "UPDATE items SET price = $2 WHERE name = $1;"
	updateItemAmountQuery = "UPDATE items SET amount = MAX(amount + $2, 0) WHERE name = $1;"
	getItemQuery          = "SELECT name, price, amount FROM items WHERE name = $1;"
	deleteItemQuery       = "DELETE FROM items WHERE name = $1;"
)
