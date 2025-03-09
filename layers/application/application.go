package application

import (
	"net/http"

	"github.com/Hadis2971/go_web/layers/dataAccess"
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
	dataAccess := dataAccess.NewDataAccess(dbConnection)

	authRouteHandler := NewAuthRouteHandler(mux, dataAccess)
	userRouteHandler := NewUserRouteHandler(mux, dataAccess)

	authRouteHandler.RegisterRoutes();
	userRouteHandler.RegisterRoutes();

	mux.Handle("/auth/", http.StripPrefix("/auth", mux));
	mux.Handle("/user/", http.StripPrefix("/user", mux));

	http.ListenAndServe(app.Port, mux);
} 