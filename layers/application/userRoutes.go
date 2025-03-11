package application

import (
	"encoding/json"
	"fmt"
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
		fmt.Fprintf(w, "%s", err.Error())
	}

	err := ur.userDomain.HandleDeleteUser(deleteUserJsonBody.ID)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, "%s", err.Error())
	}

	w.WriteHeader(http.StatusOK);
}

func (ur UserRouteHandler) HandleUpdateUser (w http.ResponseWriter, r *http.Request) {
	var updateUserRequestJsonBody dataAccess.UpdateUserRequest

	if err := json.NewDecoder(r.Body).Decode(&updateUserRequestJsonBody); err != nil {
		fmt.Fprintf(w, "%s", err.Error())
	}

	if err := ur.userDomain.HandleUpdateUser(updateUserRequestJsonBody); err != nil {
		fmt.Fprintf(w, "%s", err.Error())
	}

	w.WriteHeader(http.StatusOK)
}

func (ur UserRouteHandler) RegisterRoutes () {
	ur.mux.HandleFunc("POST /delete/", ur.HandleDeleteUser)
	ur.mux.HandleFunc("POST /update/", ur.HandleUpdateUser)
}