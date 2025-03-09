package application

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/Hadis2971/go_web/layers/dataAccess"
	"github.com/Hadis2971/go_web/models"
	"github.com/Hadis2971/go_web/util"
)

type AuthRouteHandler struct {
	mux *http.ServeMux
	dataAccess *dataAccess.DataAccess
}

func NewAuthRouteHandler (mux *http.ServeMux, dataAccess *dataAccess.DataAccess) *AuthRouteHandler {
	return &AuthRouteHandler{mux: mux, dataAccess: dataAccess}
}

func (arh *AuthRouteHandler) HandleRegisterUser (w http.ResponseWriter, r *http.Request) {
	var user models.User;
	defer r.Body.Close()

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		log.Fatal(err)
	}

	foundUser, _ := arh.dataAccess.GetUserByUsernameOrEmail(user) 
	
	fmt.Println(foundUser)
	
	if foundUser != nil {
		w.WriteHeader(http.StatusForbidden)
		fmt.Fprintf(w, "%s", errors.New("Username or Email Already Taken!!!").Error())

		return;
	}

	hash, err := util.HashPassword(user.Password)

	if err != nil {
		log.Fatal(err)
	}

	user.Password = hash;

	arh.dataAccess.CreateUser(user);

	jsonData, err := json.Marshal(user)

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

	foundUser, err := arh.dataAccess.GetUserByUsernameOrEmail(user)

	if err != nil {
		log.Fatal(err)
	}

	jsonData, err := json.Marshal(foundUser)

	if err != nil {
		log.Fatal(err)
	}

	w.WriteHeader(http.StatusFound)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)

}

func (arh AuthRouteHandler) RegisterRoutes () {
	arh.mux.HandleFunc("POST /register/", arh.HandleRegisterUser)
	arh.mux.HandleFunc("POST /login/", arh.HandleLoginUser)
}