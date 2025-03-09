package util

import "golang.org/x/crypto/bcrypt"

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