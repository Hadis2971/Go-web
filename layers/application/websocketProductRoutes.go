package application

import (
	"github.com/Hadis2971/go_web/layers/domain"
	"github.com/Hadis2971/go_web/middlewares"
	"golang.org/x/net/websocket"
)

type WebsocketProductRoutesHandler struct {
	productWsDomain *domain.WsProductDomain
}

func NewProductWebsocketRoutesHandler(productWsDomain *domain.WsProductDomain) *WebsocketProductRoutesHandler {
	return &WebsocketProductRoutesHandler{productWsDomain: productWsDomain}
}

func (wsprh *WebsocketProductRoutesHandler) Handler(ws *websocket.Conn) {
	id := ws.Request().URL.Query().Get("id")

	wsprh.productWsDomain.AddNewWsProductClient(id, ws)

}

func (wsrh *WebsocketProductRoutesHandler) RegisterWsProductRoute () websocket.Handler {
	authMiddleware := middlewares.NewAuthMiddleware()

	return authMiddleware.WithWebsocketRouthAuthentication(websocket.Handler(func(ws *websocket.Conn) {
		wsrh.Handler(ws)
	})) 
}