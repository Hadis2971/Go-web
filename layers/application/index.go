package application

import (
	"fmt"
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

	authRouteHandler := NewAuthRouteHandler(mux);
	userRouteHandler := NewUserRouteHandler(mux)

	authRouteHandler.RegisterRoutes();
	userRouteHandler.RegisterRoutes();

	mux.Handle("/auth/", http.StripPrefix("/auth", mux));
	mux.Handle("/user/", http.StripPrefix("/user", mux));

	fmt.Println(dataAccess);

	http.ListenAndServe(app.Port, mux);
} 