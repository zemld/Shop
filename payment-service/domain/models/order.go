package models

type Order struct {
	Items []Item `json:"items"`
	User  string `json:"user"`
}

type OrderMsg struct {
	Items   []Item `json:"items"`
	User    string `json:"user"`
	Id      int64  `json:"id"`
	Message string `json:"message"`
}
