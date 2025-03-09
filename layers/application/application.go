package application

import (
	"net/http"

	"github.com/Hadis2971/go_web/layers/dataAccess"
	"github.com/Hadis2971/go_web/layers/domain"
	"github.com/Hadis2971/go_web/layers/service"
)

type Application struct {
	Port string
}

func NewApplication (port string) *Application {
	return &Application{Port: port}
}

func (app Application) Run () {
	mux := http.NewServeMux();
	dbConnection := service.ConnectToDatabase()
	
	userDataAccess := dataAccess.NewUserDataAccess(dbConnection)
	authDomain := domain.NewAuthDomain(userDataAccess)
	userDomain := domain.NewUserDomain(userDataAccess)

	authRouteHandler := NewAuthRouteHandler(mux, authDomain)
	userRouteHandler := NewUserRouteHandler(mux, userDomain)

	authRouteHandler.RegisterRoutes();
	userRouteHandler.RegisterRoutes();

	mux.Handle("/auth/", http.StripPrefix("/auth", mux));
	mux.Handle("/user/", http.StripPrefix("/user", mux));

	http.ListenAndServe(app.Port, mux);
} 