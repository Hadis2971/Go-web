package application

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Hadis2971/go_web/layers/domain"
	"github.com/Hadis2971/go_web/models"
)

type AuthRouteHandler struct {
	mux *http.ServeMux
	authDomain *domain.AuthDomain
}

func NewAuthRouteHandler (mux *http.ServeMux, authDomain *domain.AuthDomain) *AuthRouteHandler {
	return &AuthRouteHandler{mux: mux, authDomain: authDomain}
}

func (arh *AuthRouteHandler) HandleRegisterUser (w http.ResponseWriter, r *http.Request) {
	var user models.User;
	defer r.Body.Close()

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)

		return
	}

	newUser, err := arh.authDomain.RegisterUser(user)
	
	if err != nil {
		http.Error(w, err.Error(), http.StatusForbidden)

		return;
	}

	
	jsonData, err := json.Marshal(newUser)

	if err != nil {
		log.Fatal(err)
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}

func (arh AuthRouteHandler) HandleLoginUser (w http.ResponseWriter, r *http.Request) {
	var user models.User

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		log.Fatal(err)
	}

	foundUser, err := arh.authDomain.LoginUser(user)

	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)

		return
	}

	jsonData, err := json.Marshal(foundUser)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	w.WriteHeader(http.StatusFound)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)

}

func (arh AuthRouteHandler) RegisterRoutes () {
	arh.mux.HandleFunc("POST /register/", arh.HandleRegisterUser)
	arh.mux.HandleFunc("POST /login/", arh.HandleLoginUser)
}