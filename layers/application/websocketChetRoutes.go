package application

import (
	"github.com/Hadis2971/go_web/layers/domain"
	"github.com/Hadis2971/go_web/middlewares"
	"golang.org/x/net/websocket"
)

type WebsocketChatRoutesHandler struct {
	chatDomain *domain.ChatDomain
}

func NewWebsocketRoutesHandler(chatDomain *domain.ChatDomain) *WebsocketChatRoutesHandler {
	return &WebsocketChatRoutesHandler{chatDomain: chatDomain}
}

func (wrh *WebsocketChatRoutesHandler) Handler(ws *websocket.Conn) {
	id := ws.Request().URL.Query().Get("id")

	wrh.chatDomain.AddNewClient(id, ws)

}

func (wsrh *WebsocketChatRoutesHandler) RegisterRoute () websocket.Handler {
	authMiddleware := middlewares.NewAuthMiddleware()

	return authMiddleware.WithWebsocketRouthAuthentication(websocket.Handler(func(ws *websocket.Conn) {
		wsrh.Handler(ws)
	})) 
}