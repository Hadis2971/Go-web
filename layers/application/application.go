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
	// I don't think you need to strip the prefix
	// Hadis => You mean it's bad design or in a techical way?
	//			As far as I can see I need it since i plan to send requests to /user/xxx
	//			Since that would be a RESTFUL design
	mux.Handle("/product/", http.StripPrefix("/product", productMux))

	mux.Handle("/chat", wsChantHandler)

	http.ListenAndServe(app.Port, mux)
}
