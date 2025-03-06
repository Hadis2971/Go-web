package application

import (
	"net/http"
)

type Application struct {
	Port string
}

func NewApplication (port string) *Application {
	return &Application{Port: port}
}

func (app Application) Run () {
	mux := http.NewServeMux();

	authRouteHandler := NewAuthRouteHandler(mux);
	userRouteHandler := NewUserRouteHandler(mux)

	authRouteHandler.RegisterRoutes();
	userRouteHandler.RegisterRoutes();

	mux.Handle("/auth/", http.StripPrefix("/auth", mux));
	mux.Handle("/user/", http.StripPrefix("/user", mux));

	
	http.ListenAndServe(app.Port, mux);
} 