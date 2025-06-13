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

type FoundUserReponse struct {
	ID models.UserId `json:"id"`
	Username string `json:"username"`
	Email string `json:"email"`
	CreatedOn string `json:"created_on"`
	UpdatedOn string `json:"updated_on"`
	Password string
}

type IUserDataAccess interface {
	CreateUser(user models.User) error
	DeleteUser(id int) error
	UpdateUser(updateUserRequest UpdateUserRequest) error
	GetUserByUsernameOrEmail(user models.User) (*FoundUserReponse, error)
	GetUserById(id int) (*FoundUserReponse, error)
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

func (da UserDataAccess) GetUserByUsernameOrEmail(user models.User) (*FoundUserReponse, error) {
	query := "SELECT * FROM User WHERE username = ? OR email = ?"
	
	var foundUserReponse FoundUserReponse

	if (user.Username == "" && user.Email == "") {
		return nil, ErrorMissingUsernameOrEmail
	}

	row := da.dbConnection.QueryRow(query, user.Username, user.Email)

	err := row.Scan(&foundUserReponse.ID, &foundUserReponse.Username, &foundUserReponse.Email, &foundUserReponse.Password, &foundUserReponse.CreatedOn, &foundUserReponse.UpdatedOn)

	if err != nil {
		return nil, err
	}

	return &foundUserReponse, nil
}

func (da UserDataAccess) GetUserById(id int) (*FoundUserReponse, error) {
	query := "SELECT * FROM User WHERE id = ?"
	var foundUserReponse FoundUserReponse

	if (id == 0) {
		return nil, ErrorMissingID
	}

	row := da.dbConnection.QueryRow(query, id)

	err := row.Scan(&foundUserReponse.ID, &foundUserReponse.Username, &foundUserReponse.Email, &foundUserReponse.Password, &foundUserReponse.CreatedOn, &foundUserReponse.UpdatedOn)


	
	if err != nil {
		return nil, ErrorUserNotFound
	}

	return &foundUserReponse, nil
}
