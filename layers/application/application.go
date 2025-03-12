package application

import (
	"net/http"

	"golang.org/x/net/websocket"

	"github.com/Hadis2971/go_web/layers/dataAccess"
	"github.com/Hadis2971/go_web/layers/domain"
	"github.com/Hadis2971/go_web/layers/service"
)

type Application struct {
	Port string
}

func NewApplication(port string) *Application {
	return &Application{Port: port}
}

func (app Application) Run() {
	mux := http.NewServeMux()
	dbConnection := service.ConnectToDatabase()
	websocketService := service.NewWebsocketService()

	userDataAccess := dataAccess.NewUserDataAccess(dbConnection)
	authDomain := domain.NewAuthDomain(userDataAccess)
	userDomain := domain.NewUserDomain(userDataAccess)
	chatDomain := domain.NewChatDomain(websocketService)

	authRouteHandler := NewAuthRouteHandler(mux, authDomain)
	userRouteHandler := NewUserRouteHandler(mux, userDomain)
	websocketRoutesHandler := NewWebsocketRoutesHandler(chatDomain)

	authRouteHandler.RegisterRoutes()
	userRouteHandler.RegisterRoutes()

	mux.Handle("/auth/", http.StripPrefix("/auth", mux)) // A shared mux is used here, so all endpoints would be accessible by all routes?
	mux.Handle("/user/", http.StripPrefix("/user", mux)) // I don't think you need to strip the prefix
	mux.Handle("/chat", websocket.Handler(func(ws *websocket.Conn) {
		websocketRoutesHandler.Handler(ws)
	}))

	http.ListenAndServe(app.Port, mux)
}
