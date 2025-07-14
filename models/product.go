package models

import "time"

type ProductId int

type Product struct {
	ID ProductId `json:"id"`
	Name string `json:"name"`
	Price float64 `json:"price"`
	Description string `json:"description"`
	Stock int `json:"stock"`
	ProductCategory ProductCategoryId `json:"product_category"`
	CreatedOn time.Time `json:"created_on"`
	UpdatedOn time.Time `json:"updated_on"`
}

type ProductReqPayload struct {
	ID string 	`json:"id"`
	Name string `json:"name"`
	Description string `json:"description"`
	Price int `json:"price"`
	Stock int `json:"stock"`
	ProductCategory ProductCategoryId `json:"product_category"`
}

