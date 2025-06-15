package models

type IsAdminResponse struct {
	Name    string `json:"admin_name"`
	IsAdmin bool   `json:"is_admin"`
}
