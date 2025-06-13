package application

import (
	"fmt"
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
	websocketProductService := service.NewProductWebsocketService()

	defer close(websocketService.ErrorChan)
	defer close(websocketService.BroadcastChan)
	defer close(websocketProductService.ErrorChan)
	defer close(websocketProductService.BroadcastChan)

	userDataAccess := dataAccess.NewUserDataAccess(dbConnection)
	productDataAccess := dataAccess.NewProductDataAccess(dbConnection)

	authDomain := domain.NewAuthDomain(userDataAccess)
	userDomain := domain.NewUserDomain(userDataAccess)
	productDomain := domain.NewProductDomain(productDataAccess)
	chatDomain := domain.NewChatDomain(websocketService)
	wsProductDomain := domain.NewWsProductDomain(websocketProductService)

	fmt.Println("wsProductDomain", wsProductDomain)

	authRouteHandler := NewAuthRouteHandler(authDomain)
	userRouteHandler := NewUserRouteHandler(userDomain)
	productRouteHandler := NewProductRoutes(productDomain, wsProductDomain)
	websocketRoutesHandler := NewWebsocketRoutesHandler(chatDomain)
	wescoketProductRouteHandler := NewProductWebsocketRoutesHandler(wsProductDomain)

	authMux := authRouteHandler.RegisterRoutes()
	userMux := userRouteHandler.RegisterRoutes()
	productMux := productRouteHandler.RegisterRoutes()
	wsChantHandler := websocketRoutesHandler.RegisterRoute()
	wsProductHandler := wescoketProductRouteHandler.RegisterWsProductRoute()

	mux.Handle("/auth/", http.StripPrefix("/auth", authMux))
	mux.Handle("/user/", http.StripPrefix("/user", userMux)) 
	mux.Handle("/product/", http.StripPrefix("/product", productMux))

	mux.Handle("/chat", wsChantHandler)
	mux.Handle("/product/ws", wsProductHandler)

	http.ListenAndServe(app.Port, mux)
}
