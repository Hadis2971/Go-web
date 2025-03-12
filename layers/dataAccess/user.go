package dataAccess

import (
	"database/sql"
	"errors"
	"log"

	"github.com/Hadis2971/go_web/models"
)

type UpdateUserRequest struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

type IUserDataAccess interface {
	CreateUser(user models.User)
	DeleteUser(id int) error
	UpdateUser(updateUserRequest UpdateUserRequest) error
	GetUserByUsernameOrEmail(user models.User) (*models.User, error)
	GetUserById(id int) (*models.User, error)
}

type UserDataAccess struct {
	dbConnection *sql.DB
}

func NewUserDataAccess(dbConnection *sql.DB) *UserDataAccess {
	return &UserDataAccess{dbConnection: dbConnection}
}

func (da UserDataAccess) CreateUser(user models.User) {
	query := "INSERT INTO User (username, email, password) VALUES (?, ?, ?)"

	_, err := da.dbConnection.Exec(query, user.Username, user.Email, user.Password)

	if err != nil {
		log.Fatal(err)
	} // fatal, never force a shutdown, handle the error
}

func (da UserDataAccess) DeleteUser(id int) error {
	query := "DELETE FROM User WHERE id = ?"

	_, err := da.dbConnection.Exec(query, id)

	if err != nil {
		return err
	} // provide context to your err otherwise won't be able to figure out what broke and where. fmt.Errorf("couldn't delete user: %v", err)

	return nil
}

func (da UserDataAccess) UpdateUser(updateUserRequest UpdateUserRequest) error {
	query := "UPDATE User SET username = ?, email = ? WHERE id = ?"

	_, err := da.dbConnection.Query(query, updateUserRequest.Username, updateUserRequest.Email, updateUserRequest.ID)

	if err != nil {
		return err
	}

	return nil
}

func (da UserDataAccess) GetUserByUsernameOrEmail(user models.User) (*models.User, error) {
	query := "SELECT * FROM User WHERE username = ? OR email = ?"
	var foundUser models.User
	hasResults := false

	rows, err := da.dbConnection.Query(query, user.Username, user.Email) // Should use QueryRow, you don't want more than 1 user here.
	defer rows.Close()                                                   // You need to call this after handling the error, other rows will be nil and cause a panic

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		hasResults = true
		rows.Scan(&foundUser.ID, &foundUser.Username, &foundUser.Email) // Handle error
	}

	if !hasResults {
		return nil, errors.New("User Not Found") // go prefers lowercase "user not found"
	}

	return &foundUser, nil
}

func (da UserDataAccess) GetUserById(id int) (*models.User, error) {
	query := "SELECT * FROM User WHETE id = ?" // Typo WHETE -> WHERE
	var foundUser models.User
	hasResults := false

	rows, err := da.dbConnection.Query(query, id)
	defer rows.Close()

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		hasResults = true
		rows.Scan(&foundUser.ID, &foundUser.Username, &foundUser.Email, &foundUser.Password)
	}

	if !hasResults {
		return nil, errors.New("No User Found")
	}

	return &foundUser, nil
}
