package domain

import (
	"errors"

	"github.com/Hadis2971/go_web/layers/dataAccess"
)

type UserDomain struct {
	userDataAccess dataAccess.IUserDataAccess
}

var ( 
	InternalServerError = errors.New("Internal Server Error!!!") 
)

func NewUserDomain(userDataAccess *dataAccess.UserDataAccess) *UserDomain {
	return &UserDomain{userDataAccess: userDataAccess}
}

func (ud *UserDomain) HandleDeleteUser(id int) error {
	

	if err := ud.userDataAccess.DeleteUser(id); err != nil {
		return InternalServerError
	}

	return nil
}

func (ud *UserDomain) HandleUpdateUser(updateUserRequest dataAccess.UpdateUserRequest) error {
	if err := ud.userDataAccess.UpdateUser(updateUserRequest); err != nil {
		return err
	}

	return nil
}
