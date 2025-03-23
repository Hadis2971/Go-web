package application

import (
	"fmt"

	"github.com/Hadis2971/go_web/layers/domain"
	"github.com/Hadis2971/go_web/middlewares"
	"golang.org/x/net/websocket"
)

type WebsocketRoutesHandler struct {
	chatDomain *domain.ChatDomain
}

func NewWebsocketRoutesHandler(chatDomain *domain.ChatDomain) *WebsocketRoutesHandler {
	return &WebsocketRoutesHandler{chatDomain: chatDomain}
}

func (wrh *WebsocketRoutesHandler) Handler(ws *websocket.Conn) {
	id := ws.Request().URL.Query().Get("id")

	wrh.chatDomain.AddNewClient(id, ws)

}

func (wsrh *WebsocketRoutesHandler) RegisterRoute () websocket.Handler {
	authMiddleware := middlewares.NewAuthMiddleware()

	return authMiddleware.WithWebsocketRouthAuthentication(websocket.Handler(func(ws *websocket.Conn) {
		fmt.Println(ws)

		wsrh.Handler(ws)

		
	})) 
}