package dataAccess

import (
	"database/sql"
	"errors"

	"github.com/Hadis2971/go_web/models"
)

type UpdateUserRequest struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

type IUserDataAccess interface {
	CreateUser(user models.User) error
	DeleteUser(id int) error
	UpdateUser(updateUserRequest UpdateUserRequest) error
	GetUserByUsernameOrEmail(user models.User) (*models.User, error)
	GetUserById(id int) (*models.User, error)
}

type UserDataAccess struct {
	dbConnection *sql.DB
}

var ( 
	InternalServerError = errors.New("Internal Server Error!!!")
	ErrorMissingID = errors.New("Missing ID Field!!!")
	ErrorMissingUsernameOrEmail = errors.New("Missing Username Or Email!!!")
	ErrorUserNotFound = errors.New("User Not Found!!!")
)

func NewUserDataAccess(dbConnection *sql.DB) *UserDataAccess {
	return &UserDataAccess{dbConnection: dbConnection}
}

func (da UserDataAccess) CreateUser(user models.User) error {
	query := "INSERT INTO User (username, email, password) VALUES (?, ?, ?)"

	_, err := da.dbConnection.Exec(query, user.Username, user.Email, user.Password)

	if err != nil {
		return InternalServerError
	}


	return nil
}

func (da UserDataAccess) DeleteUser(id int) error {
	query := "DELETE FROM User WHERE id = ?"

	if (id == 0) {
		return ErrorMissingID
	}

	_, err := da.dbConnection.Exec(query, id)

	if err != nil {
		return InternalServerError
	}

	return nil
}

func (da UserDataAccess) UpdateUser(updateUserRequest UpdateUserRequest) error {
	query := "UPDATE User SET username = ?, email = ? WHERE id = ?"

	if (updateUserRequest.ID == 0) {
		return ErrorMissingID
	}

	_, err := da.dbConnection.Query(query, updateUserRequest.Username, updateUserRequest.Email, updateUserRequest.ID)

	if err != nil {
		return InternalServerError
	}

	return nil
}

func (da UserDataAccess) GetUserByUsernameOrEmail(user models.User) (*models.User, error) {
	query := "SELECT * FROM User WHERE username = ? OR email = ?"
	var foundUser models.User

	if (user.Username == "" && user.Email == "") {
		return nil, ErrorMissingUsernameOrEmail
	}

	row := da.dbConnection.QueryRow(query, user.Username, user.Email)

	err := row.Scan(&foundUser.ID, &foundUser.Username, &foundUser.Email, &foundUser.Password)

	if err != nil {
		return nil, err
	}

	return &foundUser, nil
}

func (da UserDataAccess) GetUserById(id int) (*models.User, error) {
	query := "SELECT * FROM User WHERE id = ?"
	var foundUser models.User

	if (id == 0) {
		return nil, ErrorMissingID
	}

	row := da.dbConnection.QueryRow(query, id)

	
	if err := row.Scan(&foundUser.ID, &foundUser.Username, &foundUser.Email, &foundUser.Password); err != nil {
		return nil, ErrorUserNotFound
	}

	return &foundUser, nil
}
