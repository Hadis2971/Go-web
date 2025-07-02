package domain

import (
	"errors"

	"github.com/Hadis2971/go_web/layers/dataAccess"
	"github.com/Hadis2971/go_web/models"
)

type UserDomain struct {
	userDataAccess dataAccess.IUserDataAccess
}

func NewUserDomain(userDataAccess dataAccess.IUserDataAccess) *UserDomain {
	return &UserDomain{userDataAccess: userDataAccess}
}

func (ud *UserDomain) HandleDeleteUser(id int) error {

	err := ud.userDataAccess.DeleteUser(id)
	

	if (errors.Is(err, dataAccess.InternalServerError)) {
		return dataAccess.InternalServerError
	}

	if (errors.Is(err, dataAccess.ErrorMissingID)) {
		return dataAccess.ErrorMissingID
	}

	return nil
}

func (ud *UserDomain) HandleUpdateUser(updateUserRequest dataAccess.UpdateUserRequest) error {
	err := ud.userDataAccess.UpdateUser(updateUserRequest)

	if (errors.Is(err, dataAccess.InternalServerError)) {
		return dataAccess.InternalServerError
	}

	if (errors.Is(err, dataAccess.ErrorMissingID)) {
		return dataAccess.ErrorMissingID
	}

	return nil
}

func (ud *UserDomain) HandleGetAllUsersAndTheirOrders() ([]models.UserWithOrders, error) {
	userWithOrders, err := ud.userDataAccess.GetAllUsersAndTheirOrders()

	if (errors.Is(err, dataAccess.ErrorNoUserOrderProductsFound)) {
		return nil, dataAccess.ErrorNoUserOrderProductsFound
	}

	return userWithOrders, nil
}