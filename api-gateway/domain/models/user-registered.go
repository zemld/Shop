package models

type UserRegistered struct {
	User         string `json:"user"`
	IsRegistered bool   `json:"is_registered"`
}
