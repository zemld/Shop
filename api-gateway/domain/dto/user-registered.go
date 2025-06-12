package dto

type UserRegistered struct {
	UserID       int  `json:"user_id"`
	IsRegistered bool `json:"is_registered"`
}
