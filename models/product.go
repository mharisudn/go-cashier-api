package models

type Product struct {
	ID         int             `json:"id"`
	Name       string          `json:"name"`
	Price      int             `json:"price"`
	Stock      int             `json:"stock"`
	CategoryID *int            `json:"category_id"`
	Category   *CategoryDetail `json:"category"`
}

type ProductUpdate struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Price      int    `json:"price"`
	Stock      int    `json:"stock"`
	CategoryID *int   `json:"category_id"`
}

type ProductCreate struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Price      int    `json:"price"`
	Stock      int    `json:"stock"`
	CategoryID *int   `json:"category_id"`
}

type CategoryDetail struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}
