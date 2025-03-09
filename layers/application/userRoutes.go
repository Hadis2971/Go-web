package application

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/Hadis2971/go_web/layers/dataAccess"
)

type UserRouteHandler struct {
	mux *http.ServeMux
	dataAccess *dataAccess.DataAccess
}

type DeleteUserJsonBody struct {
	ID int
}

func NewUserRouteHandler (mux *http.ServeMux, dataAccess *dataAccess.DataAccess) *UserRouteHandler {
	return &UserRouteHandler{mux: mux, dataAccess: dataAccess}
}

func (user UserRouteHandler) HandleDeleteUser (w http.ResponseWriter, r *http.Request) {
	var deleteUserJsonBody DeleteUserJsonBody

	if err := json.NewDecoder(r.Body).Decode(&deleteUserJsonBody); err != nil {
		fmt.Println(err);
	}

	if err := user.dataAccess.DeleteUser(deleteUserJsonBody.ID); err != nil {
		log.Fatal(err);
	}

	w.WriteHeader(http.StatusOK);
}

func (user UserRouteHandler) RegisterRoutes () {
	user.mux.HandleFunc("POST /delete/", user.HandleDeleteUser)
}