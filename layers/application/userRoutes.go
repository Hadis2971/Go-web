package application

import (
	"encoding/json"
	"net/http"

	"github.com/Hadis2971/go_web/layers/dataAccess"
	"github.com/Hadis2971/go_web/layers/domain"
)

type UserRouteHandler struct {
	mux *http.ServeMux
	userDomain *domain.UserDomain
}

type DeleteUserJsonBody struct {
	ID int
}

func NewUserRouteHandler (mux *http.ServeMux, userDomain *domain.UserDomain) *UserRouteHandler {
	return &UserRouteHandler{mux: mux, userDomain: userDomain}
}

func (ur UserRouteHandler) HandleDeleteUser (w http.ResponseWriter, r *http.Request) {
	var deleteUserJsonBody DeleteUserJsonBody

	if err := json.NewDecoder(r.Body).Decode(&deleteUserJsonBody); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	err := ur.userDomain.HandleDeleteUser(deleteUserJsonBody.ID)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	w.WriteHeader(http.StatusOK);
}

func (ur UserRouteHandler) HandleUpdateUser (w http.ResponseWriter, r *http.Request) {
	var updateUserRequestJsonBody dataAccess.UpdateUserRequest

	if err := json.NewDecoder(r.Body).Decode(&updateUserRequestJsonBody); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)

		return
	}

	if err := ur.userDomain.HandleUpdateUser(updateUserRequestJsonBody); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	w.WriteHeader(http.StatusOK)
}

func (ur UserRouteHandler) RegisterRoutes () {
	ur.mux.HandleFunc("POST /delete/", ur.HandleDeleteUser)
	ur.mux.HandleFunc("POST /update/", ur.HandleUpdateUser)
}