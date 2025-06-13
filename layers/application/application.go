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

func NewApplication(port string) *Application {
	return &Application{Port: port}
}

func (app Application) Run() {
	mux := http.NewServeMux()
	dbConnection := service.ConnectToDatabase()
	websocketService := service.NewWebsocketService()

	defer close(service.NewWebsocketService().ErrorChan)
	defer close(service.NewWebsocketService().BroadcastChan)

	userDataAccess := dataAccess.NewUserDataAccess(dbConnection)
	productDataAccess := dataAccess.NewProductDataAccess(dbConnection)

	authDomain := domain.NewAuthDomain(userDataAccess)
	userDomain := domain.NewUserDomain(userDataAccess)
	productDomain := domain.NewProductDomain(productDataAccess)
	chatDomain := domain.NewChatDomain(websocketService)

	authRouteHandler := NewAuthRouteHandler(authDomain)
	userRouteHandler := NewUserRouteHandler(userDomain)
	productRouteHandler := NewProductRoutes(productDomain)
	websocketRoutesHandler := NewWebsocketRoutesHandler(chatDomain)

	authMux := authRouteHandler.RegisterRoutes()
	userMux := userRouteHandler.RegisterRoutes()
	productMux := productRouteHandler.RegisterRoutes()
	wsChantHandler := websocketRoutesHandler.RegisterRoute()

	mux.Handle("/auth/", http.StripPrefix("/auth", authMux))
	mux.Handle("/user/", http.StripPrefix("/user", userMux)) 
	mux.Handle("/product/", http.StripPrefix("/product", productMux))

	mux.Handle("/chat", wsChantHandler)

	http.ListenAndServe(app.Port, mux)
}
