package domain

import (
	"github.com/Hadis2971/go_web/layers/dataAccess"
)

type UserDomain struct {
	userDataAccess *dataAccess.UserDataAccess
}

func NewUserDomain (userDataAccess *dataAccess.UserDataAccess) *UserDomain {
	return &UserDomain{userDataAccess: userDataAccess}
}

func (ud *UserDomain) HandleDeleteUser (id int) error {
	if err := ud.userDataAccess.DeleteUser(id); err != nil {
		return err
	}

	return nil
}