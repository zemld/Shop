package models

type Order struct {
	Items []ItemInOrder `json:"items"`
	User  string        `json:"user"`
}

type ItemInOrder struct {
	Name   string `json:"name"`
	Amount int    `json:"amount"`
}
