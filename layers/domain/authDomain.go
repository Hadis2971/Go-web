package domain

import (
	"errors"
	"log"

	"github.com/Hadis2971/go_web/layers/dataAccess"
	"github.com/Hadis2971/go_web/models"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword (password string) (string, error) {
	buffer, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost);

	if err != nil {
		return "", err
	}

	return string(buffer), nil
}

func CheckPassword (password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))

	return err == nil
}

type AuthDomain struct {
	userDataAccess dataAccess.IUserDataAccess
}

func NewAuthDomain (userDataAccess *dataAccess.UserDataAccess) *AuthDomain {
	return &AuthDomain{userDataAccess: userDataAccess}
}

func (ad *AuthDomain) RegisterUser (user models.User) (*models.User, error) {
	foundUser, _ := ad.userDataAccess.GetUserByUsernameOrEmail(user) 
	
	if foundUser != nil {
		

		return nil, errors.New("Username or Email Already Taken!!!");
	}

	hash, err := HashPassword(user.Password)

	if err != nil {
		log.Fatal(err)
	}

	user.Password = hash;

	ad.userDataAccess.CreateUser(user);

	return &user, nil
}

func (ad *AuthDomain) LoginUser (user models.User) (*dataAccess.FoundUserReponse, error) {
	foundUser, err := ad.userDataAccess.GetUserByUsernameOrEmail(user)

	if err != nil {
		return nil, errors.New("User Not Found!!!")
	}


	return foundUser, nil
}