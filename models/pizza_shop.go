package models

type Pizza struct {
	ID    int     `json:"id"`
	Name  string  `json:"nombre"`
	Price float64 `json:"precio"`
}
