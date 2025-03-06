package application

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type UserRouteHandler struct {
	mux *http.ServeMux
}

type DeleteUserJsonBody struct {
	Name string
	ID int
}

func NewUserRouteHandler (mux *http.ServeMux) *UserRouteHandler {
	return &UserRouteHandler{mux: mux}
}

func (user UserRouteHandler) HandleCreateUser (w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello From UserRouteHandler -> HandleCreateUser"))
}

func (user UserRouteHandler) HandleDeleteUser (w http.ResponseWriter, r *http.Request) {
	var deleteUserJsonBody DeleteUserJsonBody

	if err := json.NewDecoder(r.Body).Decode(&deleteUserJsonBody); err != nil {
		fmt.Println(err);
	}

	fmt.Println(deleteUserJsonBody)
	fmt.Fprintf(w, "deleteUserJsonBody: %+v \n", deleteUserJsonBody)
}

func (user UserRouteHandler) RegisterRoutes () {
	user.mux.HandleFunc("POST /create/", user.HandleCreateUser)
	user.mux.HandleFunc("POST /delete/", user.HandleDeleteUser)
}