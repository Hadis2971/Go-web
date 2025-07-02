package models

import (
	"time"
)

type UserId int

type User struct {
	ID       UserId    `json:"id"` 
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	CreatedOn time.Time `json:"created_on"`
	UpdatedOn time.Time `json:"updated_on"`
}

type BasicUser struct {
	ID int `json:"id"`
	Username string `json:"username"`
	Email string `json:"email"`
}

type UserWithOrders struct {
	ID int `json:"id"`
	Username string `json:"username"`
	Email string `json:"email"`
	Orders []ProductOrderWithProduct `json:"orders"`
}