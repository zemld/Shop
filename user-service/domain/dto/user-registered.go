package dto

type UserRegistered struct {
	User         string `json:"user_id"`
	IsRegistered bool   `json:"is_registered"`
}
