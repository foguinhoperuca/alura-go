package models

type Pizza struct {
	ID    int     `json:"id"`
	Name  string  `json:"nombre"`
	Price float64 `json:"precio"`
}

type Person struct {
	ID     int    `json:"id"`
	Name   string `json:"nombre"`
	Age    int    `json:"edad"`
	Active bool   `json:"activo"`
}
