package models

import "time"

type ProductOrder struct {
	ID int `json:"id"`
	UserId UserId `json:"user_id"`
	ProductId ProductId `json:"product_id"`
	Quantity int `json:"quantity"`
	CreatedOn time.Time `json:"created_on"`
	UpdatedOn time.Time `json:"updated_on"`
}