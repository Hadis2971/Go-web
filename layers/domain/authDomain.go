package domain

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"

	"github.com/Hadis2971/go_web/layers/dataAccess"
	"github.com/Hadis2971/go_web/models"
	"github.com/Hadis2971/go_web/util"
)

var (
	ErrorUsernameOrEmailAlreadyTaken = errors.New("Username or Email Already Taken!!!")
	ErrorRegisterUserMissingFields = errors.New("All Fileds Are Mandatory!!!")
	ErrorLoginUserInvalidCredentials = errors.New("Incorrect Credentials!!!")
)

func hashPassword(password string) (string, error) {
	buffer, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return "", err
	}

	return string(buffer), nil
}

func checkPassword(password string, hash string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)); err != nil {
		return false
	}

	return true
}

func generateJWTLoginToken (user *models.User) (string, error) {
	secret := util.GetEnvVariable("JWT_LOGIN_TOKEN_SECRET")

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, 
        jwt.MapClaims{ 
		"id": user.ID,
        "username": user.Username,
		"email": user.Email,
        "exp": time.Now().Add(time.Hour * 24).Unix(), 
        })

    tokenString, err := token.SignedString([]byte(secret))

    if err != nil {
    	return "", err
    }

 return tokenString, nil
}

type AuthDomain struct {
	userDataAccess dataAccess.IUserDataAccess
}

func NewAuthDomain(userDataAccess *dataAccess.UserDataAccess) *AuthDomain {
	return &AuthDomain{userDataAccess: userDataAccess}
}

func (ad *AuthDomain) RegisterUser(user models.User) error {
	foundUser, _ := ad.userDataAccess.GetUserByUsernameOrEmail(user)

	if foundUser != nil {

		return dataAccess.ErrorMissingUsernameOrEmail
	}

	if (user.Username == "" || user.Email == "" || user.Password == "") {
		return ErrorRegisterUserMissingFields
	}

	hash, err := hashPassword(user.Password) 
	
	// Nice, this is standard for saving all user passwords
	// Hadis => Thanks :)

	if err != nil {
		return err
	}

	user.Password = hash

	if err := ad.userDataAccess.CreateUser(user); err != nil {
		return err
	}

	return nil
}

func (ad *AuthDomain) LoginUser(user models.User) (string, error) {
	foundUser, err := ad.userDataAccess.GetUserByUsernameOrEmail(user)

	if err != nil {
		return "", err
	}

	if !checkPassword(user.Password, foundUser.Password) {
		return "", ErrorLoginUserInvalidCredentials
	}

	token, err := generateJWTLoginToken(foundUser)


	if err != nil {
		return "", err
	}

	return token, nil
}
