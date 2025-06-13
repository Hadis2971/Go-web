package models

import "time"

type UserId int

type User struct {
	ID       UserId    `json:"id"` 
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	CreatedOn time.Time `json:"created_on"`
	UpdatedOn time.Time `json:"updated_on"`
}

type CreateProductReq struct {
	Name string `json:"name"`
	Description string `json:"description"`
	Price int `json:"price"`
	Stock int `json:"stock"`
}