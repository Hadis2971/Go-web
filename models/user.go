package models

type UserId int

type User struct {
	ID       UserId    `json:"id"` 
	// ID is an int? It'd be better to create your own custom type. So you don't end up with the ID of something else
	// Hadis => Can you go into a bit more details here? What is the difference now and how would I use the id of something else?
	
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
