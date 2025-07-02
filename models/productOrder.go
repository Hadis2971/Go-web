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

type ProductOrderWithProduct struct {
	ProductOrderId int `json:"product_order_id"`
	UserId int `json:"user_id"`
	OrderId string `json:"order_id"`
	ProductId ProductId `json:"product_id"`
	Quantity int `json:"quantity"`
	Product Product `json:"product"`
}

type FullOrder struct {
	Id int `json:"id"`
	Username string `json:"username"`
	Email string `json:"email"`
	ProductOrderId *int `json:"product_order_id"`
	UserId *int `json:"user_id"`
	OrderId *string `json:"order_id"`
	Quantity *int `json:"quantity"`
	ProductId *ProductId `json:"product_id"`
	ProductName *string `json:"product_name"`
	ProductDescription *string `json:"product_description"`
	ProductPrice *float64 `json:"product_price"`
	ProductStock *int `json:"product_stock"`
	ProductCreatedOn *time.Time `json:"product_created_on"`
	ProductUpdatedOn *time.Time `json:"product_updated_on"`
}