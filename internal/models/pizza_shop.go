package models

type Pizza struct {
	ID    int        `json:"id"`
	Name  string     `json:"nombre"`
	Price float64    `json:"precio"`
	Reviews []Review `json:"resenas"`
}

type Person struct {
	ID     int    `json:"id"`
	Name   string `json:"nombre"`
	Email  string `json:"correo_electronico"`
	Age    int    `json:"edad"`
	Active bool   `json:"activo"`
}
