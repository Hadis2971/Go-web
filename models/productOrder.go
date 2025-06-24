package models

import (
	"time"
)

type ProductOrderId int
type OrderId string

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
	UserId int `json:"user_id"`
	ProductId int `json:"product_id"`
	Quantity int `json:"quantity"`
}

type ProductAndUser struct {
	UserId UserId `json:"user_id"`
	ProductId ProductId `json:"product_id"`
	Username string `json:"username"`
	Quantity int `json:"quantity"`
	OrderCreated string `json:"order_created"`
	OrderUpdated string `json:"order_updated"`
}