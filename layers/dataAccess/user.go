package dataAccess

import (
	"database/sql"
	"errors"
	"log"

	"github.com/Hadis2971/go_web/models"
)

type UserDataAccess struct {
	dbConnection *sql.DB
}

type FoundUserReponse struct {
	ID int
	Username string
	Email string
}

func NewUserDataAccess (dbConnection *sql.DB) *UserDataAccess {
	return &UserDataAccess{dbConnection: dbConnection}
}

func (da UserDataAccess) CreateUser (user models.User) sql.Result {
	query := "INSERT INTO User (username, email, password) VALUES (?, ?, ?)";

	result, err := da.dbConnection.Exec(query, user.Username, user.Email, user.Password);

	if err != nil {
		log.Fatal(err)
	}

	return result
}

func (da UserDataAccess) DeleteUser (id int) error {
	query := "DELETE FROM User WHERE id = ?";

	_, err := da.dbConnection.Exec(query, id);

	if (err != nil) {
		return err
	}

	return nil;
}

func (da UserDataAccess) GetUserByUsernameOrEmail (user models.User) (*FoundUserReponse, error) {
	query := "SELECT id, username, email FROM User WHERE username = ? OR email = ?"
	var foundUser FoundUserReponse
	hasResults := false

	rows, err := da.dbConnection.Query(query, user.Username, user.Email)
	defer rows.Close()


	if err != nil {
		return nil, err
	}

	for rows.Next() {
		hasResults = true
		rows.Scan(&foundUser.ID, &foundUser.Username, &foundUser.Email)
	}

	if !hasResults {
		return nil, errors.New("User Not Found")
	}

	return &foundUser, nil
}