package application

import (
	"net/http"
)

type AuthRouteHandler struct {
	mux *http.ServeMux
}

func NewAuthRouteHandler (mux *http.ServeMux) *AuthRouteHandler {
	return &AuthRouteHandler{mux: mux}
}

func (arh AuthRouteHandler) HandleRegisterUser (w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello From AuthRouteHandler -> HandleRegisterUser"))
}

func (arh AuthRouteHandler) HandleLoginUser (w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello From AuthRouteHandler -> HandleLoginUser"))
}

func (arh AuthRouteHandler) RegisterRoutes () {
	arh.mux.HandleFunc("POST /register/", arh.HandleRegisterUser)
	arh.mux.HandleFunc("POST /login/", arh.HandleLoginUser)
}