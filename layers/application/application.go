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

	authRouteHandler := NewAuthRouteHandler(authDomain)
	userRouteHandler := NewUserRouteHandler(userDomain)
	websocketRoutesHandler := NewWebsocketRoutesHandler(chatDomain)

	authMux := authRouteHandler.RegisterRoutes()
	userMux := userRouteHandler.RegisterRoutes()

	mux.Handle("/auth/", http.StripPrefix("/auth", authMux))
	mux.Handle("/user/", http.StripPrefix("/user", userMux)) 
	// I don't think you need to strip the prefix
	// Hadis => You mean it's bad design or in a techical way?
	//			As far as I can see I need it since i plan to send requests to /user/xxx
	//			Since that would be a RESTFUL design

	mux.Handle("/chat", websocket.Handler(func(ws *websocket.Conn) {
		websocketRoutesHandler.Handler(ws)
	}))

	http.ListenAndServe(app.Port, mux)
}
