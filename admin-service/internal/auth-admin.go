package internal

import "github.com/zemld/Shop/admin-service/db"

func AuthAdmin(secretCode string) bool {
	adminName, err := db.CreateDBConnectionAndCheckAdmin(db.AdminDB, secretCode)
	if err != nil {
		return false
	}
	return adminName != ""
}
