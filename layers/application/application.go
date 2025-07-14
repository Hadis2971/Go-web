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
	websocketChatService := service.NewWebsocketService[service.ChatMessage]()
	websocketProductService := service.NewWebsocketService[service.ProductMessage]()

	defer close(websocketChatService.ErrorChan)
	defer close(websocketChatService.BroadcastChan)
	defer close(websocketProductService.ErrorChan)
	defer close(websocketProductService.BroadcastChan)

	userDataAccess := dataAccess.NewUserDataAccess(dbConnection)
	productDataAccess := dataAccess.NewProductDataAccess(dbConnection)
	productOrderDataAccess := dataAccess.NewProductOrderDataAccess(dbConnection)
	productCategoryDataAccess := dataAccess.NewProductCategoryDataAccess(dbConnection)

	authDomain := domain.NewAuthDomain(userDataAccess)
	userDomain := domain.NewUserDomain(userDataAccess)
	productDomain := domain.NewProductDomain(productDataAccess)
	productOrderDomain := domain.NewProductOrderDomain(productOrderDataAccess)
	productCategoryDomain := domain.NewProductCategoryDomain(productCategoryDataAccess)

	chatDomain := domain.NewChatDomain(websocketChatService)
	wsProductDomain := domain.NewWsProductDomain(websocketProductService)

	authRouteHandler := NewAuthRouteHandler(authDomain)
	userRouteHandler := NewUserRouteHandler(userDomain)
	productRouteHandler := NewProductRoutes(productDomain, wsProductDomain)
	productOrderRouteHandler := NewProductOrderRoutes(productOrderDomain)
	productCategoryRouteHandler := NewProductCategoryRoutes(productCategoryDomain)

	websocketRoutesHandler := NewWebsocketRoutesHandler(chatDomain)
	wescoketProductRouteHandler := NewProductWebsocketRoutesHandler(wsProductDomain)

	authMux := authRouteHandler.RegisterRoutes()
	userMux := userRouteHandler.RegisterRoutes()
	productMux := productRouteHandler.RegisterRoutes()
	productOrderMux := productOrderRouteHandler.RegisterRoutes()
	productCategoryMux := productCategoryRouteHandler.RegisterRoutes()

	wsChantHandler := websocketRoutesHandler.RegisterRoute()
	wsProductHandler := wescoketProductRouteHandler.RegisterWsProductRoute()

	mux.Handle("/auth/", http.StripPrefix("/auth", authMux))
	mux.Handle("/user/", http.StripPrefix("/user", userMux)) 
	mux.Handle("/product/", http.StripPrefix("/product", productMux))
	mux.Handle("/product_order/", http.StripPrefix("/product_order", productOrderMux))
	mux.Handle("/product_category/", http.StripPrefix("/product_category", productCategoryMux))

	mux.Handle("/chat", wsChantHandler)
	mux.Handle("/product/ws", wsProductHandler)

	http.ListenAndServe(app.Port, mux)
}
