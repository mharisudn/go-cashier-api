package models

type Category struct {
	ID          int           `json:"id"`
	Name        string        `json:"name"`
	Description string        `json:"description"`
	Products    []ProductList `json:"products"`
}

type ProductList struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
	Stock int    `json:"stock"`
}
