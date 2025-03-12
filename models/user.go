package models

type User struct {
	ID       int    `json:"id"` // ID is an int? It'd be better to create your own custom type. So you don't end up with the ID of something else
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"-"`
}
