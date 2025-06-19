package models

import "time"

type ProductOrderId int
type OrderId int

type ProductOrder struct {
	ID ProductOrderId `json:"id"`
	UserId UserId `json:"user_id"`
	ProductId ProductId `json:"product_id"`
	Quantity int `json:"quantity"`
	OrderId OrderId `json:"order_id"`
	CreatedOn time.Time `json:"created_on"`
	UpdatedOn time.Time `json:"updated_on"`
}

type CreateProductOrderReqPayload struct {
	UserId int `json:"userId"`
	ProductId int `json:"productId"`
	Quantity int `json:"quantity"`
}