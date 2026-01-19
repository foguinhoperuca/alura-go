package models

type Review struct {
	Rating  int    `json:"clasificacion"` // validation: [1-5]
	Comment string `json:"comentario"`
}
