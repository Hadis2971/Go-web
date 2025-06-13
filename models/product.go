package models

type ProductId int

type Product struct {
	ID ProductId `json:"id"`
	Name string `json:"name"`
	Price float64 `json:"price"`
	Description string `json:"description"`
	Stock int `json:"stock"`
	CreatedOn string `json:"created_on"`
	UpdatedOn string `json:"updated_on"`
}